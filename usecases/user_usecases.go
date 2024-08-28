package usecases

import (
	"errors"

	"github.com/hunderaweke/spher/domain"
	"github.com/hunderaweke/spher/infrastructures"
)

type userUsecases struct {
	userRepository domain.UserRepository
}

func NewUserUsecases(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecases{userRepository: repo}
}

func (uc *userUsecases) Register(u *domain.User) (domain.User, error) {
	hash, err := infrastructures.GenerateHash(u.Password)
	if err != nil {
		return domain.User{}, err
	}
	u.Password = hash
	createdUser, err := uc.userRepository.Create(u)
	if err != nil {
		return domain.User{}, err
	}
	return createdUser, nil
}

func (uc *userUsecases) Login(email, password string) (domain.User, error) {
	user, err := uc.userRepository.FetchByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	if !infrastructures.CompareHashAndPassword(password, user.Password) {
		return domain.User{}, errors.New("incorrect email or password")
	}
	return user, nil
}

func (uc *userUsecases) Update(userID string, data domain.User) (domain.User, error) {
	return uc.userRepository.Update(userID, data)
}

func (uc *userUsecases) Delete(userID string) error {
	return uc.userRepository.Delete(userID)
}

func (uc *userUsecases) FetchByID(userID string) (domain.User, error) {
	return uc.userRepository.FetchByID(userID)
}

func (uc *userUsecases) FetchByEmail(email string) (domain.User, error) {
	return uc.userRepository.FetchByEmail(email)
}

func (uc *userUsecases) FetchByUsername(username string) (domain.User, error) {
	return domain.User{}, nil
}

func (uc *userUsecases) Fetch() ([]domain.User, error) {
	return []domain.User{}, nil
}
