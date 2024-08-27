package loancontroller

import (
	"loanTracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)
func (l *LoanController) GetLoanById(ctx *gin.Context) {
	loanId := ctx.Param("id")
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: http.StatusInternalServerError,
			Message: "Internal server error",
			Error: "cannot get claims",
		})
		return
	}
	loan, err := l.loanusecases.GetLoanByID(loanId, claims)
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
		Message: "Success",
		Data: loan,
	})
}	