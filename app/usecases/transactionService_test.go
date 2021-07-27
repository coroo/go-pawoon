package usecases

import (
	"net/http/httptest"
	"testing"
	"time"

	entity "github.com/coroo/go-pawoon/app/entity"
	repositories "github.com/coroo/go-pawoon/app/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// dummy data
var dummyTransaction = []entity.Transaction{
	entity.Transaction{
		ID				: 1,
		UserId			: 1,
		DeviceTimestamp	: time.Now(),
		TotalAmount		: 100000,
		PaidAmount		: 110000,
		ChangeAmount	: 10000,
		PaymentMethod	: "cash",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, entity.Transaction{
		ID				: 2,
		UserId			: 1,
		DeviceTimestamp	: time.Now(),
		TotalAmount		: 100000,
		PaidAmount		: 110000,
		ChangeAmount	: 10000,
		PaymentMethod	: "cash",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	},
}

type repoMockTransaction struct {
	mock.Mock
}

func (r *repoMockTransaction) SaveTransaction(transaction entity.Transaction) (int, error) {
	return 0, nil
}

func (r *repoMockTransaction) UpdateTransaction(transaction entity.Transaction) error {
	args := r.Called(transaction)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockTransaction) DeleteTransaction(transaction entity.Transaction) error {
	args := r.Called(transaction)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (r *repoMockTransaction) GetAllTransactions() []entity.Transaction {
	return dummyTransaction
}

func (r *repoMockTransaction) GetTransaction(ctx *gin.Context) []entity.Transaction {
	return dummyTransaction
}

func (r *repoMockTransaction) CloseDB() {
}

type TransactionUsecaseTestSuite struct {
	suite.Suite
	repositoryTest repositories.TransactionRepository
}

func (suite *TransactionUsecaseTestSuite) SetupTest() {
	suite.repositoryTest = new(repoMockTransaction)
}

func (suite *TransactionUsecaseTestSuite) TestBuildTransactionService() {
	resultTest := NewTransaction(suite.repositoryTest)
	var dummyImpl *TransactionService
	assert.NotNil(suite.T(), resultTest)
	assert.Implements(suite.T(), dummyImpl, resultTest)
	// assert.NotNil(suite.T(), resultTest.(*TransactionService).repositories)
}

func (suite *TransactionUsecaseTestSuite) TestSaveTransactionUsecase() {
	suite.repositoryTest.(*repoMockTransaction).On("SaveTransaction", dummyTransaction[0]).Return(nil)
	useCaseTest := NewTransaction(suite.repositoryTest)
	// dummyTransaction[0].Password = "Change Password"
	data, _ := useCaseTest.SaveTransaction(dummyTransaction[0])
	assert.NotNil(suite.T(), data)
}

func (suite *TransactionUsecaseTestSuite) TestUpdateTransactionUsecase() {
	suite.repositoryTest.(*repoMockTransaction).On("UpdateTransaction", dummyTransaction[0]).Return(nil)
	useCaseTest := NewTransaction(suite.repositoryTest)
	err := useCaseTest.UpdateTransaction(dummyTransaction[0])
	assert.Nil(suite.T(), err)
}

func (suite *TransactionUsecaseTestSuite) TestDeleteTransactionUsecase() {
	suite.repositoryTest.(*repoMockTransaction).On("DeleteTransaction", dummyTransaction[0]).Return(nil)
	useCaseTest := NewTransaction(suite.repositoryTest)
	err := useCaseTest.DeleteTransaction(dummyTransaction[0])
	assert.Nil(suite.T(), err)
}

func (suite *TransactionUsecaseTestSuite) TestGetAllTransactions() {
	suite.repositoryTest.(*repoMockTransaction).On("GetAllTransactions", dummyTransaction).Return(dummyTransaction)
	useCaseTest := NewTransaction(suite.repositoryTest)
	dummyTransaction := useCaseTest.GetAllTransactions()
	assert.Equal(suite.T(), dummyTransaction, dummyTransaction)
}

func (suite *TransactionUsecaseTestSuite) TestGetTransaction() {
	suite.repositoryTest.(*repoMockTransaction).On("GetTransaction", dummyTransaction[0].ID).Return(dummyTransaction[0], nil)
	useCaseTest := NewTransaction(suite.repositoryTest)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	dummyTransaction := useCaseTest.GetTransaction(c)
	assert.NotNil(suite.T(), dummyTransaction[0])
	assert.Equal(suite.T(), dummyTransaction[0], dummyTransaction[0])
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUsecaseTestSuite))
}
