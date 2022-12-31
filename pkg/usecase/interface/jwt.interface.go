package interfaces

import (
	model "github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/golang-jwt/jwt/v4"
)

type JWTUseCase interface {
	GenerateToken(user_id int, username string, role string) string
	VerifyToken(signedToken string) (bool, *model.SignedDetails)
	GetTokenFromString(signedToken string, claims *model.SignedDetails) (*jwt.Token, error)
}
