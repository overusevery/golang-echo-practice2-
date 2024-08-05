package value

import "github.com/google/uuid"

type ID string

func NewID(id string) ID {
	return ID(id)
}

func GenerateNewIDString() string {
	return uuid.NewString()
}
