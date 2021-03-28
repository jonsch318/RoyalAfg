package dtos

//LoginDto
type SessionUser struct {
	Username string `json:"username" schema:"username"`
	Name     string `json:"name" schema:"name"`
	Id       string `json:"id" schema:"id"`
}

//SessionResponse is the response of a session request
type SessionResponse struct {
	User *SessionUser `json:"user"`
}

// LoginDto defines the object for the api login request
type LoginDto struct {
	Username   string `json:"username" schema:"username"`
	Password   string `json:"password" schema:"password"`
	RememberMe bool   `json:"rememberme" schema:"rememberme"`
}

// RegisterDto defines the dto for the user account registration
type RegisterDto struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"fullName"`
	Birthdate  int64  `json:"birthdate"`
	RememberMe bool   `json:"rememberme"`
}
