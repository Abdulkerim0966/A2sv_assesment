package loancontroller

import "loanTracker/domain"

type LoanController struct {
	loanusecases domain.LoanUsecase
}

func NewLoanController(lu domain.LoanUsecase) *LoanController {
	return &LoanController{
		loanusecases: lu,
	}
}

