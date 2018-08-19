package data

import (
	"database/sql"
	"fmt"
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

	// https://blog.golang.org/subtests
	d := constructBasicRepo(t)
	defer d.Close()

	testCases := []struct {
		input  []string
		result int
	}{
		{input: "Krystian,Dog,12", result: 1},
		{input: "", result: 0},
		{input: "08:08", result: 10},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input %s", tc.input), func(t *testing.T) {
			d.mock.ExpectPrepare("SELECT *")
			d.mock.ExpectQuery("SELECT *").
				WillReturnRows(sqlmock.NewRows([]string{"petName", "petType", "age"}).FromCSVString(tc.input))

			pets, err := d.repo.getPets()

			assert.Equal(t, tc.result, len(pets))

			if err != nil {
				t.Errorf("Expected err to be nil but got: %s", err)
			}
		})
	}

	//t.Run("Test No Pets", func(t *testing.T) {
	//
	//	d.mock.ExpectPrepare("SELECT *")
	//	d.mock.ExpectQuery("SELECT *").
	//		WillReturnRows(sqlmock.NewRows([]string{"petName", "petType", "age"}))
	//
	//	pets, err := d.repo.getPets()
	//
	//	assert.Equal(t, 0, len(pets))
	//
	//	if err != nil {
	//		t.Errorf("Expected err to be nil but got: %s", err)
	//	}
	//})
	//
	//t.Run("Test Get Pets", func(t *testing.T) {
	//
	//	d.mock.ExpectPrepare("SELECT *")
	//	d.mock.ExpectQuery("SELECT *").
	//		WillReturnRows(sqlmock.NewRows([]string{"petName", "petType", "age"}).
	//			FromCSVString("Krystian,Dog,12"))
	//
	//	pets, err := d.repo.getPets()
	//
	//	assert.Equal(t, 1, len(pets))
	//
	//	if err != nil {
	//		t.Errorf("Expected err to be nil but got: %s", err)
	//	}
	//})

}
