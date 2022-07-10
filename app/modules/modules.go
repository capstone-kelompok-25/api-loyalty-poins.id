package modules

import (
	"api-redeem-point/api"
	customerApi "api-redeem-point/api/customer"
	storeApi "api-redeem-point/api/store"
	customerBusiness "api-redeem-point/business/customer"
	storeBusiness "api-redeem-point/business/store"
	"api-redeem-point/config"
	customerRepo "api-redeem-point/repository/customer"
	storeRepo "api-redeem-point/repository/store"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {
	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	storePermitRepository := storeRepo.RepositoryFactory(dbCon)
	storePermitService := storeBusiness.NewService(storePermitRepository)
	storePermitController := storeApi.NewController(storePermitService)
	controller := api.Controller{
		CustomerController: customerPermitController,
		StoreController:    storePermitController,
	}
	return controller
}
