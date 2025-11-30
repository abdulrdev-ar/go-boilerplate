package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/inienam06/go-boilerplate/internal/exception"
	"github.com/inienam06/go-boilerplate/internal/response"
	"github.com/inienam06/go-boilerplate/internal/util"
)

type UserController struct {
	svc      *UserService
	validate *validator.Validate
}

func InitUserController(svc *UserService) *UserController {
	return &UserController{svc: svc, validate: validator.New()}
}

// CreateUser godoc
// @Summary Create user
// @Description create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserDTO true "User body"
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req CreateUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewValidationException(err.Error()))
		return
	}

	if err := uc.validate.Struct(req); err != nil {
		c.Error(exception.NewValidationException(err.Error()))
		return
	}

	dto := CreateUserDTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := uc.svc.CreateUser(&dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponse{Message: "success", Status: true, Data: user})
}

// GetUser godoc
// @Summary Get user by id
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Router /users/{id} [get]
func (uc *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	u, err := uc.svc.GetUserByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{Data: u, Message: "success", Status: true})
}

// ListUsers godoc
// @Summary List users
// @Tags users
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Router /users [get]
func (uc *UserController) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	var pagination = util.Pagination{Page: page, Limit: limit}
	users, err := uc.svc.ListUsers(pagination)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response.BaseResponse{Data: users, Message: "success", Status: true})
}
