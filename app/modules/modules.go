package modules

import (
	"api-redeem-point/api"
	customerApi "api-redeem-point/api/customer"
	customerBusiness "api-redeem-point/business/customer"
	"api-redeem-point/config"
	customerRepo "api-redeem-point/repository/customer"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {
	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	controller := api.Controller{
		CustomerController: customerPermitController,
	}
	return controller
}
