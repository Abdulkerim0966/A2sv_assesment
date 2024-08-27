package loanusecase

import (
	"loanTracker/config"
	"loanTracker/domain"
)

func (l *LoanUsecase) GetLoanByID(id string,claim *domain.LoginClaims) (*domain.Loan, error) {


	loan , err := l.loanRepo.GetLoanById(id)
	if loan.Owner != claim.Username {
		return nil, config.ErrOnlyAuthorOrAdminDel
	}
	
	if err != nil {
		return nil, err
	}

	return loan, nil
}