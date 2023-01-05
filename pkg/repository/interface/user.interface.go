package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type UserRepository interface {
	InsertUser(login domain.Login) (int, error)
	FindUser(email string) (domain.UserResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	UserAddProfile(Profile domain.Profile, id int) (int, error)
	UserEditProfile(Profile domain.Profile, id int) (int, error)
	UserChangePassword(changepassword string, id int) (int, error)
	// verifyPassword(changepassword string, id int) (int, error)
}
