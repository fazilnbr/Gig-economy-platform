package domain

type UserResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"first_name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
