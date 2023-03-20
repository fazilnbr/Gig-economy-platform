package usecase

import (
	"fmt"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/repository"
	"github.com/fazilnbr/project-workey/pkg/utils"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var Login = domain.User{
	UserName: utils.RandomString(6),
	Password: fmt.Sprint(utils.RandomInt(10000, 99999)),
}

func TestCreateUser(t *testing.T) {

	gormDB, _,_ := utils.MockGormDB()
	authRepoMock := repository.NewUserRepo(gormDB)
	authService := NewUserService(authRepoMock)

	t.Run("Test normal case service login", func(t *testing.T) {
		err := authService.CreateUser(Login)
		assert.NoError(t, err)
	})

	t.Run("Test exist user case service login", func(t *testing.T) {
		Login.UserName = "sethu"
		err := authService.CreateUser(Login)
		if assert.Error(t, err) {
			assert.Equal(t, "Username already exists", err.Error())
		}
	})
}
