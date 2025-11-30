package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/inienam06/go-boilerplate/internal/exception"
	"github.com/inienam06/go-boilerplate/internal/response"
)

type AuthenticationController struct {
	svc      *AuthenticationService
	validate *validator.Validate
}

func InitAuthenticationController(svc *AuthenticationService) *AuthenticationController {
	return &AuthenticationController{svc: svc, validate: validator.New()}
}

func (c *AuthenticationController) Login(ctx *gin.Context) {
	var request AuthLoginDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(exception.NewValidationException(err.Error()))
		return
	}

	if err := c.validate.Struct(request); err != nil {
		ctx.Error(exception.NewValidationException(err.Error()))
		return
	}

	dto := AuthLoginDTO{
		Email:    request.Email,
		Password: request.Password,
	}

	token, err := c.svc.Login(dto)
	if err != nil {
		ctx.Error(exception.NewUnauthorizedException(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{Message: "login success", Data: token})
}
