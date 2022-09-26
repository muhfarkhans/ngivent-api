package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(input CreateUserInput, avatarPath string) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	GetUserById(id int) (User, error)
	UpdateUser(input UpdateUserInput) (User, error)
	UpdatePassword(input UpdatePasswordUserInput) (User, error)
	UpdateAvatar(id int, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(input CreateUserInput, avatarPath string) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Avatar = avatarPath
	user.UserType = "member"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on thaht email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateUser(input UpdateUserInput) (User, error) {
	user, err := s.repository.FindById(input.User.Id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdatePassword(input UpdatePasswordUserInput) (User, error) {
	user, err := s.repository.FindById(input.User.Id)
	if err != nil {
		return user, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdateAvatar(id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, err
}
