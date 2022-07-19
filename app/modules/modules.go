package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	customerApi "api-redeem-point/api/customer"
	storeApi "api-redeem-point/api/store"
	adminBusiness "api-redeem-point/business/admin"
	customerBusiness "api-redeem-point/business/customer"
	storeBusiness "api-redeem-point/business/store"
	"api-redeem-point/config"
	adminRepo "api-redeem-point/repository/admin"
	customerRepo "api-redeem-point/repository/customer"
	storeRepo "api-redeem-point/repository/store"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, _ *config.AppConfig) api.Controller {
	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	storePermitRepository := storeRepo.RepositoryFactory(dbCon)
	storePermitService := storeBusiness.NewService(storePermitRepository)
	storePermitController := storeApi.NewController(storePermitService)

	controller := api.Controller{
		CustomerController: customerPermitController,
		AdminController:    adminPermitController,
		StoreController:    storePermitController,
	}
	return controller
}
