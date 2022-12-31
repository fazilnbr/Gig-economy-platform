package usecase

import (
	"fmt"
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

// // GetTokenFromString implements interfaces.JWTUseCase
func (j *JWTUseCase) GetTokenFromString(signedToken string, claims *model.SignedDetails) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

}

// VerifyToken implements interfaces.JWTUseCase
func (j *JWTUseCase) VerifyToken(signedToken string) (bool, *model.SignedDetails) {
	claims := &model.SignedDetails{}
	token, _ := j.GetTokenFromString(signedToken, claims)

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func NewJWTUserService() services.JWTUseCase {
	return &JWTUseCase{
		SecretKey: os.Getenv("USER_KEY"),
	}
}
