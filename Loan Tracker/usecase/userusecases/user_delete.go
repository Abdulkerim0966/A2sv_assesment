package userusecase

import (
	"loanTracker/config"
	"loanTracker/domain"
)

func (u *UserUsecase) DeleteUser(username string, claim *domain.LoginClaims) error {
	user, err := u.UserRepo.GetUserByUsernameorEmail(claim.Username)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		return config.ErrOnlyAuthorOrAdminDel
	}

	return u.UserRepo.DeleteUser(username)
}