package usecase

import (
	"errors"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/mock"
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
