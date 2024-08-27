package loanusecase

import "loanTracker/domain"

func (l *LoanUsecase) CreateLoan(loan *domain.Loan) (*domain.Loan ,error) {
	newloan,err := l.loanRepo.CreateLoan(loan)
	if err != nil {
		return nil, err

	}
	return newloan, nil

}