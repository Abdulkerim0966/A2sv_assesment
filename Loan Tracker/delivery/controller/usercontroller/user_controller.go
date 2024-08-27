package usercontroller

import "loanTracker/domain"

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(uu domain.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: uu,
	}
}
