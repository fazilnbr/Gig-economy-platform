package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type WorkerRepository interface {
	InsertWorker(newWorker domain.Login) (int, error)
	AddProfile(workerProfile domain.Profile, id int) (int, error)
	FindWorker(email string) (domain.WorkerResponse, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
}
