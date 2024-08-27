package loanusecase

import (
	"loanTracker/config"
	"loanTracker/domain"
)


func (l *LoanUsecase) GetAllLoans(claim *domain.LoginClaims) ([]*domain.Loan, error) {
	if claim.Role != "admin" {
		return nil, config.ErrOnlyAuthorOrAdminDel
	}
	loans, err := l.loanRepo.GetAllLoans()
	if err != nil {
		return nil, err
	}
	return loans, nil
}
