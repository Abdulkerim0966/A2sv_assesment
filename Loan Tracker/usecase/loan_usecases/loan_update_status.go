package loanusecase

import (
	"errors"
	"loanTracker/domain"
)

func (l *LoanUsecase) UpdateLoanStatus(id string, status string,claim *domain.LoginClaims) (*domain.Loan, error) {
    if claim.Role != "admin" {
		return nil, errors.New("you are not authorized to perform this action")
	}
	loan, err := l.loanRepo.GetLoanById(id)
	if err != nil {
		return nil, err
	}
	err = l.loanRepo.UpdateLoanStatus(id, status)
	if err != nil {
		return nil, err
	}
	loan.LoanStatus = status
	return loan, nil


}

	