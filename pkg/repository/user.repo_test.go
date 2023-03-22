package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/mock"
	"github.com/fazilnbr/project-workey/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var User = domain.User{
	IdLogin:  1,
	UserName: utils.RandomString(6),
	Password: fmt.Sprint(utils.RandomInt(10000, 99999)),
}

func TestInsertUser0(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockRep := mock.NewMockUserRepository(ctrl)

	mockRep.EXPECT().InsertUser(User).Times(1).Return(User.IdLogin, nil)

	id, err := mockRep.InsertUser(User)

	t.Run("normal test case", func(t *testing.T) {
		assert.NoError(t, err, "error found in insert user")
		assert.Equal(t, User.IdLogin, id, "miss match user id")
	})

}

func TestInsertUser(t *testing.T) {

	_, mock, mockDB := utils.MockGormDB()
	t.Run("test normal case repo register", func(t *testing.T) {

		query := "INSERT INTO users (user_name,password,user_type) VALUES (?,?,?) RETURNING id_login;"
		mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), User.UserName, User.Password, "user").
			WillReturnResult(sqlmock.NewResult(1, 1))

		authRepo := NewUserRepo(mockDB)
		id, err := authRepo.InsertUser(User)

		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
			assert.NotEqual(t, 0, id)

		})
	})
}

func TestFindUser(t *testing.T) {
	hashedPassword := "12345"

	t.Run("test normal case repo login", func(t *testing.T) {
		gormDB, mock, _ := utils.MockGormDB()

		rows := sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword)
		mock.ExpectQuery("SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1").
			WillReturnRows(rows)

		authRepo := NewUserRepo(gormDB)
		User.UserName = "sethu"
		user, err := authRepo.FindUser(User.UserName)
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
