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

const (
	listjob = `select COUNT(*) OVER(),t1.id_job,t2.user_name, t3.category        
	from jobs t1 
	inner join users t2 on t1.id_worker = t2.id_login
	inner join categories t3 on t1.category_id=t3.id_category
	LIMIT $1 OFFSET $2;`
	listjobsearch = `select COUNT(*) OVER(),t1.id_job,t2.user_name, t3.category        
	from jobs t1 
	inner join users t2 on t1.id_worker = t2.id_login
	inner join categories t3 on t1.category_id=t3.id_category WHERE category ILIKE '%' || $1 || '%'
	LIMIT $2 OFFSET $3;`
	ListFavorite = `SELECT COUNT(*) OVER(),F.id_favorite,P.name,P.photo,C.category,J.wage,J.description
	FROM jobs J
	INNER JOIN categories C ON J.category_id=C.id_category
	INNER JOIN profiles P ON J.id_worker=P.login_id
	INNER JOIN favorites F ON F.job_id=J.id_job
	WHERE F.user_id=$1
	LIMIT $2 OFFSET $3;`
)

type userRepo struct {
	db *sql.DB
}

// SendJobRequest implements interfaces.UserRepository
func (c *userRepo) SendJobRequest(request domain.Request) (int, error) {
	var Id int
	query := `INSERT INTO requests (user_id,job_id,address_id) VALUES($1,$2,$3) RETURNING id_requset;`

	err := c.db.QueryRow(query,
		request.UserId,
		request.JobId,
		request.AddressId,
	).Scan(
		&Id,
	)

	return Id, err
}

// DeleteAddress implements interfaces.UserRepository
func (c *userRepo) DeleteAddress(id int, userid int) error {
	query := `DELETE FROM addresses WHERE id_address=$1 AND user_id=$2 RETURNING id_address;`

	var row int
	sql := c.db.QueryRow(query, id, userid)

	sql.Scan(&row)
	if row == 0 {
		return errors.New("There is no item to delete")
	}

	return sql.Err()
}

// ListAddress implements interfaces.UserRepository
func (c *userRepo) ListAddress(id int) ([]domain.Address, error) {
	var addresses []domain.Address

	query := `select id_address,user_id,house_name,place,city,post,pin,phone from addresses where user_id=$1;`
	rows, err := c.db.Query(query, id)

	if err != nil {
		return addresses, err
	}

	defer rows.Close()

	for rows.Next() {
		var address domain.Address

		err = rows.Scan(
			&address.IdAddress,
			&address.UserId,
			&address.HouseName,
			&address.Place,
			&address.City,
			&address.Post,
			&address.Pin,
			&address.Phone,
		)

		if err != nil {
			return addresses, err
		}

		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		return addresses, err
	}
	return addresses, nil
}

// AddAddress implements interfaces.UserRepository
func (c *userRepo) AddAddress(address domain.Address) (int, error) {
	var Id int
	query := `INSERT INTO addresses (user_id,house_name,place,city,post,pin,phone)
			VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_address;`

	err := c.db.QueryRow(query,
		address.UserId,
		address.HouseName,
		address.Place,
		address.City,
		address.Post,
		address.Pin,
		address.Phone,
	).Scan(
		&Id,
	)

	return Id, err
}

// CheckInFevorite implements interfaces.UserRepository
func (c *userRepo) CheckInFevorite(favorite domain.Favorite) (int, error) {
	query := ` select count(*) from favorites where user_id=$1 and job_id=$2;
	`

	rows, err := c.db.Query(query,
		favorite.UserId,
		favorite.JobId,
	)
	fmt.Println("rows : ", rows)
	if err != nil {
		return 0, err
	}
	var id int

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&id,
		)

		if err != nil {
			return 0, err
		}
	}

	fmt.Println("id : ", id)

	return id, err
}

// ListFevorite implements interfaces.UserRepository
func (c *userRepo) ListFevorite(pagenation utils.Filter, id int) ([]domain.ListFavorite, utils.Metadata, error) {
	var favorites []domain.ListFavorite

	rows, err := c.db.Query(ListFavorite, id, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return favorites, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var favorite domain.ListFavorite

		err = rows.Scan(
			&totalRecords,
			&favorite.FavoriteId,
			&favorite.Name,
			&favorite.Photo,
			&favorite.JobCategory,
			&favorite.Wage,
			&favorite.Description,
		)

		if err != nil {
			return favorites, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		favorites = append(favorites, favorite)
	}
	fmt.Printf("\n\nusers : %v\n\n", favorites)

	if err := rows.Err(); err != nil {
		return favorites, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(favorites)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return favorites, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// AddToFavorite implements interfaces.UserRepository
func (c *userRepo) AddToFavorite(favorite domain.Favorite) (int, error) {
	var Id int
	query := `INSERT INTO favorites (user_id,job_id) VALUES ($1,$2) RETURNING id_favorite;`

	err := c.db.QueryRow(query,
		favorite.UserId,
		favorite.JobId,
	).Scan(
		&Id,
	)

	return Id, err
}

// SearchWorkersWithJob implements interfaces.UserRepository
func (c *userRepo) SearchWorkersWithJob(pagenation utils.Filter, key string) ([]domain.ListJobsWithWorker, utils.Metadata, error) {
	var jobs []domain.ListJobsWithWorker

	rows, err := c.db.Query(listjobsearch, key, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return jobs, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var job domain.ListJobsWithWorker

		err = rows.Scan(
			&totalRecords,
			&job.IdJob,
			&job.WorkerName,
			&job.CategoryName,
		)

		if err != nil {
			return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		jobs = append(jobs, job)
	}
	fmt.Printf("\n\nusers : %v\n\n", jobs)

	if err := rows.Err(); err != nil {
		return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(jobs)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// ListWorkers implements interfaces.UserRepository
func (c *userRepo) ListWorkersWithJob(pagenation utils.Filter) ([]domain.ListJobsWithWorker, utils.Metadata, error) {
	var jobs []domain.ListJobsWithWorker

	rows, err := c.db.Query(listjob, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return jobs, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var job domain.ListJobsWithWorker

		err = rows.Scan(
			&totalRecords,
			&job.IdJob,
			&job.WorkerName,
			&job.CategoryName,
		)

		if err != nil {
			return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		jobs = append(jobs, job)
	}
	fmt.Printf("\n\nusers : %v\n\n", jobs)

	if err := rows.Err(); err != nil {
		return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(jobs)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// // verifyPassword implements interfaces.UserRepository
// func (*userRepo) verifyPassword(changepassword string, id int) (int, error) {
// 	panic("unimplemented")
// }

// ChangePassword implements interfaces.UserRepository
func (c *userRepo) UserChangePassword(changepassword string, id int) (int, error) {
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
func (c *userRepo) UserEditProfile(userProfile domain.Profile, id int) (int, error) {
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

	return Id, err
}

// AddProfile implements interfaces.UserRepository
func (c *userRepo) UserAddProfile(userProfile domain.Profile, id int) (int, error) {
	var Id int
	query := ` INSERT INTO Profiles 
	(login_id,name,gender,date_of_birth,house_name,place,post,pin,contact_number,email_id,photo) 
	VALUES 
	($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING login_id;`

	err := c.db.QueryRow(query,
		id,
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
	).Scan(
		&Id,
	)

	return Id, err
}

// StoreVerificationDetails implements interfaces.UserRepository
func (c *userRepo) StoreVerificationDetails(email string, code int) error {
	query := `INSERT INTO 
		verifications(email, code)
		VALUES( $1, $2);`

	err := c.db.QueryRow(query, email, code).Err()

	return err
}

// VerifyAccount implements interfaces.UserRepository
func (c *userRepo) VerifyAccount(email string, code int) error {
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

	query = `DELETE FROM 
	verifications WHERE 
	email = $1;`
	err = c.db.QueryRow(query, email).Scan(&id)

	return nil
}

// FindUser implements interfaces.UserRepository
func (c *userRepo) FindUser(email string) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `SELECT id_login,user_name,password,verification  FROM users WHERE user_name=$1 AND user_type='user' ;`

	err := c.db.QueryRow(query,
		email).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
		&user.Verification,
	)
	fmt.Printf("\n\n\nuser : %v\n\n\n", user)
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}

	return user, err
}

// InsertUser implements interfaces.UserRepository
func (c *userRepo) InsertUser(login domain.User) (int, error) {
	var id int

	query := `INSERT INTO users (user_name,password,user_type) VALUES ($1,$2,$3) RETURNING id_login;`

	err := c.db.QueryRow(query,
		login.UserName,
		login.Password,
		"user",
	).Scan(
		&id,
	)

	return id, err
}

func NewUserRepo(db *sql.DB) interfaces.UserRepository {
	return &userRepo{
		db: db,
	}
}
