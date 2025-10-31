package token

type Token interface {
	Kind() string
	Format() string
}
