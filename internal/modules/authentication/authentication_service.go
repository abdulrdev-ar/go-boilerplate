package authentication

import (
	"github.com/inienam06/go-boilerplate/internal/core/jwt"
	"github.com/inienam06/go-boilerplate/internal/exception"
	"github.com/inienam06/go-boilerplate/internal/modules/user"
	"github.com/inienam06/go-boilerplate/internal/util"
)

type IAuthenticationService interface {
	Login(dto AuthLoginDTO) (string, error)
}

type AuthenticationService struct {
	userRepo user.IUserRepository
}

func InitAuthenticationService(userRepo user.IUserRepository) *AuthenticationService {
	return &AuthenticationService{userRepo: userRepo}
}

// Login authenticates user and returns token
// @Summary Login user
// @Description Authenticates user and returns token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body AuthLoginDTO true "Login request"
// @Router /auth/login [post]
func (s *AuthenticationService) Login(dto AuthLoginDTO) (string, error) {
	user, err := s.userRepo.GetByEmail(dto.Email)
	if err != nil {
		return "", exception.NewInternalException(err.Error())
	}

	if user == nil {
		return "", exception.NewNotFoundException("user not found")
	}

	if match, err := util.VerifyPassword(dto.Password, *user.Password); err != nil || !match {
		return "", exception.NewValidationException("invalid password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", exception.NewInternalException(err.Error())
	}

	return token, nil
}
