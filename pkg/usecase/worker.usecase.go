package usecase

import (
	"database/sql"
	"errors"

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
func (c *workerService) CreateUser(newWorker domain.Login) error {
	_, err := c.workerRepo.FindWorker(newWorker.UserName)

	if err == nil {
		return errors.New("Username Already Exists")
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Hashing the password

	newWorker.Password = HashPassword(newWorker.Password)

	_, err = c.workerRepo.InsertWorker(newWorker)

	return err

}

// FindWorker implements interfaces.WorkerUseCase
func (c *workerService) FindWorker(email string) (*domain.WorkerResponse, error) {
	user, err := c.workerRepo.FindWorker(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
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
