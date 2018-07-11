package data

import (
	"go-hello-skeleton/models"
	"database/sql"
	"fmt"
)

type petRepo interface {
	putPet(pet models.Pet) error
}

type repo struct {
	db *sql.DB
}

// NewRepo
func NewRepo(db *sql.DB) *repo{
	return &repo{
		db: db,
	}
}

func (p *repo) putPet(pet models.Pet) error {
	stmt, err := p.db.Prepare("INSERT INTO pets (PetName, PetType, Age) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(pet.PetName, pet.PetType, pet.Age)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return fmt.Errorf("1 row was not affected")
	}

	return nil
}
