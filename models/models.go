package models

type PetTypes string

const (
	Dog    PetTypes = "Dog"
	Cat    PetTypes = "Cat"
	Lizard PetTypes = "Lizard"
)

type Pet struct {
	// camel case for json and db
	PetName string   `json:"petName" db:"petName"`
	PetType PetTypes `json:"petType" db:"petType"`
	Age     int      `json:"age" db:"age"`
}

type PutPetRequest struct {
	Pet Pet
}
