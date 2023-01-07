package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
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
func (c *adminRepo) ListNewWorkers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	var workers []domain.UserResponse

	query := `SELECT COUNT(*) OVER(),
			  id_login,
			  user_name,
			  password 		
			  FROM logins 
			  WHERE user_type='worker' 
			  AND verification='true' 
			  AND status='newuser'
			  LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var worker domain.UserResponse

		err = rows.Scan(
			&totalRecords,
			&worker.ID,
			&worker.UserName,
			&worker.Password,
		)

		if err != nil {
			return workers, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		workers = append(workers, worker)
	}
	if err := rows.Err(); err != nil {
		return workers, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(workers)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return workers, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// ListUsers implements interfaces.AdminRepository
func (c *adminRepo) ListWorkers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	var workers []domain.UserResponse

	query := `SELECT COUNT(*) OVER(),
			  id_login,
			  user_name,
			  password 		
			  FROM logins 
			  WHERE user_type='worker' 
			  AND verification='true' 
			  AND status='unblocked'
			  LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var worker domain.UserResponse

		err = rows.Scan(
			&totalRecords,
			&worker.ID,
			&worker.UserName,
			&worker.Password,
		)

		if err != nil {
			return workers, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		workers = append(workers, worker)
	}
	if err := rows.Err(); err != nil {
		return workers, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(workers)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return workers, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

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
func (c *adminRepo) ListBlockedUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	var users []domain.UserResponse

	query := `SELECT COUNT(*) OVER(),
			  id_login,
			  user_name,
			  password 		
			  FROM logins 
			  WHERE user_type='user' 
			  AND verification='true' 
			  AND status='blocked'
			  LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&totalRecords,
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(users)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// ListNewUsers implements interfaces.AdminRepository
func (c *adminRepo) ListNewUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	var users []domain.UserResponse

	query := `SELECT COUNT(*) OVER(),
			  id_login,
			  user_name,
			  password 		
			  FROM logins 
			  WHERE user_type='user' 
			  AND verification='true' 
			  AND status='newuser'
			  LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&totalRecords,
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(users)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// ListUsers implements interfaces.AdminRepository
func (c *adminRepo) ListUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	var users []domain.UserResponse

	query := `SELECT COUNT(*) OVER(),
			  id_login,
			  user_name,
			  password 		
			  FROM logins 
			  WHERE user_type='user' 
			  AND verification='true' 
			  AND status='unblocked'
			  LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())

	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var user domain.UserResponse

		err = rows.Scan(
			&totalRecords,
			&user.ID,
			&user.UserName,
			&user.Password,
		)

		if err != nil {
			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}

		users = append(users, user)
	}
	fmt.Printf("\n\nusers : %v\n\n", users)

	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(users)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
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
