package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	customerApi "api-redeem-point/api/customer"
	adminBusiness "api-redeem-point/business/admin"
	customerBusiness "api-redeem-point/business/customer"
	"api-redeem-point/config"
	adminRepo "api-redeem-point/repository/admin"
	customerRepo "api-redeem-point/repository/customer"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {
	customerPermitRepository := customerRepo.RepositoryFactory(dbCon)
	customerPermitService := customerBusiness.NewService(customerPermitRepository)
	customerPermitController := customerApi.NewController(customerPermitService)

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	controller := api.Controller{
		CustomerController: customerPermitController,
		AdminController:    adminPermitController,
	}
	return controller
}
