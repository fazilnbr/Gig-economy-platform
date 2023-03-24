package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// Create module test
func TestAddJobCategory(t *testing.T) {
	// Create mock DB and mock query
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	mockQuery := "INSERT INTO categories \\(category\\) VALUES \\(\\$1\\) RETURNING id_category;"

	// Create Category
	Category := "welder"
	expectedId := 1
	mock.ExpectQuery(mockQuery).WithArgs(Category).WillReturnRows(sqlmock.NewRows([]string{"id_category"}).AddRow(expectedId))

	// Create admin repocitory with mock

	adminRepo := NewAdminRepo(db)

	// Call AddJobCategory methord
	err = adminRepo.AddJobCategory(Category)

	assert.NoError(t, err)

	// assert that the QueryRow function of the mock DB was called with the correct query and parameters
	assert.NoError(t, mock.ExpectationsWereMet())

}

// Read module test
func TestFindAdmin(t *testing.T) {
	// define mock data
	mockAdmin := domain.AdminResponse{
		ID:       1,
		Username: "admin@gmail.com",
		Password: "12345",
	}

	// initialize mock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// set up mock DB to return predefined result
	mock.ExpectQuery("SELECT id_login,user_name,password  FROM users WHERE user_name=\\$1 AND user_type='admin'").
		WithArgs(mockAdmin.Username).
		WillReturnRows(sqlmock.NewRows([]string{"id_login", "user_name", "password"}).
			AddRow(mockAdmin.ID, mockAdmin.Username, mockAdmin.Password))

	// initialize repository with mock DB
	repo := &adminRepo{db: db}

	// call FindUserWithId function
	admin, err := repo.FindAdmin(mockAdmin.Username)

	// assert that the function returns the expected output
	assert.NoError(t, err)
	assert.Equal(t, mockAdmin, admin)

	// assert that the QueryRow function of the mock DB was called with the correct query and parameters
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Read All module test
func TestListNewUsers(t *testing.T) {
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


	rows := sqlmock.NewRows([]string{"count", "id_login", "user_name", "password"}).
		AddRow(2, 1, "john", "12345").
		AddRow(2, 2, "jane", "54321")

	mock.ExpectQuery("SELECT COUNT").WithArgs(pagination.Limit(), pagination.Offset()).WillReturnRows(rows)

	expectedResult := []domain.UserResponse{
		{
			ID: 1,
			UserName: "john",
			Password: "12345",
		},
		{
			ID: 2,
			UserName: "jane",
			Password: "54321",
		},
	}

	result, metadata, err := adminRepo.ListNewUsers(pagination)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, utils.Metadata{CurrentPage: 1, PageSize: 10, FirstPage: 1, LastPage: 1, TotalRecords: 2}, metadata)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Update module test
func TestUpdateJobCategory(t *testing.T){
	db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    adminRepo := NewAdminRepo(db)

    category := domain.Category{
		IdCategory: 1,
		Category: "welder",
    }


    // Mocking the query and its expected result
    mock.ExpectQuery("^UPDATE categories").WithArgs(
		category.Category,
		category.IdCategory,
    ).WillReturnRows(sqlmock.NewRows([]string{"id_user"}).AddRow(category.IdCategory))

    id, err := adminRepo.UpdateJobCategory(category)

    // Assert the query was called with expected arguments
    assert.NoError(t, mock.ExpectationsWereMet())

    // Assert the function returned the expected values
    assert.Equal(t, category.IdCategory, id)
    assert.NoError(t, err)
}

// Delete module test

