package facebook

type Service interface {
	GetMe(token string, params ...string) (*User, error)
}
