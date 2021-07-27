package repositories

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-pawoon/app/entity"
	"github.com/coroo/go-pawoon/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *TransactionRepositoryTestSuite) TestBuildNewTransactionRepository() {
	repoTest := NewTransactionRepository()
	var dummyImpl *TransactionRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *TransactionRepositoryTestSuite) TestTransactionCreate() {
	repoTest := NewTransactionRepository()
	dummyTransaction := entity.Transaction{
		ID				: 1,
		UserId			: 1,
		DeviceTimestamp	: time.Now(),
		TotalAmount		: 100000,
		PaidAmount		: 110000,
		ChangeAmount	: 10000,
		PaymentMethod	: "cash",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}
	_, err := repoTest.SaveTransaction(dummyTransaction)
	assert.Nil(suite.T(), err)
}

func (suite *TransactionRepositoryTestSuite) TestTransactionUpdate() {
	repoTest := NewTransactionRepository()
	dummyTransaction := entity.Transaction{
		ID				: 1,
		UserId			: 1,
		DeviceTimestamp	: time.Now(),
		TotalAmount		: 130000,
		PaidAmount		: 140000,
		ChangeAmount	: 10000,
		PaymentMethod	: "cash",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}
	transactionDummy := repoTest.UpdateTransaction(dummyTransaction)
	assert.Nil(suite.T(), transactionDummy)
}

func (suite *TransactionRepositoryTestSuite) TestTransactionDelete() {
	repoTest := NewTransactionRepository()
	dummyTransaction := entity.Transaction{
		ID: 1,
	}
	transactionDummy := repoTest.DeleteTransaction(dummyTransaction)
	assert.Nil(suite.T(), transactionDummy)
}

func (suite *TransactionRepositoryTestSuite) TestGetAllTransactions() {
	repoTest := NewTransactionRepository()
	transactionDummy := repoTest.GetAllTransactions()
	assert.NotNil(suite.T(), transactionDummy)
}

func (suite *TransactionRepositoryTestSuite) TestGetTransaction() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	repoTest := NewTransactionRepository()
	transactionDummy := repoTest.GetTransaction(c)
	assert.NotNil(suite.T(), transactionDummy)
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}
