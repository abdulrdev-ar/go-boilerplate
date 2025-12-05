package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inienam06/go-boilerplate/internal/response"
)

type ProfileController struct {
	svc *ProfileService
}

func InitProfileController(svc *ProfileService) *ProfileController {
	return &ProfileController{svc: svc}
}

// GetProfile godoc
// @Summary Get profile
// @Tags profile
// @Produce json
// @Router /profile [get]
// @Security Bearer
func (c *ProfileController) GetProfile(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	user, err := c.svc.GetProfile(userID.(uint))

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Data:    user,
		Message: "success",
		Status:  true,
	})
}
