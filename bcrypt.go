package hashalgorithm

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	Identifier uuid.UUID
	Cost       int
}

func (b Bcrypt) ID() uuid.UUID {
	return b.Identifier
}

func (b Bcrypt) HashAndSalt(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.Cost)
	if err != nil {
		return nil, fmt.Errorf("hash password with cost '%d': %w", b.Cost, err)
	}
	return hash, nil
}

func (b Bcrypt) Validate(hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if nil != err {
		return invalidPasswordError{Original: err}
	}
	return nil
}

func Bcrypt14() Bcrypt {
	return Bcrypt{
		Identifier: uuid.MustParse("01d1de6c-fa67-4caa-84da-684dc5640626"),
		Cost:       14,
	}
}
