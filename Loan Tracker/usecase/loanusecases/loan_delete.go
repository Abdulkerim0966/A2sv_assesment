package loanusecase

import (
	"loanTracker/config"
	"loanTracker/domain"
)


func (l *LoanUsecase) DeleteLoan(id string, claim *domain.LoginClaims) error {
	if claim.Role != "admin" {
		return config.ErrOnlyAuthorOrAdminDel
	}
	
	err := l.loanRepo.DeleteLoan(id)
	if err != nil {
		return err
	}
	return nil
}