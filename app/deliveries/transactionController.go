package deliveries

import (
	"net/http"

	entity "github.com/coroo/go-pawoon/app/entity"
	repositories "github.com/coroo/go-pawoon/app/repositories"
	usecases "github.com/coroo/go-pawoon/app/usecases"
	utils "github.com/coroo/go-pawoon/app/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	transactionRepository repositories.TransactionRepository = repositories.NewTransactionRepository()
	transactionService    usecases.TransactionService        = usecases.NewTransaction(transactionRepository)
	// transactionController deliveries.TransactionController   = deliveries.NewTransaction(transactionService)
)

// GetTransactionsIndex godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show all existing Transactions
// @Description Get all existing Transactions
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Transaction
// @Failure 401 {object} dto.Response
// @Router /transaction/index [get]
func TransactionsIndex(c *gin.Context) {
	transactions := transactionService.GetAllTransactions()
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// GetTransactionsDetail godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Show an existing Transactions
// @Description Get detail the existing Transactions
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param  id path int true "Master Question ID"
// @Success 200 {array} entity.Transaction
// @Failure 401 {object} dto.Response
// @Router /transaction/detail/{id} [get]
func TransactionsDetail(c *gin.Context) {
	transaction := transactionService.GetTransaction(c)
	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// CreateTransactions godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Create new Transactions
// @Description Create a new Transactions
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param transaction body entity.TransactionCreate true "Create transaction"
// @Success 200 {object} entity.Transaction
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /transaction/create [post]
func TransactionCreate(c *gin.Context) {
	claims, _ := utils.ExtractClaims(c.Request.Header["Authorization"][0])
	var transactionEntity entity.Transaction
	c.ShouldBindJSON(&transactionEntity)
	if claims["user_id"] != nil {
		transactionPK, err := transactionService.SaveTransaction(transactionEntity, int(claims["user_id"].(float64)))
		if(err!=nil){
		c.JSON(http.StatusConflict, err)
		} else {
			transactionEntity.ID = transactionPK
			c.JSON(http.StatusOK, transactionEntity)
		}
	}
}

// UpdateTransactions godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Update Transactions
// @Description Update a Transactions
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param transaction body entity.Transaction true "Update transaction"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /transaction/update [put]
func TransactionUpdate(c *gin.Context) {
	var transactionEntity entity.Transaction
	c.ShouldBindJSON(&transactionEntity)
	transaction := transactionService.UpdateTransaction(transactionEntity)
	c.JSON(http.StatusOK, transaction)
}

// DeleteTransactions godoc
// @Param Authorization header string true "Bearer"
// @Security basicAuth
// @Summary Delete Transactions
// @Description Delete a Transactions
// @Tags Transactions
// @Accept  json
// @Produce  json
// @Param transaction body entity.TransactionDelete true "Delete transaction"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /transaction/delete [delete]
func TransactionDelete(c *gin.Context) {
	var transactionEntity entity.Transaction
	c.ShouldBindJSON(&transactionEntity)
	transaction := transactionService.DeleteTransaction(transactionEntity)
	c.JSON(http.StatusOK, transaction)
}