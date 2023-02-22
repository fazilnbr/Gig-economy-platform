package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type workerService struct {
	workerRepo interfaces.WorkerRepository
}

// ListAcceptedJobRequsetFromUser implements interfaces.WorkerUseCase
func (c *workerService) ListAcceptedJobRequsetFromUser(pagenation utils.Filter, id int) (*[]domain.RequestResponse, *utils.Metadata, error) {
	requsets, Metadata, err := c.workerRepo.ListAcceptedJobRequsetFromUser(pagenation, id)

	return &requsets, &Metadata, err
}

// ListJobRequsetFromUser implements interfaces.WorkerUseCase
func (c *workerService) ListPendingJobRequsetFromUser(pagenation utils.Filter, id int) (*[]domain.RequestResponse, *utils.Metadata, error) {
	requsets, Metadata, err := c.workerRepo.ListPendingJobRequsetFromUser(pagenation, id)

	return &requsets, &Metadata, err
}

// DeleteJob implements interfaces.WorkerUseCase
func (c *workerService) DeleteJob(id int) error {
	err := c.workerRepo.DeleteJob(id)

	return err
}

// ViewJob implements interfaces.WorkerUseCase
func (c *workerService) ViewJob(id int) ([]domain.WorkerJob, error) {
	jobs, err := c.workerRepo.ViewJob(id)

	return jobs, err
}

// AddJob implements interfaces.WorkerUseCase
func (c *workerService) AddJob(job domain.Job) (int, error) {
	id, err := c.workerRepo.AddJob(job)

	return id, err
}

// ListJobCategoryUser implements interfaces.WorkerUseCase
func (c *workerService) ListJobCategoryUser(pagenation utils.Filter) (*[]domain.Category, utils.Metadata, error) {
	categories, Metadata, err := c.workerRepo.ListJobCategoryUser(pagenation)

	return &categories, Metadata, err
}

// VerifyPassword implements interfaces.UserUseCase
func (c *workerService) WorkerVerifyPassword(changepassword domain.ChangePassword, id int) error {
	user, err := c.workerRepo.FindWorker(changepassword.Email)
	if err != nil {
		return errors.New("Invalid User")
	}

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", user, changepassword.OldPassword)

	isValidPassword := VerifyPassword(changepassword.OldPassword, user.Password)
	if !isValidPassword {
		return errors.New("Invalid Password")
	}
	return nil
}

// ChangePassword implements interfaces.UserUseCase
func (c *workerService) WorkerChangePassword(changepassword string, id int) error {
	//hashing password
	changepassword = HashPassword(changepassword)
	_, err := c.workerRepo.WorkerChangePassword(changepassword, id)

	return err

}

// EditProfile implements interfaces.UserUseCase
func (c *workerService) WorkerEditProfile(userProfile domain.Profile, id int) error {
	_, err := c.workerRepo.WorkerEditProfile(userProfile, id)

	return err

}

// AddProfile implements interfaces.WorkerUseCase
func (c *workerService) AddProfile(workerProfile domain.Profile, id int) error {
	_, err := c.workerRepo.WorkerAddProfile(workerProfile, id)

	return err
}

// CreateUser implements interfaces.WorkerUseCase
func (c *workerService) CreateUser(newWorker domain.User) error {
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
