package data

import (
	"database/sql"
	"fmt"
	"go-hello-skeleton/models"
)

type petRepo interface {
	putPet(pet models.Pet) error
	getPets() ([]models.Pet, error)
}

type repo struct {
	db *sql.DB
}

// NewRepo
func NewRepo(db *sql.DB) *repo {
	return &repo{
		db: db,
	}
}

func (p *repo) getPets() ([]models.Pet, error) {
	stmt, err := p.db.Prepare("SELECT * FROM pets")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	pets := make([]models.Pet, 0)
	for rows.Next() {
		pet := models.Pet{}
		// no need to override global err and do err :=
		err = rows.Scan(&pet.PetName, &pet.PetType, &pet.Age)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (p *repo) putPet(pet models.Pet) error {
	stmt, err := p.db.Prepare("INSERT INTO pets (petName, petType, age) VALUES ($1, $2, $3)")
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
