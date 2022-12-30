package interfaces

import (
	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	User_Id  int
	Username string
	Role     string
	jwt.StandardClaims
}
type JWTUseCase interface {
	GenerateToken(user_id int, username string, role string) string
	VerifyToken(signedToken string) (bool, *SignedDetails)
	GetTokenFromString(signedToken string, claims *SignedDetails) (*jwt.Token, error)
}
