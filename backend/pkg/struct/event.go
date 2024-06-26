package structure

type CreateEventRequest struct {
	OrganizerId    string  `json:"organizer_id" binding:"required"`
	StartDate      string  `json:"start_date" binding:"required"`
	EndDate        string  `json:"end_date" binding:"required"`
	Status         string  `json:"status" binding:"required"`
	ParticipantFee float64 `json:"participant_fee" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	EventName      string  `json:"event_name" binding:"required"`
	Activities     string  `json:"activities" binding:"required"`
	EventImage     string  `json:"event_image" binding:"required"`
	LocationName   string  `json:"location_name" binding:"required"`
	District       string  `json:"district" binding:"required"`
	City           string  `json:"city" binding:"required"`
	EventId        string  `json:"event_id"`
}

type CreateEventResponse struct {
	EventId string `json:"event_id"`
}

type UpdateEventRequest struct {
	EventId        string  `json:"event_id"`
	StartDate      string  `json:"start_date" binding:"required"`
	EndDate        string  `json:"end_date" binding:"required"`
	Status         string  `json:"status" binding:"required"`
	ParticipantFee float64 `json:"participant_fee" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	EventName      string  `json:"event_name" binding:"required"`
	Activities     string  `json:"activities" binding:"required"`
	LocationName   string  `json:"location_name" binding:"required"`
	District       string  `json:"district" binding:"required"`
	City           string  `json:"city" binding:"required"`
}

type GetEventList struct {
	EventId     string `json:"event_id"`
	EventName   string `json:"event_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
	Status      string `json:"status"`
	EventImage  string `json:"event_image"`
	City        string `json:"city"`
	District    string `json:"district"`
}

type GetEventListsRequest struct {
	OrganizerId string `json:"organizer_id"`
	Filter      string `json:"filter"`
	Sort        string `json:"sort"`
	Search      string `json:"search"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}

type GetEventListsResponse struct {
	TotalPages int            `json:"total_pages"`
	TotalEvent int            `json:"total_events"`
	EventLists []GetEventList `json:"event_lists"`
}

type GetEventListByStartDate struct {
	EventId      string `json:"event_id"`
	OrganizerId  string `json:"organizer_id"`
	EventName    string `json:"event_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	EventImage   string `json:"event_image"`
	LocationName string `json:"location_name"`
}

type GetEventListsByStartDateRequest struct {
	StartDate string `json:"start_date"`
}

type GetEndedEventList struct {
	EventId     string  `json:"event_id"`
	EventName   string  `json:"event_name"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	EventImage  string  `json:"event_image"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	AverageRate float64 `json:"average_rate"`
}

type GetEndedEventListsResponse struct {
	EventLists []GetEndedEventList `json:"event_lists"`
}

type GetEventListsByStartDateResponse struct {
	EventLists []GetEventListByStartDate `json:"event_lists"`
}

type GetEventDataByIdResponse struct {
	EventId          string         `json:"event_id"`
	OrganizerId      string         `json:"organizer_id"`
	UserId           string         `json:"admin_id"`
	LocationId       string         `json:"location_id"`
	FirstName        string         `json:"first_name"`
	LastName         string         `json:"last_name"`
	StartDate        string         `json:"start_date"`
	EndDate          string         `json:"end_date"`
	Status           string         `json:"status"`
	ParticipantFee   float64        `json:"participant_fee"`
	Description      string         `json:"description"`
	EventName        string         `json:"event_name"`
	Deadline         string         `json:"deadline"`
	Activities       string         `json:"activities"`
	EventImage       string         `json:"event_image"`
	Country          string         `json:"country"`
	City             string         `json:"city"`
	District         string         `json:"district"`
	LocationName     string         `json:"location_name"`
	ParticipantCount int            `json:"participant_count"`
	AnnouncementList []Announcement `json:"announcement_list"`
}

type Participant struct {
	UserId         string  `json:"user_id"`
	FirstName      string  `json:"first_name" binding:"required"`
	LastName       string  `json:"last_name" binding:"required"`
	Username       string  `json:"username" binding:"required"`
	NumParticipant int     `json:"num_participant" binding:"required"`
	PhoneNumber    *string `json:"phone_number" `
	Email          *string `json:"email"`
	BirthDate      string  `json:"birth_date" binding:"required"`
	UserImage      string  `json:"user_image"`
	Address        string  `json:"address" binding:"required"`
	District       string  `json:"district" binding:"required"`
	Province       string  `json:"province" binding:"required"`
}

type GetParticipantListsRequest struct {
	EventId string `json:"event_id" binding:"required"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

type GetParticipantListsResponse struct {
	ParticipantList []Participant `json:"participant_list"`
}

type VerifyEventRequest struct {
	EventId string `json:"event_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
	AdminId string `json:"admin_id" binding:"required"`
}
