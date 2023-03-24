package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

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

// FindUserWithId implements interfaces.UserRepository
func (c *userRepo) FindUserWithId(id int) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `SELECT id_login,user_name,password,verification  FROM users WHERE id_login=$1 AND user_type='user' ;`

	err := c.db.QueryRow(query,
		id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
		&user.Verification,
	)
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}

	return user, err
}

// UpdatePaymentId implements interfaces.UserRepository
func (c *userRepo) UpdatePaymentId(razorPaymentId string, idPayment int) error {
	query := `UPDATE job_payments SET razor_paymet_id=$1,payment_status='completed' WHERE id_payment=$2 RETURNING id_payment;`
	var row int
	sql := c.db.QueryRow(query, razorPaymentId, idPayment)

	sql.Scan(&row)
	if row == 0 {
		return errors.New("There is no accepted job to complition")
	}

	return sql.Err()
}

// CheckOrderId implements interfaces.UserRepository
func (c *userRepo) CheckOrderId(userId int, orderId string) (int, error) {
	var id int

	query := `SELECT id_payment FROM job_payments WHERE order_id=$1 AND user_id=$2;`

	err := c.db.QueryRow(query, strings.Join(strings.Fields(orderId), ""), userId).Scan(
		&id,
	)
	if err == sql.ErrNoRows {
		return id, errors.New("Fake payment order id ")
	}
	if id == 0 {
		return id, errors.New("no data")
	}
	if err != nil && err != sql.ErrNoRows {
		return id, err
	}

	return id, err
}

// SaveOrderId implements interfaces.UserRepository
func (c *userRepo) SavePaymentOrderDeatials(payment domain.JobPayment) (int, error) {
	var Id int
	query := `insert into job_payments (request_id,order_id,user_id,amount,date) values($1,$2,$3,$4,$5) RETURNING id_payment;`

	time := time.Now()
	date := fmt.Sprintf("%v/%v/%v", time.Day(), time.Month(), time.Year())

	err := c.db.QueryRow(query,
		payment.RequestId,
		payment.OrderId,
		payment.UserId,
		payment.Amount,
		date,
	).Scan(
		&Id,
	)

	return Id, err
}

// FetchRazorPayDetials implements interfaces.UserRepository
func (c *userRepo) FetchRazorPayDetials(userId int, requestId int) (domain.RazorPayVariables, error) {
	var razordata domain.RazorPayVariables

	query := `SELECT U.user_name,U.user_name,A.phone,J.wage from users AS U
			INNER JOIN addresses AS A ON A.user_id=id_login 
			INNER JOIN requests AS R ON R.user_id=U.id_login
			INNER JOIN jobs AS J ON id_job=R.job_id 
			WHERE U.id_login=$1 
			AND R.id_requset=$2
			AND R.status='completed';`

	err := c.db.QueryRow(query, userId, requestId).Scan(
		&razordata.Email,
		&razordata.Name,
		&razordata.Contact,
		&razordata.Amount,
	)
	if err == sql.ErrNoRows {
		return razordata, errors.New("There is no job completed to done payment")
	}
	if err != nil && err != sql.ErrNoRows {
		return razordata, err
	}

	return razordata, err
}

// UpdateJobComplition implements interfaces.UserRepository
func (c *userRepo) UpdateJobComplition(userId int, requestId int) error {
	query := `UPDATE requests SET status='completed' WHERE id_requset=$1 AND user_id=$2 AND status='accepted' RETURNING id_requset;`
	var row int
	sql := c.db.QueryRow(query, requestId, userId)

	sql.Scan(&row)
	if row == 0 {
		return errors.New("There is no accepted job to complition")
	}

	return sql.Err()
}

// ViewSendRequest implements interfaces.UserRepository
func (c *userRepo) ViewSendOneRequest(userId int, requestId int) (domain.RequestUserResponse, error) {
	var request domain.RequestUserResponse

	query := `SELECT R.id_requset,U.user_name,C.category,R.date,R.status,A.* FROM requests AS R
				INNER JOIN jobs AS J ON R.job_id=J.id_job 
				INNER JOIN categories AS C ON J.category_id=C.id_category
				INNER JOIN users AS U ON J.id_worker=U.id_login INNER JOIN addresses AS A ON A.user_id=$1 WHERE R.user_id=$1 AND R.id_requset=$2;
	`

	err := c.db.QueryRow(query, userId, requestId).Scan(
		&request.IdRequest,
		&request.UserName,
		&request.JobCategory,
		&request.JobDate,
		&request.RequestStatus,
		&request.Address.IdAddress,
		&request.Address.UserId,
		&request.Address.HouseName,
		&request.Address.Place,
		&request.Address.City,
		&request.Address.Post,
		&request.Address.Pin,
		&request.Address.Phone,
	)
	if err != nil && err != sql.ErrNoRows {
		return request, err
	}

	return request, err
}

// ListSendRequests implements interfaces.UserRepository
func (c *userRepo) ListSendRequests(pagenation utils.Filter, id int) ([]domain.RequestUserResponse, utils.Metadata, error) {
	var requests []domain.RequestUserResponse

	query := `SELECT COUNT(*) OVER(),R.id_requset, U.user_name,C.category,R.date,R.status FROM requests AS R
			INNER JOIN jobs AS J ON R.job_id=J.id_job 
			INNER JOIN categories AS C ON J.category_id=C.id_category
			INNER JOIN users AS U ON J.id_worker=U.id_login WHERE R.user_id=$1 ORDER BY R.date LIMIT $2 OFFSET $3;`
	rows, err := c.db.Query(query, id, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return requests, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var request domain.RequestUserResponse

		err = rows.Scan(
			&totalRecords,
			&request.IdRequest,
			&request.UserName,
			&request.JobCategory,
			&request.JobDate,
			&request.RequestStatus,
		)

		if err != nil {
			return requests, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		requests = append(requests, request)
	}
	if err := rows.Err(); err != nil {
		return requests, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	return requests, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// DeleteJobRequest implements interfaces.UserRepository
func (c *userRepo) DeleteJobRequest(requestId int, userid int) error {
	query := `DELETE FROM requests WHERE id_requset=$1 AND user_id=$2 RETURNING id_requset;`

	var row int
	sql := c.db.QueryRow(query, requestId, userid)

	if sql.Err()!=nil{
		return sql.Err()	
	}
	sql.Scan(&row)
	if row == 0 {
		return errors.New("There is no request to cancel")
	}
	return nil

}

// CheckInRequest implements interfaces.UserRepository
func (c *userRepo) CheckInRequest(request domain.Request) (int, error) {
	query := `SELECT COUNT(*) FROM requests WHERE user_id=$1 AND job_id=$2 AND address_id=$3;`

	rows, err := c.db.Query(query,
		request.UserId,
		request.JobId,
		request.AddressId,
	)
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

	return id, err
}

// SendJobRequest implements interfaces.UserRepository
func (c *userRepo) SendJobRequest(request domain.Request) (int, error) {
	var Id int
	query := `INSERT INTO requests (user_id,job_id,address_id,date) VALUES($1,$2,$3,$4) RETURNING id_requset;`

	time := time.Now()
	date := fmt.Sprintf("%v/%v/%v", time.Day(), time.Month(), time.Year())

	err := c.db.QueryRow(query,
		request.UserId,
		request.JobId,
		request.AddressId,
		date,
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
	query := ` select count(*) from favorites where user_id=$1 and job_id=$2;`

	rows, err := c.db.Query(query,
		favorite.UserId,
		favorite.JobId,
	)
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

	if err := rows.Err(); err != nil {
		return favorites, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
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

	if err := rows.Err(); err != nil {
		return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
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

	if err := rows.Err(); err != nil {
		return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	return jobs, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// ChangePassword implements interfaces.UserRepository
func (c *userRepo) UserChangePassword(changepassword string, id int) (int, error) {
	var Id int
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
func (c *userRepo) StoreVerificationDetails(email string, code string) error {
	query := `INSERT INTO 
		verifications(email, code)
		VALUES( $1, $2);`

	err := c.db.QueryRow(query, email, code).Err()

	return err
}

// VerifyAccount implements interfaces.UserRepository
func (c *userRepo) VerifyAccount(email string, code string) error {
	var id int

	query := `SELECT id FROM 
	verifications WHERE 
	email = $1 AND code = $2;`
	err := c.db.QueryRow(query, email, code).Scan(&id)

	if err == sql.ErrNoRows {
		return errors.New("varification link alredy userd or expierd try again")
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

	// query := `SELECT id_login,user_name,password,verification  FROM users WHERE user_name=$1 AND user_type='user' ;`
	query := `SELECT id_login,user_name,password,verification,user_type  FROM users WHERE user_name=$1;`

	err := c.db.QueryRow(query,
		email).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
		&user.Verification,
		&user.UserType,
	)
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}

	return user, err
}

// InsertUser implements interfaces.UserRepository
func (c *userRepo) InsertUser(login domain.User) (int, error) {
	var id int

	query := `INSERT INTO users (user_name,password,user_type,verification) VALUES ($1,$2,$3,$4) RETURNING id_login;`

	err := c.db.QueryRow(query,
		login.UserName,
		login.Password,
		login.UserType,
		login.Verification,
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
