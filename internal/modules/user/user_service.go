package user

import (
	"errors"

	"github.com/inienam06/go-boilerplate/internal/exception"
	"github.com/inienam06/go-boilerplate/internal/model"
	"github.com/inienam06/go-boilerplate/internal/util"
)

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	repo IUserRepository
}

func InitUserService(repo IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(dto *CreateUserDTO) (*model.User, error) {

	getUser, err := s.repo.GetByEmail(dto.Email)
	if err != nil {
		return nil, exception.NewInternalException(err.Error())
	}

	if getUser != nil {
		return nil, exception.NewValidationException("email already exists")
	}

	hash, err := util.HashPassword(dto.Password)
	if err != nil {
		return nil, exception.NewInternalException(err.Error())
	}
	payload := &model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: &hash,
	}

	user, err := s.repo.Create(payload)

	if err != nil {
		return nil, exception.NewInternalException(err.Error())
	}

	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return nil, exception.NewInternalException(err.Error())
	}
	if u == nil {
		return nil, exception.NewNotFoundException("user not found")
	}
	return u, nil
}

func (s *UserService) ListUsers(pagination util.Pagination) (*util.Pagination, error) {
	users, err := s.repo.List(pagination)
	if err != nil {
		return nil, exception.NewInternalException(err.Error())
	}
	return users, nil
}
