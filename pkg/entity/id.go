package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

// ParseId parses a string into an ID
func ParseId(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
