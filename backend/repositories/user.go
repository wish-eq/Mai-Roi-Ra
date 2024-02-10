package repository

import (
	"log"
	"time"

	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/models"
	st "github.com/2110366-2566-2/Mai-Roi-Ra/backend/pkg/struct"
	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/utils"
	"gorm.io/gorm"
)

// UserRepository represents the repository for the User model.
type UserRepository struct {
	DB *gorm.DB
}

type IUserRepository interface {
	GetUserByID(userID string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *st.CreateUserRequest) (*st.CreateUserResponse, error)
}

// NewUserRepository creates a new instance of the UserRepository.
// oldone-func NewUserRepository(c *gin.Context, db *gorm.DB) *UserRepository {
func NewUserRepository(
	DB *gorm.DB,
) IUserRepository {
	return &UserRepository{
		DB: DB,
	}
}

// GetAllUsers retrieves all users from the database.
func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUserByID retrieves a user by their ID.
func (repo *UserRepository) GetUserByID(userID string) (*models.User, error) {
	log.Println("[REPO: GetUserByID]: Called")
	var user models.User
	result := repo.DB.Where("user_id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email address.
func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByPhoneNumber retrieves a user by their phone number.
func (repo *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser adds a new user to the database.
func (r *UserRepository) CreateUser(req *st.CreateUserRequest) (*st.CreateUserResponse, error) {
	log.Println("[Repo: CreateUser]: Called")
	birthDate, err := utils.StringToTime(req.BirthDate)
	if err != nil {
		log.Println("[Repo: CreateUser] Error parsing BirthDate to time.Time format:", err)
		return nil, err
	}

	userModel := models.User{
		UserID:                   utils.GenerateUUID(),
		PaymentGatewayCustomerID: utils.GenerateUUID(),
		PhoneNumber:              req.PhoneNumber,
		BirthDate:                birthDate,
		Email:                    req.Email,
		FirstName:                req.FirstName,
		LastName:                 req.LastName,
		UserImage:                req.UserImage,
		CreatedAt:                time.Now(),
	}

	trans := r.DB.Begin().Debug()
	if err := trans.Create(&userModel).Error; err != nil {
		trans.Rollback()
		log.Println("[Repo: CreateUser]: Insert data in Users table error:", err)
		return nil, err
	}

	if err := trans.Commit().Error; err != nil {
		trans.Rollback()
		log.Println("[Repo: CreateUser]: Call ORM DB Commit error:", err)
		return nil, err
	}

	return &st.CreateUserResponse{
		UserId: userModel.UserID, // Assuming your User model has an ID field
	}, nil
}

// UpdateUser updates an existing user in the database.
func (repo *UserRepository) UpdateUser(user *models.User) error {
	result := repo.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user from the database.
func (repo *UserRepository) DeleteUser(userID string) error {
	result := repo.DB.Delete(&models.User{}, "user_id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
