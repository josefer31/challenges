package errorService

import (
	"fmt"
)

type DescriptionLenError struct {
	value string
}

func (descriptionLenError *DescriptionLenError) Error() string {
	return fmt.Sprintf("Description has invalid len %d", len(descriptionLenError.value))
}

func NewDescriptionLenError(description string) *DescriptionLenError {
	return &DescriptionLenError{description}
}
