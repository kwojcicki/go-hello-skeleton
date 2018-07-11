package models

type PetTypes string

const (
	Dog PetTypes = "Dog"
	Cat PetTypes = "Cat"
	Lizard PetTypes = "Lizard"
)

type Pet struct {
	PetName string
	PetType PetTypes
	Age int
}

type PutPetRequest struct {
	Pet Pet
}