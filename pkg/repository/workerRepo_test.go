package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// Create module test
func TestInsertWorker(t *testing.T) {
	// Create mock DB and mock query
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	mockQuery := "INSERT INTO users \\(user_name,password,user_type\\) VALUES \\(\\$1,\\$2,\\$3\\) RETURNING id_login"

	// Create test user
	testWorker := domain.User{
		UserName:     "testuser",
		Password:     "testpassword",
		UserType:     "testusertype",
		Verification: true,
	}

	// Create expected result
	expectedID := 1

	// Set up mock query to expect query and return expected result
	mock.ExpectQuery(mockQuery).WithArgs(testWorker.UserName, testWorker.Password, "worker").WillReturnRows(sqlmock.NewRows([]string{"id_login"}).AddRow(expectedID))

	// Create user repository with mock DB
	workerRepo := NewWorkerRepo(db)

	// Call InsertUser method
	actualID, actualErr := workerRepo.InsertWorker(testWorker)

	assert.NoError(t, actualErr)
	assert.Equal(t, expectedID, actualID)
}

// Read methord test
func TestFindWorkerWithId(t *testing.T) {
	// define mock data
	mockWorker := domain.WorkerResponse{
		ID:           1,
		UserName:     "test_user",
		Password:     "test_password",
		Verification: false,
	}

	// initialize mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// set up mock DB to return predefined result
	mock.ExpectQuery("SELECT id_login,user_name,password FROM users WHERE id_login=\\$1 AND user_type='worker'").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id_login", "user_name", "password"}).
			AddRow(mockWorker.ID, mockWorker.UserName, mockWorker.Password))

	// initialize repository with mock DB
	repo := NewWorkerRepo(db)

	// call FindUserWithId function
	worker, err := repo.FindWorkerWithId(1)

	// assert that the function returns the expected output
	assert.NoError(t, err)
	assert.Equal(t, mockWorker, worker)

	// assert that the QueryRow function of the mock DB was called with the correct query and parameters
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Read all methord
func TestListJobCategoryUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %s", err)
	}
	defer db.Close()

	adminRepo := NewAdminRepo(db)

	pagination := utils.Filter{
		Page:     1,
		PageSize: 10,
	}


	rows := sqlmock.NewRows([]string{"count", "id_category", "user_name"}).
		AddRow(2, 1, "welder").
		AddRow(2, 2, "plumber")

	mock.ExpectQuery("SELECT COUNT").WithArgs( pagination.Limit(), pagination.Offset()).WillReturnRows(rows)

	expectedResult := []domain.Category{
		{
			IdCategory: 1,
			Category: "welder",
		},
		{
			IdCategory: 2,
			Category: "plumber",
		},
	}

	result, metadata, err := adminRepo.ListJobCategory(pagination)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, utils.Metadata{CurrentPage: 1, PageSize: 10, FirstPage: 1, LastPage: 1, TotalRecords: 2}, metadata)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Update methord test

func TestWorkerEditProfile(t *testing.T){
	db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    workerRepo := NewWorkerRepo(db)

    workerProfile := domain.Profile{
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
        workerProfile.Name,
        workerProfile.Gender,
        workerProfile.DateOfBirth,
        workerProfile.HouseName,
        workerProfile.Place,
        workerProfile.Post,
        workerProfile.Pin,
        workerProfile.ContactNumber,
        workerProfile.EmailID,
        workerProfile.Photo,
        expectedId,
    ).WillReturnRows(sqlmock.NewRows([]string{"id_user"}).AddRow(expectedId))

    id, err := workerRepo.WorkerEditProfile(workerProfile,expectedId)

    // Assert the query was called with expected arguments
    assert.NoError(t, mock.ExpectationsWereMet())

    // Assert the function returned the expected values
    assert.Equal(t, expectedId, id)
    assert.NoError(t, err)
}