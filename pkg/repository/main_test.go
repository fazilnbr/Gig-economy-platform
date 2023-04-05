package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

const (
	dbURI = "postgresql://developer:dev123@localhost:5432/workey?sslmode=disable"
)

var testuserrepo interfaces.UserRepository

func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalln("Cannot connect database : ", err)
	}

	testuserrepo = NewUserRepo(db)

	os.Exit(m.Run())
}
