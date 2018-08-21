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

	// getPet
	testCases := []struct {
		input  []string
		result int
	}{
		{input: []string{"Krystian,Dog,12"}, result: 1},
		{input: []string{""}, result: 0},
		{input: []string{"Krystian,Dog,12",
			"Tom,Dog,12",
			"Bob,Dog,12",
			"Michael,Dog,12",
			"Jimmy,Dog,12",
			"Timmy,Dog,12",
			"Alice,Dog,12",
			"Pat,Dog,12",
			"Tara,Dog,12",
			"Alex,Dog,12"}, result: 10},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input %s", tc.input), func(t *testing.T) {
			d.mock.ExpectPrepare("SELECT *")

			var returnRows *sqlmock.Rows
			returnRows = sqlmock.NewRows([]string{"petName", "petType", "age"})
			for _, row := range tc.input {
				returnRows = returnRows.FromCSVString(row)
			}
			d.mock.ExpectQuery("SELECT *").
				WillReturnRows(returnRows)

			pets, err := d.repo.getPets()

			assert.Equal(t, tc.result, len(pets))

			if err != nil {
				t.Errorf("Expected err to be nil but got: %s", err)
			}
		})
	}

	// testPutPet
}
