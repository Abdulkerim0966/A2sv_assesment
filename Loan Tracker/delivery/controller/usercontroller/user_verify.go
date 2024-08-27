package usercontroller

import (
	"loanTracker/config"
	"loanTracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) VerifyUser(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "Token is required",
		})
		return
	}

	err := u.UserUsecase.VerifyUser(token)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to verify user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully verified user",
	})
}
