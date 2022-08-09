package users

import (
	"gomp/logger"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(RegisterInput) (Users, error)
	LoginUser(LoginInput) (Users, error)
}

type DefaultUserService struct {
	repo UserRepository
}

func NewUserService(repository UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func (s DefaultUserService) CreateUser(input RegisterInput) (Users, error) {
	u := Users{}
	u.Username = input.Username
	u.Role = input.Role
	u.CompanyID = input.CompanyID

	hashPassword, errHash := Hash(input.Password)
	if errHash != nil {
		return u, errHash
	}

	u.Password = string(hashPassword)

	user, err := s.repo.RegisterUser(u)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u DefaultUserService) LoginUser(input LoginInput) (Users, error) {
	username := input.Username
	password := input.Password

	user, err := u.repo.FindUser(username)
	if err != nil {
		return user, err
	}

	if user.Username == "" {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logger.Error("password is not valid" + err.Error())
		return user, err
	}

	return user, nil
}
