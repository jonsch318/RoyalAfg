package dto

// LoginDto defines the object for the api login request
type LoginDto struct {
	Username   string `json:"username" schema:"username"`
	Password   string `json:"password" schema:"password"`
	RememberMe bool   `json:"rememberme" schema:"rememberme"`
}
