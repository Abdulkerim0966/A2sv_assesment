package loancontroller

import (
	"loanTracker/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *LoanController) UpdateLoanStatus(ctx *gin.Context) {
	idHex := ctx.Param("id")

	id ,err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status: http.StatusBadRequest,
			Message: "Invalid ID",
			Error: err.Error(),
		})
		return
	}

	status := ctx.Param("status")

	if status != "approved" && status != "rejected" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status: http.StatusBadRequest,
			Message: "Invalid status",
		})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: http.StatusInternalServerError,
			Message: "Internal server error",
			Error: "cannot get claims",
		})
		return
	}


	loanData, err := uc.loanusecases.UpdateLoanStatus(id.Hex(), status, claim)
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
		Message: "Loan status updated",
		Data: loanData,
	})
}