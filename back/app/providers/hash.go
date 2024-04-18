package providers

type HashProvider interface {
	Hash(password string) (string, error)
	Compare(password, hashedPassword string) error
}
