package deliveries

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	usecases "github.com/coroo/go-pawoon/app/usecases"
	entity "github.com/coroo/go-pawoon/app/entity"

	"github.com/gin-gonic/gin"
)

// dummy data
var dummyTransaction = []*entity.Transaction{
	&entity.Transaction{
		ID				: 1,
		UserId			: 1,
		DeviceTimestamp	: time.Now(),
		TotalAmount		: 100000,
		PaidAmount		: 110000,
		ChangeAmount	: 10000,
		PaymentMethod	: "cash",
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
	}, &entity.Transaction{
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

type transactionRouteMock struct {
	mock.Mock
}

func (s *transactionRouteMock) SaveTransaction(transaction entity.Transaction) (int, error) {
	return 0, nil
}

func (s *transactionRouteMock) UpdateTransaction(transaction entity.Transaction) error {
	return nil
}

func (s *transactionRouteMock) DeleteTransaction(transaction entity.Transaction) error {
	return nil
}

func (s *transactionRouteMock) GetAllTransactions() []entity.Transaction {
	return nil
}

func (s *transactionRouteMock) GetTransaction(ctx *gin.Context) []entity.Transaction {
	return nil
}

func (s *transactionRouteMock) AuthTransaction(transaction entity.Transaction) int {
	return 200
}

type TransactionRouteTestSuite struct {
	suite.Suite
	serviceTest usecases.TransactionService
}

func (suite *TransactionRouteTestSuite) SetupTest() {
	suite.serviceTest = new(transactionRouteMock)
}

func (suite *TransactionRouteTestSuite) TestSaveDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("transaction/create", TransactionCreate)

	jsonValue, _ := json.Marshal(dummyTransaction[0])
	req, _ := http.NewRequest(http.MethodPost, "/transaction/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *TransactionRouteTestSuite) TestUpdateDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("transaction/update", TransactionUpdate)

	jsonValue, _ := json.Marshal(dummyTransaction[0])
	req, _ := http.NewRequest(http.MethodPost, "/transaction/update", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *TransactionRouteTestSuite) TestDeleteDelivery() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("transaction/delete", TransactionDelete)

	jsonValue, _ := json.Marshal(dummyTransaction[0])
	req, _ := http.NewRequest(http.MethodPost, "/transaction/delete", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *TransactionRouteTestSuite) TestTransactionsIndexRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("transaction/index", TransactionsIndex)
	req, _ := http.NewRequest(http.MethodGet, "/transaction/index", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func (suite *TransactionRouteTestSuite) TestTransactionsDetailRoute() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("transaction/detail/1", TransactionsDetail)
	req, _ := http.NewRequest(http.MethodGet, "/transaction/detail/1", nil)

	r.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 200)
}

func TestTransactionRouteTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRouteTestSuite))
}
