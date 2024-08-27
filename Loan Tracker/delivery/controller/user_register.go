package usercontroller

import (
	"loanTracker/config"
	"loanTracker/domain"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *UserController) RegisterUser(ctx *gin.Context) {
	var userData struct {
		FirstName string `form:"firstname"`
		LastName  string `form:"lastname"`
		Bio       string `form:"bio"`
		Username  string `form:"username"`
		Password  string `form:"password"`
		Email     string `form:"email"`
		Roles     string `form:"roles"`
		Address   string `form:"address"`
	
	}

	err := ctx.ShouldBind(&userData)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	log.Println("User data:", userData)

	if userData.Username == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "username is required",
		})
		return
	}

	if userData.Email == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "email is required",
		})
		return
	}

	if userData.Password == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "password is required",
		})
		return
	}

	user := &domain.User{
		FirstName:  userData.FirstName,
		LastName:   userData.LastName,
		Username:   userData.Username,
		Password:   userData.Password,
		Email:      userData.Email,
		Bio:        userData.Bio,
		Address:    userData.Address,
		Role:       userData.Roles,
		JoinedDate: time.Now(),
	}



	err = u.UserUsecase.RegisterUser(user)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to register user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Verification email has been sent",
	})
}
