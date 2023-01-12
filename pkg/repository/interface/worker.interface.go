package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type WorkerRepository interface {
	InsertWorker(newWorker domain.Login) (int, error)
	// AddProfile(workerProfile domain.Profile, id int) (int, error)
	FindWorker(email string) (domain.WorkerResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	WorkerAddProfile(Profile domain.Profile, id int) (int, error)
	WorkerEditProfile(Profile domain.Profile, id int) (int, error)
	WorkerChangePassword(changepassword string, id int) (int, error)
	ListJobCategoryUser() ([]domain.Category, error)
	AddJob(job domain.Job) (int, error)
	ViewJob(id int) ([]domain.WorkerJob, error)
	DeleteJob(id int) error
}
