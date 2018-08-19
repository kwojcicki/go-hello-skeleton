package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

func TestPetRepo(t *testing.T) {

	d := constructBasicRepo(t)
	defer d.Close()

	t.Run("Test No Pets", func(t *testing.T) {

		d.mock.ExpectPrepare("SELECT *")
		d.mock.ExpectQuery("SELECT *").
			WillReturnRows(sqlmock.NewRows([]string{"petName", "petType", "age"}))

		pets, err := d.repo.getPets()

		assert.Equal(t, 0, len(pets))

		if err != nil {
			t.Errorf("Expected err to be nil but got: %s", err)
		}
	})

	t.Run("Test Get Pets", func(t *testing.T) {

		d.mock.ExpectPrepare("SELECT *")
		d.mock.ExpectQuery("SELECT *").
			WillReturnRows(sqlmock.NewRows([]string{"petName", "petType", "age"}).
				FromCSVString("Krystian,Dog,12"))

		pets, err := d.repo.getPets()

		assert.Equal(t, 1, len(pets))

		if err != nil {
			t.Errorf("Expected err to be nil but got: %s", err)
		}
	})

}
