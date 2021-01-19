package dtos

//LoginDto
type SessionUser struct {
	Username string `json:"username" schema:"username"`
	Name string `json:"name" schema:"name"`
	Id string `json:"id" schema:"id"`
}

//SessionResponse is the response of a session request
type SessionResponse struct {
	User *SessionUser `json:"user"`
}