package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	LoanAmount float64            `json:"loan_amount" bson:"loan_amount"`
	LoanStatus string             `json:"loan_status" bson:"loan_status"`
	Owner      string             `json:"owner" bson:"owner"`
	RequestedAt time.Time		  `json:"requested_at" bson:"requested_at"`

}

type LoanRepository interface {
	CreateLoan(loan *Loan) (*Loan, error)
	GetLoanById(id string) (*Loan, error)
	UpdateLoanStatus(id string, status string) error
	DeleteLoan(id string) error
	GetAllLoans() ([]*Loan, error)

}


type LoanUsecase interface {
	CreateLoan(loan *Loan) (*Loan, error)
	GetLoanByID(id string,claim *LoginClaims) (*Loan, error)
	UpdateLoanStatus(id string, status string,claim *LoginClaims) (*Loan, error)
	DeleteLoan(id string, claim *LoginClaims) error
	GetAllLoans(claim *LoginClaims) ([]*Loan, error)

}
	//

