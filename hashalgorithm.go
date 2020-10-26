package hashalgoritm

import (
	"github.com/google/uuid"
)

type Algorithm interface {
	ID() uuid.UUID
	HashAndSalt(password string) ([]byte, error)
	Validate(hash []byte, password string) error
}

type Algorithms []Algorithm

func (h Algorithms) FindByID(id uuid.UUID) Algorithm {
	for _, algorithm := range h {
		if id == algorithm.ID() {
			return algorithm
		}
	}
	return nil
}

func LoadAll() Algorithms {
	return []Algorithm{
		Bcrypt14(),
	}
}
