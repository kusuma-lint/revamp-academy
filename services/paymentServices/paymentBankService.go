package services

import (
	"net/http"

	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentBankService struct {
	repositoriesManager *repositories.RepositoriesManager
}

func NewPaymentBankService(repoMgr *repositories.RepositoriesManager) *PaymentBankService {
	return &PaymentBankService{
		repositoriesManager: repoMgr,
	}
}

// 1a. ambil get list untuk payment bank
func (pbr PaymentBankService) GetListPaymentBank(ctx *gin.Context) ([]*dbContext.Bank, *models.ResponseError) {
	return pbr.repositoriesManager.PaymentBankRepository.GetListPaymentBank(ctx)
}

// 1a. ambil get by name untuk payment bank
func (pbr PaymentBankService) GetPaymentBankByName(ctx *gin.Context, name string) (*dbContext.Bank, *models.ResponseError) {
	return pbr.repositoriesManager.PaymentBankRepository.GetPaymentBankByName(ctx, name)
}

// 1b. buat create paymentbank
func (pbr PaymentBankService) CreateNewPaymentBank(ctx *gin.Context, paymentBankParams *dbContext.CreatePaymentBankParams) (*dbContext.Bank, *models.ResponseError) {
	responseErr := validatePaymentBank(paymentBankParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pbr.repositoriesManager.PaymentBankRepository.CreateNewPaymentBank(ctx, paymentBankParams)
}

// 1b. update payment data bank
func (pbr PaymentBankService) UpdatePaymentBank(ctx *gin.Context, paymentBankParams *dbContext.CreatePaymentBankParams, bankEntityID int64) *models.ResponseError {
	responseErr := validatePaymentBank(paymentBankParams)
	if responseErr != nil {
		return responseErr
	}
	return pbr.repositoriesManager.PaymentBankRepository.UpdatePaymentBank(ctx, paymentBankParams, bankEntityID)

}

// 1b. delet payment bank
func (pbr PaymentBankService) DeletePaymentBank(ctx *gin.Context, id int64) *models.ResponseError {
	return pbr.repositoriesManager.PaymentBankRepository.DeletePaymentBank(ctx, id)
}

func validatePaymentBank(paymentBankParams *dbContext.CreatePaymentBankParams) *models.ResponseError {
	// if paymentBankParams.BankEntityID == 0 {
	// 	return &models.ResponseError{
	// 		Message: "Invalid category id",
	// 		Status:  http.StatusBadRequest,
	// 	}
	// }

	if paymentBankParams.BankCode == "" {
		return &models.ResponseError{
			Message: "Invalid category code",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
