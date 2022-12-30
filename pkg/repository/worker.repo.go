package repository

import (
	"database/sql"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type workerRepository struct {
	db *sql.DB
}

// AddProfile implements interfaces.WorkerRepository
func (*workerRepository) AddProfile(workerProfile domain.Worker, id int) (int, error) {
	panic("unimplemented")
}

// FindWorker implements interfaces.WorkerRepository
func (*workerRepository) FindWorker(email string) (domain.WorkerResponse, error) {
	panic("unimplemented")
}

// InsertWorker implements interfaces.WorkerRepository
func (*workerRepository) InsertWorker(newWorker domain.Login) (int, error) {
	panic("unimplemented")
}

// StoreVerificationDetails implements interfaces.WorkerRepository
func (*workerRepository) StoreVerificationDetails(email string, code int) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.WorkerRepository
func (*workerRepository) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewWorkerRepo(db *sql.DB) interfaces.WorkerRepository {
	return &workerRepository{
		db: db,
	}
}
