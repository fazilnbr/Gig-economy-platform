package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type workerRepository struct {
	db *sql.DB
}

// AddProfile implements interfaces.WorkerRepository
func (c *workerRepository) AddProfile(workerProfile domain.Profile, id int) (int, error) {
	var Id int
	query := ` INSERT INTO Profiles 
	(id_login,name,gender,date_of_birth,house_name,place,post,pin,contact_number,email_id,photo) 
	VALUES 
	($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id_login;`

	err := c.db.QueryRow(query,
		id,
		workerProfile.Gender,
		workerProfile.Name,
		workerProfile.DateOfBirth,
		workerProfile.HouseName,
		workerProfile.Place,
		workerProfile.Post,
		workerProfile.Pin,
		workerProfile.ContactNumber,
		workerProfile.EmailID,
		workerProfile.Photo,
	).Scan(
		&Id,
	)

	return Id, err
}

// FindWorker implements interfaces.WorkerRepository
func (c *workerRepository) FindWorker(email string) (domain.WorkerResponse, error) {
	var worker domain.WorkerResponse

	query := `SELECT id_login,user_name,password  FROM logins WHERE user_name=$1 AND user_type='worker';`

	err := c.db.QueryRow(query,
		email).Scan(
		&worker.ID,
		&worker.UserName,
		&worker.Password,
	)
	fmt.Print("\n", email, worker, err)

	return worker, err
}

// InsertWorker implements interfaces.WorkerRepository
func (c *workerRepository) InsertWorker(newWorker domain.Login) (int, error) {
	var id int

	query := `INSERT INTO logins (user_name,password,user_type) VALUES ($1,$2,$3) RETURNING id_login;`

	err := c.db.QueryRow(query,
		newWorker.UserName,
		newWorker.Password,
		"worker",
	).Scan(
		&id,
	)

	return id, err
}

// StoreVerificationDetails implements interfaces.WorkerRepository
func (*workerRepository) StoreVerificationDetails(email string, code int) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.WorkerRepository
func (c *workerRepository) VerifyAccount(email string, code int) error {
	var id int

	query := `SELECT id FROM 
	verifications WHERE 
	email = $1 AND code = $2;`
	err := c.db.QueryRow(query, email, code).Scan(&id)

	if err == sql.ErrNoRows {
		return errors.New("Invalid verification code/Email")
	}

	if err != nil {
		return err
	}

	query = `UPDATE logins 
	SET
	verification = $1
	WHERE
	user_name = $2 ;`
	err = c.db.QueryRow(query, true, email).Err()
	log.Println("Updating User verification: ", err)
	if err != nil {
		return err
	}

	return nil
}

func NewWorkerRepo(db *sql.DB) interfaces.WorkerRepository {
	return &workerRepository{
		db: db,
	}
}
