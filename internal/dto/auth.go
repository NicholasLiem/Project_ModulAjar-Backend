package dto

type LoginDTO struct {
	Identifier string `json:"identifier,omitempty"`
	Password   string `json:"password,omitempty"`
}

type RegisterDTO struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
