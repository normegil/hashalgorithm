package hashalgorithm

import (
	"errors"
	"fmt"
)

type invalidPasswordError struct {
	Original error
}

func (e invalidPasswordError) Error() string {
	return fmt.Errorf("invalid password: %w", e.Original).Error()
}

func IsInvalidPassword(err error) bool {
	return errors.As(err, &invalidPasswordError{})
}
