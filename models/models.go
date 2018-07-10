package models

type PetTypes string

const (
	Dog PetTypes = "Dog"
	Cat PetTypes = "Cat"
	Lizard PetTypes = "Lizard"
)

type PutPetRequest struct {
	PetName string
	PetType PetTypes
}