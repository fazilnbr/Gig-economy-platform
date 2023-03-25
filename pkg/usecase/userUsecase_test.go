package usecase

import (
	"errors"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/domain"
	mock "github.com/fazilnbr/project-workey/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockUserRepository(ctrl)

	userusecase := NewUserService(c)
	testData := []struct {
		name       string
		email      string
		beforeTest func(userRepo *mock.MockUserRepository)
		expectErr  error
	}{
		{
			name:  "Test success response",
			email: "jon",
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUser("jon").Return(domain.UserResponse{
					UserName: "jon",
					Password: "12345",
				}, nil)
			},
			expectErr: nil,
		},
		{
			name:  "Test when user alredy exist response",
			email: "jon",
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUser("jon").Return(domain.UserResponse{}, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(c)
			actualUser, err := userusecase.FindUser(tt.email)
			assert.Equal(t, tt.expectErr, err)
			if err == nil {
				assert.Equal(t, tt.email, actualUser.UserName)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockUserRepository(ctrl)

	userusecase := NewUserService(c)
	testData := []struct {
		name       string
		user       domain.User
		beforeTest func(userRepo *mock.MockUserRepository)
		expectErr  error
	}{
		{
			name: "Test success response",
			user: domain.User{
				UserName: "jon",
				Password: "12345",
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUser("jon").Return(domain.UserResponse{
					UserName: "jon",
					Password: "12345",
				}, errors.New("find user repo error"))
				userRepo.EXPECT().InsertUser(domain.User{
					UserName: "jon",
					Password: "827ccb0eea8a706c4c34a16891f84e7b",
					Verification: false,
				}).Return(1, nil)
			},
			expectErr: nil,
		},
		{
			name: "Test user already exist response",
			user: domain.User{
				UserName: "jon",
				Password: "12345",
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUser("jon").Return(domain.UserResponse{}, nil)
			},
			expectErr: errors.New("Username already exists"),
		},
		{
			name: "Test insert repo error response",
			user: domain.User{
				UserName: "jon",
				Password: "12345",
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUser("jon").Return(domain.UserResponse{
					UserName: "jon",
					Password: "12345",
				}, errors.New("find user repo error"))
				userRepo.EXPECT().InsertUser(domain.User{
					UserName: "jon",
					Password: "827ccb0eea8a706c4c34a16891f84e7b",
				}).Return(1, errors.New("insertuser repo error"))
			},
			expectErr: errors.New("insertuser repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(c)
			 err := userusecase.CreateUser(tt.user)
			assert.Equal(t, tt.expectErr, err)
		})
	}
}
