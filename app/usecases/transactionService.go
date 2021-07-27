package usecases

import (
	entity "github.com/coroo/go-pawoon/app/entity"
	repositories "github.com/coroo/go-pawoon/app/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type TransactionService interface {
	SaveTransaction(entity.Transaction, int) (int, error)
	UpdateTransaction(entity.Transaction) error
	DeleteTransaction(entity.Transaction) error
	GetAllTransactions() []entity.Transaction
	GetTransaction(ctx *gin.Context) []entity.Transaction
}

type transactionService struct {
	repositories repositories.TransactionRepository
}

func NewTransaction(transactionRepository repositories.TransactionRepository) TransactionService {
	return &transactionService{
		repositories: transactionRepository,
	}
}

func (usecases *transactionService) GetAllTransactions() []entity.Transaction {
	return usecases.repositories.GetAllTransactions()
}

func (usecases *transactionService) GetTransaction(ctx *gin.Context) []entity.Transaction {
	return usecases.repositories.GetTransaction(ctx)
}

func (usecases *transactionService) SaveTransaction(transaction entity.Transaction, userId int) (int, error) {
	transaction.Uuid = uuid.New().String()
	transaction.UserId = userId
	return usecases.repositories.SaveTransaction(transaction)
}

func (usecases *transactionService) UpdateTransaction(transaction entity.Transaction) error {
	return usecases.repositories.UpdateTransaction(transaction)
}

func (usecases *transactionService) DeleteTransaction(transaction entity.Transaction) error {
	return usecases.repositories.DeleteTransaction(transaction)
}