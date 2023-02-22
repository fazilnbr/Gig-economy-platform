package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type WorkerRepository interface {
	InsertWorker(newWorker domain.User) (int, error)
	// AddProfile(workerProfile domain.Profile, id int) (int, error)
	FindWorker(email string) (domain.WorkerResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
	WorkerAddProfile(Profile domain.Profile, id int) (int, error)
	WorkerEditProfile(Profile domain.Profile, id int) (int, error)
	WorkerChangePassword(changepassword string, id int) (int, error)
	ListJobCategoryUser(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error)
	AddJob(job domain.Job) (int, error)
	ViewJob(id int) ([]domain.WorkerJob, error)
	DeleteJob(id int) error
	ListPendingJobRequsetFromUser(pagenation utils.Filter, id int) ([]domain.RequestResponse, utils.Metadata, error)
	ListAcceptedJobRequsetFromUser(pagenation utils.Filter, id int) ([]domain.RequestResponse, utils.Metadata, error)
	AcceptJobRequest(id int) error
}
