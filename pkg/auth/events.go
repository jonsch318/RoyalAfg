package auth

//AccountCreatedEvent is the name of the event that a new account was created
const AccountCreatedEvent = "AccountCreated"

//AccizbtDeletedEvent is the name of the event that a account was deleted
const AccountDeletedEvent = "AccountDeleted"

type AccountCommand struct {
	EventType string `json:"type"`
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
