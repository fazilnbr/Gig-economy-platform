package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type WorkerUseCase interface {
	CreateUser(newWorker domain.Login) error
	FindWorker(email string) (*domain.WorkerResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
	AddProfile(workerProfile domain.Worker, id int) error
}
