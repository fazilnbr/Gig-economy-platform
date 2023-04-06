package usecase

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/domain"
	mock "github.com/fazilnbr/project-workey/pkg/mock/repoMock"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRejectJobRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		requestId  int
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name:      "test with  return error",
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().RejectJobRequest(1).Return(errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
		{
			name:      "test with  return no error",
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().RejectJobRequest(1).Return(nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			err := workerUsecase.RejectJobRequest(tt.requestId)
			assert.Equal(t, tt.expectErr, err)

		})
	}
}

func TestAcceptJobRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		requestId  int
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name:      "test with  return error",
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().AcceptJobRequest(1).Return(errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
		{
			name:      "test with  return no error",
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().AcceptJobRequest(1).Return(nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			err := workerUsecase.AcceptJobRequest(tt.requestId)
			assert.Equal(t, tt.expectErr, err)

		})
	}
}

func TestListAcceptedJobRequsetFromUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name            string
		pagenation      utils.Filter
		requestId       int
		beforeTest      func(workerRepo *mock.MockWorkerRepository)
		expectRequest   *[]domain.RequestResponse
		expectMeatadata *utils.Metadata
		expectErr       error
	}{
		{
			name: "test fails repo return with error",
			pagenation: utils.Filter{
				Page:     1,
				PageSize: 5,
			},
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ListAcceptedJobRequsetFromUser(utils.Filter{
					Page:     1,
					PageSize: 5,
				}, 1).Return([]domain.RequestResponse{}, utils.Metadata{}, errors.New("repo error"))
			},
			expectRequest:   &[]domain.RequestResponse{},
			expectMeatadata: &utils.Metadata{},
			expectErr:       errors.New("repo error"),
		},
		{
			name: "test success repo return without error",
			pagenation: utils.Filter{
				Page:     1,
				PageSize: 5,
			},
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ListAcceptedJobRequsetFromUser(utils.Filter{
					Page:     1,
					PageSize: 5,
				}, 1).Return([]domain.RequestResponse{}, utils.Metadata{}, nil)
			},
			expectRequest:   &[]domain.RequestResponse{},
			expectMeatadata: &utils.Metadata{},
			expectErr:       nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualRequests, actualMetadata, actualErr := workerUsecase.ListAcceptedJobRequsetFromUser(tt.pagenation, tt.requestId)
			assert.Equal(t, tt.expectErr, actualErr)
			assert.Equal(t, tt.expectRequest, actualRequests)
			assert.Equal(t, tt.expectMeatadata, actualMetadata)

		})
	}
}

func TestListPendingJobRequsetFromUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name            string
		pagenation      utils.Filter
		requestId       int
		beforeTest      func(workerRepo *mock.MockWorkerRepository)
		expectRequest   *[]domain.RequestResponse
		expectMeatadata *utils.Metadata
		expectErr       error
	}{
		{
			name: "test fails repo return with error",
			pagenation: utils.Filter{
				Page:     1,
				PageSize: 5,
			},
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ListPendingJobRequsetFromUser(utils.Filter{
					Page:     1,
					PageSize: 5,
				}, 1).Return([]domain.RequestResponse{}, utils.Metadata{}, errors.New("repo error"))
			},
			expectRequest:   &[]domain.RequestResponse{},
			expectMeatadata: &utils.Metadata{},
			expectErr:       errors.New("repo error"),
		},
		{
			name: "test success repo return without error",
			pagenation: utils.Filter{
				Page:     1,
				PageSize: 5,
			},
			requestId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ListPendingJobRequsetFromUser(utils.Filter{
					Page:     1,
					PageSize: 5,
				}, 1).Return([]domain.RequestResponse{}, utils.Metadata{}, nil)
			},
			expectRequest:   &[]domain.RequestResponse{},
			expectMeatadata: &utils.Metadata{},
			expectErr:       nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualRequests, actualMetadata, actualErr := workerUsecase.ListPendingJobRequsetFromUser(tt.pagenation, tt.requestId)
			assert.Equal(t, tt.expectErr, actualErr)
			assert.Equal(t, tt.expectRequest, actualRequests)
			assert.Equal(t, tt.expectMeatadata, actualMetadata)

		})
	}
}

func TestDeleteJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		jobId      int
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name:  "test fail with repo return error",
			jobId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().DeleteJob(1).Return(errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
		{
			name:  "test success with  return no error",
			jobId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().DeleteJob(1).Return(nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			err := workerUsecase.DeleteJob(tt.jobId)
			assert.Equal(t, tt.expectErr, err)

		})
	}
}

func TestViewJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		jobId      int
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectJobs []domain.WorkerJob
		expectErr  error
	}{
		{
			name:  "test fail with repo return error",
			jobId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ViewJob(1).Return([]domain.WorkerJob{}, errors.New("repo error"))
			},
			expectJobs: []domain.WorkerJob{},
			expectErr:  errors.New("repo error"),
		},
		{
			name:  "test success with  return no error",
			jobId: 1,
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ViewJob(1).Return([]domain.WorkerJob{}, nil)
			},
			expectJobs: []domain.WorkerJob{},
			expectErr:  nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualJobs, actualErr := workerUsecase.ViewJob(tt.jobId)
			assert.Equal(t, tt.expectErr, actualErr)
			assert.Equal(t, tt.expectJobs, actualJobs)

		})
	}
}

func TestAddJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		job        domain.Job
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectId   int
		expectErr  error
	}{
		{
			name: "test fail with repo return error",
			job:  domain.Job{},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().AddJob(domain.Job{}).Return(1, errors.New("repo error"))
			},
			expectId:  1,
			expectErr: errors.New("repo error"),
		},
		{
			name: "test success with  return no error",
			job:  domain.Job{},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().AddJob(domain.Job{}).Return(1, nil)
			},
			expectId:  1,
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualId, actualErr := workerUsecase.AddJob(tt.job)
			assert.Equal(t, tt.expectErr, actualErr)
			assert.Equal(t, tt.expectId, actualId)

		})
	}
}

func TestListJobCategoryUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name            string
		pagenation      utils.Filter
		beforeTest      func(workerRepo *mock.MockWorkerRepository)
		expectRequest   *[]domain.Category
		expectMeatadata utils.Metadata
		expectErr       error
	}{
		{
			name: "test fails repo return with error",
			pagenation: utils.Filter{
				Page:     1,
				PageSize: 5,
			},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ListJobCategoryUser(utils.Filter{
					Page:     1,
					PageSize: 5,
				}).Return([]domain.Category{}, utils.Metadata{}, errors.New("repo error"))
			},
			expectRequest:   &[]domain.Category{},
			expectMeatadata: utils.Metadata{},
			expectErr:       errors.New("repo error"),
		},
		{
			name: "test success repo return without error",
			pagenation: utils.Filter{
				Page:     1,
				PageSize: 5,
			},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().ListJobCategoryUser(utils.Filter{
					Page:     1,
					PageSize: 5,
				}).Return([]domain.Category{}, utils.Metadata{}, nil)
			},
			expectRequest:   &[]domain.Category{},
			expectMeatadata: utils.Metadata{},
			expectErr:       nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualRequests, actualMetadata, actualErr := workerUsecase.ListJobCategoryUser(tt.pagenation)
			assert.Equal(t, tt.expectErr, actualErr)
			assert.Equal(t, tt.expectRequest, actualRequests)
			assert.Equal(t, tt.expectMeatadata, actualMetadata)

		})
	}
}

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
			err := workerUsecase.WorkerVerifyPassword(tt.changepassword, tt.userId)
			assert.Equal(t, tt.expectErr, err)

		})
	}
}

func TestWorkerChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		workerId   int
		password   string
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name:     "test fail with repo return error",
			workerId: 1,
			password: "12345",
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().WorkerChangePassword("827ccb0eea8a706c4c34a16891f84e7b", 1).Return(1, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
		{
			name:     "test success with  return no error",
			workerId: 1,
			password: "12345",
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().WorkerChangePassword("827ccb0eea8a706c4c34a16891f84e7b", 1).Return(1, nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualErr := workerUsecase.WorkerChangePassword(tt.password, tt.workerId)
			assert.Equal(t, tt.expectErr, actualErr)

		})
	}
}

func TestWorkerEditProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		workerId   int
		profile    domain.Profile
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name:     "test fail with repo return error",
			workerId: 1,
			profile:  domain.Profile{},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().WorkerEditProfile(domain.Profile{}, 1).Return(1, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
		{
			name:     "test success with  return no error",
			workerId: 1,
			profile:  domain.Profile{},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().WorkerEditProfile(domain.Profile{}, 1).Return(1, nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualErr := workerUsecase.WorkerEditProfile(tt.profile, tt.workerId)
			assert.Equal(t, tt.expectErr, actualErr)

		})
	}
}

func TestAddProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		workerId   int
		profile    domain.Profile
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name:     "test fail with repo return error",
			workerId: 1,
			profile:  domain.Profile{},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().WorkerAddProfile(domain.Profile{}, 1).Return(1, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
		{
			name:     "test success with  return no error",
			workerId: 1,
			profile:  domain.Profile{},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().WorkerAddProfile(domain.Profile{}, 1).Return(1, nil)
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualErr := workerUsecase.AddProfile(tt.profile, tt.workerId)
			assert.Equal(t, tt.expectErr, actualErr)

		})
	}
}

func TestCreateWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name       string
		user       domain.User
		beforeTest func(workerRepo *mock.MockWorkerRepository)
		expectErr  error
	}{
		{
			name: "Test success response",
			user: domain.User{
				UserName: "jon",
				Password: "12345",
			},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorker("jon").Return(domain.WorkerResponse{
					UserName: "jon",
					Password: "12345",
				}, sql.ErrNoRows)
				workerRepo.EXPECT().InsertWorker(domain.User{
					UserName:     "jon",
					Password:     "827ccb0eea8a706c4c34a16891f84e7b",
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
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorker("jon").Return(domain.WorkerResponse{}, nil)
			},
			expectErr: errors.New("Username Already Exists"),
		},
		{
			name: "Test insert repo error response",
			user: domain.User{
				UserName: "jon",
				Password: "12345",
			},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorker("jon").Return(domain.WorkerResponse{
					UserName: "jon",
					Password: "12345",
				}, sql.ErrNoRows)
				workerRepo.EXPECT().InsertWorker(domain.User{
					UserName: "jon",
					Password: "827ccb0eea8a706c4c34a16891f84e7b",
				}).Return(1, errors.New("insertuser repo error"))
			},
			expectErr: errors.New("insertuser repo error"),
		},
		{
			name: "Test fail other repo error",
			user: domain.User{
				UserName: "jon",
				Password: "12345",
			},
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorker("jon").Return(domain.WorkerResponse{
					UserName: "jon",
					Password: "12345",
				}, errors.New("find repo error"))
			},
			expectErr: errors.New("find repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			err := workerUsecase.CreateUser(tt.user)
			assert.Equal(t, tt.expectErr, err)
		})
	}
}

func TestFindWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	workerRepo := mock.NewMockWorkerRepository(ctrl)
	workerUsecase := NewWorkerService(workerRepo)
	testData := []struct {
		name         string
		email        string
		beforeTest   func(workerRepo *mock.MockWorkerRepository)
		expectWorker *domain.WorkerResponse
		expectErr    error
	}{
		{
			name:  "test fail with repo return error",
			email: "jon",
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorker("jon").Return(domain.WorkerResponse{}, errors.New("repo error"))
			},
			expectWorker: &domain.WorkerResponse{},
			expectErr:    errors.New("repo error"),
		},
		{
			name:  "test success with  return no error",
			email: "jon",
			beforeTest: func(workerRepo *mock.MockWorkerRepository) {
				workerRepo.EXPECT().FindWorker("jon").Return(domain.WorkerResponse{}, nil)
			},
			expectWorker: &domain.WorkerResponse{},
			expectErr:    nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(workerRepo)
			actualWorker, actualErr := workerUsecase.FindWorker(tt.email)
			assert.Equal(t, tt.expectErr, actualErr)
			assert.Equal(t, tt.expectWorker, actualWorker)

		})
	}
}
