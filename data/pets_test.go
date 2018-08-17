package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type TestData struct {
	repo petRepo
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func constructBasicRepo(t *testing.T) *TestData {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error '%s' when opening a stub database connection", err)
	}

	logger := logrus.New()
	logger.Out = ioutil.Discard

	repo := NewRepo(db)

	repoTestData := TestData{
		repo: repo,
		db:   db,
		mock: mock,
	}

	return &repoTestData
}

func (r *TestData) Close() {
	r.db.Close()
}

func TestAddTcaSetting(t *testing.T) {
	d := constructBasicRepo(t)
	defer d.Close()

	//d.mock.ExpectBegin()
	d.mock.ExpectPrepare("SELECT (.+)")
	//d.mock.ExpectExec("COPY (.+)").WithArgs("tca1", 10, 10, 5, 10, true).WillReturnResult(sqlmock.NewResult(1, 1))
	//d.mock.ExpectExec("COPY (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	//d.mock.ExpectCommit()

	pets, err := d.repo.getPets()

	assert.Equal(t, 0, len(pets))

	if err != nil {
		t.Errorf("Expected err to be nil but got: %s", err)
	}
}
