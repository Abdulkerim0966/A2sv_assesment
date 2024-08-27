package repository

import (
	"context"
	"loanTracker/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type LoanRepository struct {
	LoanRepository *mongo.Collection
}
func NewLoanRepository(db *mongo.Database) domain.LoanRepository {
	return &LoanRepository{
		LoanRepository: db.Collection("loans"),
	}
}	

func (lr *LoanRepository) CreateLoan(loan *domain.Loan) (*domain.Loan ,error ){
	newloan ,err := lr.LoanRepository.InsertOne(context.Background(), loan)
	if err != nil {

		return nil, err
	}
	loan.ID = newloan.InsertedID.(primitive.ObjectID)
	return loan, nil

}
	

func (lr *LoanRepository) GetLoanById(id string) (*domain.Loan, error) {
	loanid ,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var loan domain.Loan
	err = lr.LoanRepository.FindOne(context.Background(), bson.M{"_id": loanid}).Decode(&loan)
	if err != nil {
		return nil, err
	
}
	return &loan, nil
}

func (lr *LoanRepository) UpdateLoanStatus(id string, status string) error {
	loanid ,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = lr.LoanRepository.UpdateOne(context.Background(), bson.M{"_id": loanid}, bson.M{"$set": bson.M{"loan_status": status}})
	if err != nil {
		return err
	}
	return nil
}									

func (lr *LoanRepository) DeleteLoan(id string) error {
	loanid ,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = lr.LoanRepository.DeleteOne(context.Background(), bson.M{"_id": loanid})
	if err != nil {
		return err
	}
	return nil
}

func (lr *LoanRepository) GetAllLoans() ([]*domain.Loan, error) {
	cursor, err := lr.LoanRepository.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var loans []*domain.Loan
	for cursor.Next(context.Background()) {
		var loan domain.Loan
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}

	return loans, nil
}

