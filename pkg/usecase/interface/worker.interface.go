package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type WorkerUseCase interface {
	CreateUser(newWorker domain.Login) error
	FindWorker(email string) (*domain.WorkerResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
	AddProfile(workerProfile domain.Profile, id int) error
	WorkerEditProfile(userProfile domain.Profile, id int) error
	WorkerVerifyPassword(changepassword domain.ChangePassword, id int) error
	WorkerChangePassword(changepassword string, id int) error
	ListJobCategoryUser() ([]domain.Category, error)
	AddJob(job domain.Job) (int, error)
	ViewJob(id int) ([]domain.WorkerJob, error)
	DeleteJob(id int) error
}
