package usecases

import (
	"github.com/hunderaweke/spher/domain"
)

type userUsecases struct {
	userRepository domain.UserRepository
}

func (uc *userUsecases) Register(u *domain.User) (domain.User, error) {
	return domain.User{}, nil
}
