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

// AddProfile implements interfaces.UserRepository
func (*userRepo) AddProfile(login domain.User, id int) (int, error) {
	panic("unimplemented")
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
