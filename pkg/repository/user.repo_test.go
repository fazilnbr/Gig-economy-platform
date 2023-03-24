package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/stretchr/testify/assert"
	// "github.com/fazilnbr/project-workey/pkg/repository"
)


// CREATE TEST
func TestInsertUser(t *testing.T) {
	// Create mock DB and mock query
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	mockQuery := "INSERT INTO users \\(user_name,password,user_type,verification\\) VALUES \\(\\$1,\\$2,\\$3,\\$4\\) RETURNING id_login"

	// Create test user
	testUser := domain.User{
		UserName:     "testuser",
		Password:     "testpassword",
		UserType:     "testusertype",
		Verification: true,
	}

	// Create expected result
	expectedID := 1

	// Set up mock query to expect query and return expected result
	mock.ExpectQuery(mockQuery).WithArgs(testUser.UserName, testUser.Password, testUser.UserType, testUser.Verification).WillReturnRows(sqlmock.NewRows([]string{"id_login"}).AddRow(expectedID))

	// Create user repository with mock DB
	userRepo := NewUserRepo(db)

	// Call InsertUser method
	actualID, actualErr := userRepo.InsertUser(testUser)

	assert.NoError(t, actualErr)
	assert.Equal(t, expectedID, actualID)
}
// READ TEST
func TestFindUserWithId(t *testing.T) {
	// define mock data
	mockUser := domain.UserResponse{
		ID:           1,
		UserName:     "test_user",
		Password:     "test_password",
		Verification: true,
	}

	// initialize mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// set up mock DB to return predefined result
	mock.ExpectQuery("SELECT id_login,user_name,password,verification FROM users WHERE id_login=\\$1 AND user_type='user'").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id_login", "user_name", "password", "verification"}).
			AddRow(mockUser.ID, mockUser.UserName, mockUser.Password, mockUser.Verification))

	// initialize repository with mock DB
	repo := &userRepo{db: db}

	// call FindUserWithId function
	user, err := repo.FindUserWithId(1)

	// assert that the function returns the expected output
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	// assert that the QueryRow function of the mock DB was called with the correct query and parameters
	assert.NoError(t, mock.ExpectationsWereMet())
}
// READ ALL TEST
func TestListSendRequests(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %s", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	pagination := utils.Filter{
		Page:     1,
		PageSize: 10,
	}

	id := 1

	rows := sqlmock.NewRows([]string{"count", "id_request", "user_name", "category", "date", "status"}).
		AddRow(2, 1, "john", "programming", "2022-01-01", "approved").
		AddRow(2, 2, "jane", "design", "2022-02-01", "pending")

	mock.ExpectQuery("SELECT COUNT").WithArgs(id, pagination.Limit(), pagination.Offset()).WillReturnRows(rows)

	expectedResult := []domain.RequestUserResponse{
		{
			IdRequest:     1,
			UserName:      "john",
			JobCategory:   "programming",
			JobDate:       "2022-01-01",
			RequestStatus: "approved",
		},
		{
			IdRequest:     2,
			UserName:      "jane",
			JobCategory:   "design",
			JobDate:       "2022-02-01",
			RequestStatus: "pending",
		},
	}

	result, metadata, err := userRepo.ListSendRequests(pagination, id)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, utils.Metadata{CurrentPage: 1, PageSize: 10, FirstPage: 1, LastPage: 1, TotalRecords: 2}, metadata)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// UPDATE TEST
func TestUserEditProfile(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    userRepo := NewUserRepo(db)

    userProfile := domain.Profile{
        Name:           "John Doe",
        Gender:         "Male",
        DateOfBirth:    "1990-01-01",
        HouseName:      "1st Avenue",
        Place:          "New York",
        Post:           "10001",
        Pin:            "123456",
        ContactNumber:  "1234567890",
        EmailID:        "john.doe@example.com",
        Photo:          "photo.jpg",
    }

    expectedId := 1

    // Mocking the query and its expected result
    mock.ExpectQuery("^UPDATE profiles").WithArgs(
        userProfile.Name,
        userProfile.Gender,
        userProfile.DateOfBirth,
        userProfile.HouseName,
        userProfile.Place,
        userProfile.Post,
        userProfile.Pin,
        userProfile.ContactNumber,
        userProfile.EmailID,
        userProfile.Photo,
        expectedId,
    ).WillReturnRows(sqlmock.NewRows([]string{"id_user"}).AddRow(expectedId))

    id, err := userRepo.UserEditProfile(userProfile, expectedId)

    // Assert the query was called with expected arguments
    assert.NoError(t, mock.ExpectationsWereMet())

    // Assert the function returned the expected values
    assert.Equal(t, expectedId, id)
    assert.NoError(t, err)
}


// DELETE TEST
func TestUserRepo_DeleteJobRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	userRepo := NewUserRepo(db)

	tests := []struct {
		name          string
		requestId     int
		userid        int
		mockQueryFunc func()
		expectedErr   error
	}{
		{
			name:      "request found",
			requestId: 1,
			userid:    1,
			mockQueryFunc: func() {
				mock.ExpectQuery("^DELETE FROM requests WHERE id_requset=\\$1 AND user_id=\\$2 RETURNING id_requset;$").
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id_requset"}).AddRow(1))
			},
			expectedErr: nil,
		},
		{
			name:      "no request to cancel",
			requestId: 1,
			userid:    1,
			mockQueryFunc: func() {
				mock.ExpectQuery("^DELETE FROM requests WHERE id_requset=\\$1 AND user_id=\\$2 RETURNING id_requset;$").
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id_requset"}))
			},
			expectedErr: errors.New("There is no request to cancel"),
		},
		{
			name:      "DB error",
			requestId: 1,
			userid:    1,
			mockQueryFunc: func() {
				mock.ExpectQuery("^DELETE FROM requests WHERE id_requset=\\$1 AND user_id=\\$2 RETURNING id_requset;$").
					WithArgs(1, 1).
					WillReturnError(errors.New("DB error"))
			},
			expectedErr: errors.New("DB error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQueryFunc()

			err := userRepo.DeleteJobRequest(tt.requestId, tt.userid)

			assert.Equal(t, tt.expectedErr, err)
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
