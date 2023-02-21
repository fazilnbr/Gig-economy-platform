package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type workerRepository struct {
	db *sql.DB
}

// ListJobRequsetFromUser implements interfaces.WorkerRepository
func (c *workerRepository) ListJobRequsetFromUser(pagenation utils.Filter, id int) ([]domain.RequestResponse, utils.Metadata, error) {
	var requests []domain.RequestResponse

	query := `SELECT COUNT(*) OVER(),U.user_name,C.category,R.date,A.house_name,A.place,A.city,A.post,A.pin,A.phone 
				FROM users AS U 
				INNER JOIN requests AS R  ON U.id_login=R.user_id
				INNER JOIN jobs AS J ON J.id_job=R.job_id
				INNER JOIN addresses AS A ON A.id_address=R.address_id
				INNER JOIN categories AS C ON J.category_id=C.id_category 
				WHERE J.id_worker=$1 
				LIMIT $2 OFFSET $3;
	`

	rows, err := c.db.Query(query, id, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var request domain.RequestResponse

		err = rows.Scan(
			&totalRecords,
			&request.Username,
			&request.JobCategory,
			&request.JobDate,
			&request.HouseName,
			request.Place,
			&request.City,
			&request.Post,
			&request.Pin,
			&request.Phone,
		)

		if err != nil {
			return requests, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		requests = append(requests, request)
	}
	if err := rows.Err(); err != nil {
		return requests, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(requests)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return requests, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// DeleteJob implements interfaces.WorkerRepository
func (c *workerRepository) DeleteJob(id int) error {
	query := `DELETE FROM jobs WHERE id_job=$1 RETURNING id_job;`

	var row int
	sql := c.db.QueryRow(query, id)

	sql.Scan(&row)
	if row == 0 {
		return errors.New("There is no item to delete")
	}

	return sql.Err()
}

// ViewJob implements interfaces.WorkerRepository
func (c *workerRepository) ViewJob(id int) ([]domain.WorkerJob, error) {
	var jobs []domain.WorkerJob

	query := `SELECT 
				C.id_category,C.category,J.wage,J.description 
				FROM categories AS C 
				INNER JOIN jobs AS J 
				ON C.id_category = J.category_id 
				WHERE J.id_worker=$1;`

	rows, err := c.db.Query(query, id)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var job domain.WorkerJob

		err = rows.Scan(
			&job.Id,
			&job.JobTitile,
			&job.Wage,
			&job.Description,
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

	query := ` SELECT COUNT (*) FROM jobs WHERE category_id=$1 AND id_worker=$2;`

	rows, err := c.db.Query(query,
		job.CategoryId,
		job.IdWorker,
	)
	fmt.Println("rows : ", rows)
	if err != nil {
		return 0, err
	}
	var Id int

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&Id,
		)

		if err != nil {
			return 0, err
		}
	}

	fmt.Println("id : ", Id)
	if Id != 0 {
		return 0, errors.New("This Job is already added by you")
	}

	query = ` INSERT INTO jobs 
	(category_id,id_worker,wage,description) 
	VALUES 
	($1,$2,$3,$4)  RETURNING id_job;`

	err = c.db.QueryRow(query,
		job.CategoryId,
		job.IdWorker,
		job.Wage,
		job.Description,
	).Scan(
		&Id,
	)

	return Id, err

}

// ListJobCategoryUser implements interfaces.WorkerRepository
func (c *workerRepository) ListJobCategoryUser(pagenation utils.Filter) ([]domain.Category, utils.Metadata, error) {
	var categories []domain.Category

	query := `SELECT COUNT(*) OVER(),
	id_category,
		  category	
		  FROM categories 
		  LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var category domain.Category

		err = rows.Scan(
			&totalRecords,
			&category.IdCategory,
			&category.Category,
		)

		if err != nil {
			return categories, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return categories, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(categories)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return categories, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// ChangePassword implements interfaces.UserRepository
func (c *workerRepository) WorkerChangePassword(changepassword string, id int) (int, error) {
	var Id int
	fmt.Println("id", id)
	query := ` UPDATE users
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
	WHERE login_id = $11
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
	(login_id,name,gender,date_of_birth,house_name,place,post,pin,contact_number,email_id,photo) 
	VALUES 
	($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING login_id;`

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

	query := `SELECT id_login,user_name,password  FROM users WHERE user_name=$1 AND user_type='worker';`

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
func (c *workerRepository) InsertWorker(newWorker domain.User) (int, error) {
	var id int

	query := `INSERT INTO users (user_name,password,user_type) VALUES ($1,$2,$3) RETURNING id_login;`

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

	query = `UPDATE users 
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
