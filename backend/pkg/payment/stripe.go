package payment

import (
	"fmt"
	"log"

	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/app/config"
	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/client"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

type StripeService struct {
	Client *client.API
}

type IStripeService interface {
	GetPaymentIntent(req *st.GetPaymentIntentByIdRequest) (*st.GetPaymentIntentByIdResponse, error)
	CreatePaymentIntent(amount int64, currency string) (*stripe.PaymentIntent, error)
	IssueRefund(paymentIntentID string) (*stripe.Refund, error)
	TransferToOrganizer(amount int64, currency, destinationAccountID string) (*stripe.Transfer, error)
	CreatePaymentIntentID(amount int64, currency string) (*string, error)
}

func NewStripeService() *StripeService {
	cfg, err := config.NewConfig(func() string {
		return ".env"
	}())
	if err != nil {
		log.Println("[Pkg]: Error initializing .env")
		// return nil
	}
	log.Println("Config path from Stripe:", cfg)

	publicKey, secretKey := cfg.Stripe.PublicKey, cfg.Stripe.SecretKey
	if publicKey == "" || secretKey == "" {
		log.Println("[Pkg: NewStripeService] Please specify public and secret key in your .env")
		return nil
	}

	stripe.Key = cfg.Stripe.SecretKey
	sc := &client.API{}
	sc.Init(stripe.Key, nil)
	return &StripeService{
		Client: sc,
	}
}

func (s *StripeService) GetPaymentIntent(req *st.GetPaymentIntentByIdRequest) (*stripe.PaymentIntent, error) {
	log.Println("[Pkg: GetPaymentIntent] Called")
	pi, err := s.Client.PaymentIntents.Get(req.PaymentIntentId, nil)
	if err != nil {
		log.Println("[Pkg: GetPaymentIntent]: Error fetching payment intent:", err)
		return nil, err
	}
	return pi, nil
}

func (s *StripeService) CreatePaymentIntent(amount int64, currency string, eventId string, userId string, paymentType int) (*st.CreatePaymentResponse, error) {
	log.Println("[Pkg: CreatePaymentIntent] Called")
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount * 100), // Amount in smallest currency unit (e.g., cents for USD)
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card", "promptpay"}),
		Metadata: map[string]string{
			"event_id": eventId,
			"user_id":  userId,
		},
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		log.Println("[Pkg CreatePaymentIntent", err)
	}

	responseData := &st.CreatePaymentResponse{
		PaymentIntentId:     pi.ID,
		PaymentClientSecret: pi.ClientSecret,
		PaymentMethodType:   pi.PaymentMethodTypes[paymentType-1], // 1: card, 2: PromptPay
		TransactionAmount:   float64(int(pi.Amount) / 100),
		Status:              string(pi.Status),
		MetaData:            pi.Metadata,
	}

	return responseData, nil
}

func (s *StripeService) IssueRefund(paymentIntentID string) (*stripe.Refund, error) {
	log.Println("[Pkg: IssueRefund] Called")
	refundParams := &stripe.RefundParams{
		PaymentIntent: stripe.String(paymentIntentID),
	}
	return s.Client.Refunds.New(refundParams)
}

// stripe acc -> stripe acc
func (s *StripeService) TransferToOrganizer(amount int64, currency, destinationAccountID string, paymentIntentID string, eventID string, userID string) (*stripe.Transfer, error) {
	log.Println("[Pkg: TransferToOrganizer] Called")
	transferParams := &stripe.TransferParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(currency),
		Destination: stripe.String(destinationAccountID),
		Metadata: map[string]string{
			"payment_intent_id": paymentIntentID,
			"event_id":          eventID,
			"user_id":           userID,
		},
	}
	return s.Client.Transfers.New(transferParams)
}

func (s *StripeService) ConfirmPaymentIntent(paymentIntentId string) error {
	// Set Stripe API key
	cfg, err := config.NewConfig(func() string {
		return ".env"
	}())
	if err != nil {
		log.Println("[Pkg]: Error initializing .env")
		// return nil
	}
	paymentMethod := "pm_card_th_credit"

	redirectURL := fmt.Sprintf("%s/homepage", cfg.App.FrontendURL)

	_, err = paymentintent.Confirm(
		paymentIntentId,
		&stripe.PaymentIntentConfirmParams{
			CaptureMethod: stripe.String("automatic"),
			PaymentMethod: stripe.String(paymentMethod),
			ReturnURL:     stripe.String(redirectURL),
		},
	)
	if err != nil {
		log.Println("[Pkg: ConfirmPaymentIntent]: err from stripe when calling confirm method")
		return err
	}
	return nil
}

func (s *StripeService) CreateTransferPaymentIntent(amount int64, currency string, receiverStripeAccountID string, userId string, eventId string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount), // Amount in smallest currency unit (e.g., cents for USD)
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card", "promptpay"}), // Optional: specify payment method types
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(receiverStripeAccountID), // Specify the Stripe account ID of the receiver
		},
		Metadata: map[string]string{
			"event_id": eventId,
			"user_id":  userId,
		},
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		log.Println("[Pkg CreateTransferPaymentIntent] Error creating PaymentIntent:", err)
		return nil, err
	}

	return intent, nil
}
