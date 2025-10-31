package token

type EOF struct {
}

func (e EOF) Kind() string {
	return "EOF"
}

func (e EOF) Format() string {
	return e.Kind()
}
