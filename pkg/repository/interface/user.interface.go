package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type UserRepository interface {
	InsertUser(login domain.User) (int, error)
	FindUser(email string) (domain.UserResponse, error)
	StoreVerificationDetails(email string, code string) error
	VerifyAccount(email string, code string) error
	UserAddProfile(Profile domain.Profile, id int) (int, error)
	UserEditProfile(Profile domain.Profile, id int) (int, error)
	UserChangePassword(changepassword string, id int) (int, error)
	ListWorkersWithJob(pagenation utils.Filter) ([]domain.ListJobsWithWorker, utils.Metadata, error)
	SearchWorkersWithJob(pagenation utils.Filter, key string) ([]domain.ListJobsWithWorker, utils.Metadata, error)
	AddToFavorite(favorite domain.Favorite) (int, error)
	CheckInFevorite(favorite domain.Favorite) (int, error)
	ListFevorite(pagenation utils.Filter, id int) ([]domain.ListFavorite, utils.Metadata, error)
	AddAddress(address domain.Address) (int, error)
	ListAddress(id int) ([]domain.Address, error)
	DeleteAddress(id int, userid int) error
	SendJobRequest(request domain.Request) (int, error)
	CheckInRequest(request domain.Request) (int, error)
	DeleteJobRequest(requestId int, userid int) error
}
