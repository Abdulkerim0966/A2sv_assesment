package userusecase

import "loanTracker/domain"

type UserUsecase struct {
	UserRepo   domain.UserRepository
	Oauth2Repo domain.OAuthStateRepository
}

// CheckUsernameAndEmail implements domain.UserUsecase.
func (u *UserUsecase) CheckUsernameAndEmail(username string, email string) error {
	panic("unimplemented")
}

func NewUserUsecase(ur domain.UserRepository, or domain.OAuthStateRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo:   ur,
		Oauth2Repo: or,
	}
}
