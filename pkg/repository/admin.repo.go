package repository

import (
	"database/sql"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type adminRepo struct {
	db *sql.DB
}

// AddJobCategory implements interfaces.AdminRepository
func (c *adminRepo) AddJobCategory(category string) error {
	var id int
	fmt.Printf("\n\n\ninsert : %v\n\n\n", category)

	query := `INSERT INTO categories (category) VALUES ($1) RETURNING id_category;`

	err := c.db.QueryRow(query,
		category,
	).Scan(
		&id,
	)

	return err
}

// ActivateUser implements interfaces.AdminRepository
func (c *adminRepo) ActivateWorker(id int) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `UPDATE logins SET status = 'unblocked' WHERE id_login=$1  RETURNING id_login,user_name,password;`

	err := c.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)

	return user, err
}

// BlockUser implements interfaces.AdminRepository
func (c *adminRepo) BlockWorker(id int) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `UPDATE logins SET status = 'blocked' WHERE id_login=$1  RETURNING id_login,user_name,password;`

	err := c.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)

	return user, err
}

// ListBlockedUsers implements interfaces.AdminRepository
func (c *adminRepo) ListBlockedWorkers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT id_login,user_name,password FROM logins WHERE user_type='worker' and verification='true' and status='blocked';`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// fmt.Printf("\n\nlist : %v\n\n", users)
	return users, nil
}

// ListNewUsers implements interfaces.AdminRepository
func (c *adminRepo) ListNewWorkers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT id_login,user_name,password FROM logins WHERE user_type='worker' and verification='true' and status='newuser';`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// fmt.Printf("\n\nlist : %v\n\n", users)
	return users, nil
}

// ListUsers implements interfaces.AdminRepository
func (c *adminRepo) ListWorkers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT id_login,user_name,password FROM logins WHERE user_type='worker' and verification='true' and status='unblocked';`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// fmt.Printf("\n\nlist : %v\n\n", users)
	return users, nil
}

// ActivateUser implements interfaces.AdminRepository
func (c *adminRepo) ActivateUser(id int) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `UPDATE logins SET status = 'unblocked' WHERE id_login=$1  RETURNING id_login,user_name,password;`

	err := c.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)

	return user, err
}

// BlockUser implements interfaces.AdminRepository
func (c *adminRepo) BlockUser(id int) (domain.UserResponse, error) {
	var user domain.UserResponse

	query := `UPDATE logins SET status = 'blocked' WHERE id_login=$1  RETURNING id_login,user_name,password;`

	err := c.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)

	return user, err
}

// FindAdmin implements interfaces.AdminRepository
func (c *adminRepo) FindAdmin(email string) (domain.AdminResponse, error) {
	var admin domain.AdminResponse

	query := `SELECT id_login,user_name,password  FROM logins WHERE user_name=$1 AND user_type='admin';`

	err := c.db.QueryRow(query, email).Scan(
		&admin.ID,
		&admin.Username,
		&admin.Password,
	)

	return admin, err
}

// ListBlockedUsers implements interfaces.AdminRepository
func (c *adminRepo) ListBlockedUsers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT id_login,user_name,password FROM logins WHERE user_type='user' and verification='true' and status='blocked';`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// fmt.Printf("\n\nlist : %v\n\n", users)
	return users, nil
}

// ListNewUsers implements interfaces.AdminRepository
func (c *adminRepo) ListNewUsers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT id_login,user_name,password FROM logins WHERE user_type='user' and verification='true' and status='newuser';`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// fmt.Printf("\n\nlist : %v\n\n", users)
	return users, nil
}

// ListUsers implements interfaces.AdminRepository
func (c *adminRepo) ListUsers() ([]domain.UserResponse, error) {
	var users []domain.UserResponse

	query := `SELECT id_login,user_name,password FROM logins WHERE user_type='user' and verification='true' and status='unblocked';`

	rows, err := c.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// fmt.Printf("\n\nlist : %v\n\n", users)
	return users, nil
}

// StoreVerificationDetails implements interfaces.AdminRepository
func (*adminRepo) StoreVerificationDetails(string, int) error {
	panic("unimplemented")
}

func NewAdminRepo(db *sql.DB) interfaces.AdminRepository {
	return &adminRepo{
		db: db,
	}
}
