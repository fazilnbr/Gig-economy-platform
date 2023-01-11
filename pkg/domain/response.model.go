package domain

import "github.com/golang-jwt/jwt/v4"

type AdminResponse struct {
	ID       int    `json:"id_login"`
	Username string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     int    `json:"role"`
	Token    string `json:"token,omitempty"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	UserName     string `json:"first_name"`
	Password     string `json:"password"`
	Verification bool   `json:"verification"`
	Token        string `json:"token"`
}

type WorkerResponse struct {
	ID           int    `json:"id"`
	UserName     string `json:"first_name"`
	Password     string `json:"password"`
	Verification bool   `json:"verification"`
	Token        string `json:"token"`
}

type SignedDetails struct {
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	Role     string `json:"role"`
	Source   string `json:"source"`
	jwt.StandardClaims
}

type ChangePassword struct {
	Email       string `json:"email" binding:"required"`
	OldPassword string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}
