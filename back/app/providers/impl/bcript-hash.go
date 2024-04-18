package providers_impl

import "golang.org/x/crypto/bcrypt"

type BcryptHashProvider struct{}

func (b BcryptHashProvider) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err
}

func (b BcryptHashProvider) Compare(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
