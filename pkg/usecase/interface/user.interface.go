package interfaces

import "github.com/fazilnbr/project-workey/pkg/domain"

type UserUseCase interface {
	CreateUser(newUser domain.Login) error
	FindUser(email string) (*domain.UserResponse, error)
	VerifyUser(email string, password string) error
	AddProfile(userProfile domain.Profile, id int) error
	UserEditProfile(userProfile domain.Profile, id int) error
	UserVerifyPassword(changepassword domain.ChangePassword, id int) error
	UserChangePassword(changepassword string, id int) error
}
