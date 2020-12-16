package dto

// RegisterDto defines the dto for the user account registration
type RegisterDto struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"fullName"`
	Birthdate  int64  `json:"birthdate"`
	RememberMe bool   `json:"rememberme"`
}
