package token

type Token interface {
	Kind() string
	Format() string
}

type Expression []Token
