package api

import (
	"api-redeem-point/api/customer"
	"api-redeem-point/api/middleware"
	"api-redeem-point/api/store"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	CustomerController *customer.Controller
	StoreController    *store.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	c := e.Group("/v1")
	c.POST("/customer/register", controller.CustomerController.Register)
	c.POST("/customer", controller.CustomerController.Login)
	c.PUT("/customer", controller.CustomerController.UpdateCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/customer/:id", controller.CustomerController.FindCustomersByID, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/history", controller.CustomerController.HistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/dethistory/:idtransaction", controller.CustomerController.DetailHistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/pulsa", controller.CustomerController.OrderPulsa, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/paketdata", controller.CustomerController.OrderPaketData, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/cashout", controller.CustomerController.OrderCashout, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/emoney", controller.CustomerController.OrderEmoney, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/callback", controller.CustomerController.CallbackXendit)
	//store
	s := c.Group("/store")
	s.POST("", controller.CustomerController.RegisterStore)
	s.POST("/login", controller.StoreController.LoginStore)
	s.POST("/poin", controller.StoreController.InputPoinStore, middleware.StoreSetupAuthenticationJWT())
}
