package loancontroller

import (
	"loanTracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *LoanController) DeleteLoan(ctx *gin.Context) {
	loanId := ctx.Param("id")
	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: http.StatusInternalServerError,
			Message: "Internal server error",
			Error: "cannot get claims",
		})
		return
	}
	err := l.loanusecases.DeleteLoan(loanId, claim)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: http.StatusInternalServerError,
			Message: "Internal server error",
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status: http.StatusOK,
		Message: "Loan deleted successfully",
	})
}
