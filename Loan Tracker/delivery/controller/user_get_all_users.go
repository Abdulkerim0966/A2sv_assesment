package usercontroller

import (
	"loanTracker/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetUsers(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}


	users, err := uc.UserUsecase.GetUsers(claims)
	if err != nil {
		code := http.StatusInternalServerError

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to get user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    users,
	})
}
