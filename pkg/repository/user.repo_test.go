package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var Login = domain.User{
	UserName: utils.RandomString(6),
	Password: fmt.Sprint(utils.RandomInt(10000, 99999)),
}

func TestInsertUser(t *testing.T) {

	t.Run("test normal case repo register", func(t *testing.T) {
		gormDB, mock,_ := utils.MockGormDB()

		query := "INSERT INTO users (user_name,password,user_type) VALUES (?,?,?) RETURNING id_login;"
		mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), Login.UserName, Login.Password, "user").
			WillReturnResult(sqlmock.NewResult(1, 1))

		authRepo := NewUserRepo(gormDB)
		id, err := authRepo.InsertUser(Login)

		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
			assert.NotEqual(t, 0, id)

		})
	})

}

func TestFindUser(t *testing.T) {
	hashedPassword := "12345"

	t.Run("test normal case repo login", func(t *testing.T) {
		gormDB, mock,_ := utils.MockGormDB()

		rows := sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword)
		mock.ExpectQuery("SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1").
			WillReturnRows(rows)

		authRepo := NewUserRepo(gormDB)
		Login.UserName = "sethu"
		user, err := authRepo.FindUser(Login.UserName)
		assert.NoError(t, err)

		t.Run("test get stored password by username is hashed", func(t *testing.T) {
			assert.Equal(t, hashedPassword, user.Password)
		})
		t.Run("test return the value", func(t *testing.T) {
			assert.NotEmpty(t, user)
			assert.Equal(t, "sethu", user.UserName)
		})

	})
}
