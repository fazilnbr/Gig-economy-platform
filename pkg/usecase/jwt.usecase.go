package usecase

import (
	"log"
	"os"
	"time"

	model "github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/golang-jwt/jwt/v4"
)

type JWTUseCase struct {
	SecretKey string
}

// GenerateToken implements interfaces.JWTUseCase
func (j *JWTUseCase) GenerateToken(user_id int, username string, role string) string {
	claims := &model.SignedDetails{
		User_Id:  user_id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken
}

// GetTokenFromString implements interfaces.JWTUseCase
func (*JWTUseCase) GetTokenFromString(signedToken string, claims *services.SignedDetails) (*jwt.Token, error) {
	panic("unimplemented")
}

// VerifyToken implements interfaces.JWTUseCase
func (*JWTUseCase) VerifyToken(signedToken string) (bool, *services.SignedDetails) {
	panic("unimplemented")
}

func NewJWTUserService() services.JWTUseCase {
	return &JWTUseCase{
		SecretKey: os.Getenv("USER_KEY"),
	}
}
