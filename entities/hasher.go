package entities

type Hasher interface {
	HashPassword(pass string) (string, error)
	Equal(pass string, hashedPass string) (bool, error)
}

