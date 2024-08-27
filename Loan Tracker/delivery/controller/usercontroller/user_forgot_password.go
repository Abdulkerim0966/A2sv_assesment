package usercontroller

import (
	"loanTracker/config"
	"loanTracker/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) ForgotPassword(ctx *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if input.Email == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Email is required",
		})
		return
	}

	if input.NewPassword == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "New password is required",
		})
		return
	}

	err = u.UserUsecase.ForgotPassword(input.Email, input.NewPassword)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, domain.APIResponse{
				Status:  code,
				Message: "Internal server error",
				Error:   "Failed to send password reset token",
			})
			return
		}

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to send password reset token",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Password reset token sent",
	})
}
