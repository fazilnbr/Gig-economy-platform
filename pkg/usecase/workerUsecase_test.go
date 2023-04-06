package usecase

import (
	"errors"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/domain"
	mock "github.com/fazilnbr/project-workey/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestWorkerVerifyPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name           string
		changepassword domain.ChangePassword
		userId         int
		beforeTest     func(workerRepo *mock.MockWorkerRepository)
		expectErr      error
	}{
		{
			name:           "change password user with invalid user",
			changepassword: domain.ChangePassword{},
			userId:         1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorkerWithId(1).Return(domain.WorkerResponse{}, errors.New("Invalid User"))
			},
			expectErr: errors.New("Invalid User"),
		},
		{
			name: "change password user with invalid oldpassword",
			changepassword: domain.ChangePassword{
				OldPassword: "12345",
				NewPassword: "54321",
			},
			userId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorkerWithId(1).Return(domain.WorkerResponse{
					Password: "pwd",
				}, nil)
			},
			expectErr: errors.New("Invalid Password"),
		},
		{
			name: "change password user success",
			changepassword: domain.ChangePassword{
				OldPassword: "12345",
				NewPassword: "54321",
			},
			userId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorkerWithId(1).Return(domain.WorkerResponse{
					Password: "827ccb0eea8a706c4c34a16891f84e7b",
				}, nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			 err := workerUsecase.WorkerVerifyPassword(tt.changepassword,tt.userId)
			assert.Equal(t, tt.expectErr, err)
			
		})
	}
}
