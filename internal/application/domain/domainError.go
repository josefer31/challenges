package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type AdNotFoundError struct {
	id uuid.UUID
}

func (adNotFoundError AdNotFoundError) Error() string {
	return fmt.Sprintf("ad %v not found", adNotFoundError.id.String())
}

func NewAdNotFoundError(id uuid.UUID) AdNotFoundError {
	return AdNotFoundError{id}
}
