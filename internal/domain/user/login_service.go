package user

import "context"

type LoginService struct {
	repo Repository
}

func NewLoginService(repo Repository) *LoginService {
	return &LoginService{repo: repo}
}

func (s *LoginService) Handle(ctx context.Context) (*User, error) {
	return nil, nil
}
