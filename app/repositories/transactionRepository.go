package repositories

import (
	"time"

	entity "github.com/coroo/go-pawoon/app/entity"
	"github.com/coroo/go-pawoon/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type TransactionRepository interface {
	SaveTransaction(transaction entity.Transaction) (int, error)
	UpdateTransaction(transaction entity.Transaction) error
	DeleteTransaction(transaction entity.Transaction) error
	GetAllTransactions() []entity.Transaction
	GetTransaction(ctx *gin.Context) []entity.Transaction
}

type transactionDatabase struct {
	connection *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&entity.Transaction{})
	return &transactionDatabase{
		connection: db,
	}
}

func (db *transactionDatabase) SaveTransaction(transaction entity.Transaction) (int, error) {
	data := &transaction
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := db.connection.Create(data)
	if(err.Error != nil){
		return 0, err.Error
	}
	return data.ID, nil
}

func (db *transactionDatabase) UpdateTransaction(transaction entity.Transaction) error {
	data := &transaction
	data.UpdatedAt = time.Now()
	db.connection.Save(data)
	return nil
}

func (db *transactionDatabase) DeleteTransaction(transaction entity.Transaction) error {
	db.connection.Delete(&transaction)
	return nil
}

func (db *transactionDatabase) GetAllTransactions() []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Set("gorm:auto_preload", true).Find(&transactions)
	return transactions
}

func (db *transactionDatabase) GetTransaction(ctx *gin.Context) []entity.Transaction {
	var transaction []entity.Transaction
	db.connection.Set("gorm:auto_preload", true).Where("id = ?", ctx.Param("id")).First(&transaction)
	return transaction
}
