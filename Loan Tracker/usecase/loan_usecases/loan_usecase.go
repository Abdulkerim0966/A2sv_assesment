package loanusecase

import "loanTracker/domain"

type LoanUsecase struct {
	loanRepo domain.LoanRepository
}

func NewLoanUsecase(lr domain.LoanRepository) domain.LoanUsecase {
	return &LoanUsecase{
		loanRepo: lr,
	}
}