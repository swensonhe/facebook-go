package mock

import "github.com/swensonhe/facebook-go"

// Service is a mock facebook.Service.
type Service struct {
	GetMeFn      func(token string, params ...string) (*facebook.User, error)
	GetMeInvoked bool
}

// NewService returns a default mock.Service instance.
func NewService() *Service {
	return &Service{
		GetMeFn: func(token string, params ...string) (*facebook.User, error) {
			return &facebook.User{}, nil
		},
	}
}

func (s *Service) GetMe(token string, params ...string) (*facebook.User, error) {
	s.GetMeInvoked = true
	return s.GetMeFn(token, params...)
}
