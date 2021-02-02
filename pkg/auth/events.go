package auth

const AccountCreatedEvent = "AccountCreated"
const AccountDeletedEvent = "AccountDeleted"

type AccountCommand struct {
	UserID    string `json:"userId"`
	EventType string `json:"type"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
