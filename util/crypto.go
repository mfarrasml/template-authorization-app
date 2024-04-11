package util

import "golang.org/x/crypto/bcrypt"

type PasswordUtil interface {
	HashPassword(password string) ([]byte, error)
	ComparePwdWithHash(password string, hashed []byte) error
}

type bcryptHasherUtil struct {
	cost int
}

func NewBcryptHasherUtil(cost int) *bcryptHasherUtil {
	return &bcryptHasherUtil{
		cost: cost,
	}
}

func (h *bcryptHasherUtil) HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return nil, err
	}
	return hashed, nil
}

func (h *bcryptHasherUtil) ComparePwdWithHash(password string, hashed []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
