package domain

import "github.com/google/uuid"

type IdGenerator interface {
	Next() uuid.UUID
}

type UUIDGenerator struct{}

func (receiver *UUIDGenerator) Next() uuid.UUID {
	return uuid.New()
}
