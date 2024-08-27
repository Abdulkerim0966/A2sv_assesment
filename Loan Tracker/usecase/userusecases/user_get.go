package userusecase

import (
	"loanTracker/config"
	"loanTracker/domain"
)

func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	user, err := u.UserRepo.GetUserByUsernameorEmail(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) GetUsers(claims *domain.LoginClaims) ([]domain.User, error) {
	user, err := u.UserRepo.GetUserByUsernameorEmail(claims.Username)
	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		return nil, config.ErrOnlyAuthorOrAdminDel
	}

	users, err := u.UserRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}


	
