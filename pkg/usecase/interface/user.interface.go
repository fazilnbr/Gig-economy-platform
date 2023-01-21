package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type UserUseCase interface {
	CreateUser(newUser domain.User) error
	FindUser(email string) (*domain.UserResponse, error)
	VerifyUser(email string, password string) error
	AddProfile(userProfile domain.Profile, id int) error
	UserEditProfile(userProfile domain.Profile, id int) error
	UserVerifyPassword(changepassword domain.ChangePassword, id int) error
	UserChangePassword(changepassword string, id int) error
	ListWorkersWithJob(pagenation utils.Filter) (*[]domain.ListJobsWithWorker, *utils.Metadata, error)
	SearchWorkersWithJob(pagenation utils.Filter, key string) (*[]domain.ListJobsWithWorker, *utils.Metadata, error)
	AddToFavorite(favorite domain.Favorite) (int, error)
	ListFevorite(pagenation utils.Filter, id int) (*[]domain.ListFavorite, *utils.Metadata, error)
	AddAddress(address domain.Address) (int, error)
	ListAddress(id int) (*[]domain.Address, error)
	DeleteAddress(id int) error
}
