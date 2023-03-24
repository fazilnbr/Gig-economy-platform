package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/stretchr/testify/assert"
	// "github.com/fazilnbr/project-workey/pkg/repository"
)

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