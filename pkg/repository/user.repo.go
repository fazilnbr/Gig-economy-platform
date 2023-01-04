package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type userRepo struct {
	db *sql.DB
}

// EditProfile implements interfaces.UserRepository
func (c *userRepo) EditProfile(userProfile domain.Profile, id int) (int, error) {
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

	return Id, err
}

// AddProfile implements interfaces.UserRepository
func (c *userRepo) AddProfile(userProfile domain.Profile, id int) (int, error) {
	var Id int
	query := ` INSERT INTO Profiles 
	(id_login,name,gender,date_of_birth,house_name,place,post,pin,contact_number,email_id,photo) 
	VALUES 
	($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id_login;`

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

// FindUser implements interfaces.UserRepository
func (c *userRepo) FindUser(email string) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `SELECT id_login,user_name,password  FROM logins WHERE user_name=$1 AND user_type='user' ;`

	err := c.db.QueryRow(query,
		email).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)
	fmt.Printf("\n\n\nuser : %v\n\n\n", user)
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}

	return user, err
}

// InsertUser implements interfaces.UserRepository
func (c *userRepo) InsertUser(login domain.Login) (int, error) {
	var id int
	fmt.Printf("\n\n\ninsert : %v\n\n\n", login)

	query := `INSERT INTO logins (user_name,password,user_type) VALUES ($1,$2,$3) RETURNING id_login;`

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
