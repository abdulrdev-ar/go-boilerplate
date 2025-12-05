package profile

import (
	"github.com/inienam06/go-boilerplate/internal/exception"
	"github.com/inienam06/go-boilerplate/internal/model"
	"github.com/inienam06/go-boilerplate/internal/modules/user"
)

type IProfile interface {
	GetProfile(userID uint) (*model.User, error)
}

type ProfileService struct {
	userRepo user.IUserRepository
}

func NewProfileService(userRepo user.IUserRepository) *ProfileService {
	return &ProfileService{
		userRepo: userRepo,
	}
}

func (s *ProfileService) GetProfile(userID uint) (*model.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, exception.NewNotFoundException("user not found")
	}
	return user, nil
}
