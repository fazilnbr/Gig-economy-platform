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

// ViewJob implements interfaces.WorkerRepository
func (c *workerRepository) ViewJob(id int) ([]domain.WorkerJob, error) {
	var jobs []domain.WorkerJob

	query := `	SELECT
 			   	c.category
				FROM
    			categories AS C
				INNER JOIN jobs AS J 
   				ON C.id_category = J.id_category
				WHERE J.id_worker=$1;`

	rows, err := c.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var job domain.WorkerJob

		err = rows.Scan(
			&job.JobTitile,
		)

		if err != nil {
			return jobs, err
		}

		jobs = append(jobs, job)
	}
	if err := rows.Err(); err != nil {
		return jobs, err
	}
	return jobs, nil
}

// AddJob implements interfaces.WorkerRepository
func (c *workerRepository) AddJob(job domain.Job) (int, error) {

	var Id int
	query := ` INSERT INTO jobs 
	(id_category,id_worker) 
	VALUES 
	($1,$2) RETURNING id_job;`

	err := c.db.QueryRow(query,
		job.IdCategory,
		job.IdWorker,
	).Scan(
		&Id,
	)

	return Id, err

}

// ListJobCategoryUser implements interfaces.WorkerRepository
func (c *workerRepository) ListJobCategoryUser() ([]domain.Category, error) {
	var categories []domain.Category

	query := `select * from categories;`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category domain.Category

		err = rows.Scan(
			&category.IdCategory,
			&category.Category,
		)

		if err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}

// ChangePassword implements interfaces.UserRepository
func (c *workerRepository) WorkerChangePassword(changepassword string, id int) (int, error) {
	var Id int
	fmt.Println("id", id)
	query := ` UPDATE logins
	SET password = $1
	WHERE id_login = $2
	RETURNING id_login;
	`

	err := c.db.QueryRow(query,
		changepassword,
		id,
	).Scan(
		&Id,
	)

	return Id, err
}

// EditProfile implements interfaces.UserRepository
func (c *workerRepository) WorkerEditProfile(userProfile domain.Profile, id int) (int, error) {
	var Id int

	query := ` UPDATE profiles
	SET name = $1,
		gender = $2,
		date_of_birth = $3, 
		house_name = $4,
		place = $5, 
		post = $6, 
		pin = $7,
		contact_number = $8, 
		email_id = $9, 
		photo = $10
	WHERE id_login = $11
	RETURNING id_user;
	`

	err := c.db.QueryRow(query,
		userProfile.Name,
		userProfile.Gender,
		userProfile.DateOfBirth,
		userProfile.HouseName,
		userProfile.Place,
		userProfile.Post,
		userProfile.Pin,
		userProfile.ContactNumber,
		userProfile.EmailID,
		userProfile.Photo,
		id,
	).Scan(
		&Id,
	)

	fmt.Printf("\n\nerr : %v\n%v\n\n", userProfile, err)

	return Id, err
}

// AddProfile implements interfaces.WorkerRepository
func (c *workerRepository) WorkerAddProfile(workerProfile domain.Profile, id int) (int, error) {
	var Id int
	query := ` INSERT INTO Profiles 
	(id_login,name,gender,date_of_birth,house_name,place,post,pin,contact_number,email_id,photo) 
	VALUES 
	($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id_login;`

	err := c.db.QueryRow(query,
		id,
		workerProfile.Name,
		workerProfile.Gender,
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
