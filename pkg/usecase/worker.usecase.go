package usecase

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type workerService struct {
	workerRepo interfaces.WorkerRepository
}

// AddProfile implements interfaces.WorkerUseCase
func (*workerService) AddProfile(workerProfile domain.Worker, id int) error {
	panic("unimplemented")
}

// CreateUser implements interfaces.WorkerUseCase
func (*workerService) CreateUser(newWorker domain.Login) error {
	panic("unimplemented")
}

// FindWorker implements interfaces.WorkerUseCase
func (*workerService) FindWorker(email string) (*domain.WorkerResponse, error) {
	panic("unimplemented")
}

// SendVerificationEmail implements interfaces.WorkerUseCase
func (*workerService) SendVerificationEmail(email string) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.WorkerUseCase
func (*workerService) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewWorkerService(workerRepo interfaces.WorkerRepository) services.WorkerUseCase {
	return &workerService{
		workerRepo: workerRepo,
	}
}
