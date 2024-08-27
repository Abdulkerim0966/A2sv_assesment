package loancontroller

import (
	"loanTracker/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func (l *LoanController) CreateLoan(ctx *gin.Context) {
	var loan struct {
		LoanAmount float64 `json:"amount"`
	}
	err := ctx.ShouldBind(&loan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status: http.StatusBadRequest,
			
			Message: "Invalid request",
			Error: err.Error(),
		})
		return

	}
	if loan.LoanAmount == 0 || loan.LoanAmount < 0  {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status: http.StatusBadRequest,
			Message: "Loan amount is required",
		})
		return
	}

	claim ,ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: http.StatusInternalServerError,
			Message: "Internal server error",
			Error: "cannot get claims",
		})
		return
	}

	loanData := &domain.Loan{
		LoanAmount: loan.LoanAmount,
		Owner: claim.Username,
		LoanStatus: "pending",
		RequestedAt: time.Now(),
	}

	newloan,err := l.loanusecases.CreateLoan(loanData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: http.StatusInternalServerError,
			Message: "Internal server error",
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Status: http.StatusCreated,
		Message: "Loan created",
		Data: newloan,
	})
}





