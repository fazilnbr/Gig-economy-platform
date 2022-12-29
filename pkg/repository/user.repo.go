package repository

import (
	"database/sql"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) interfaces.UserRepository {
	return &userRepo{
		db: db,
	}
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
