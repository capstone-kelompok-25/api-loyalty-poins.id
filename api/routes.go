package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/customer"
	"api-redeem-point/api/middleware"
	"api-redeem-point/api/store"

	//auth "api-redeem-point/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminController    *admin.Controller
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
	//admin
	g := c.Group("/admin")
	g.POST("/login", controller.AdminController.LoginAdmin)
	g.POST("", controller.AdminController.CreateAdmin)
	g.GET("", controller.AdminController.Dashboard)
	g.GET("/:id", controller.AdminController.FindAdminByID)
	g.GET("/transaction/pending", controller.AdminController.TransactionPending)
	g.POST("/approve/:idtransaction", controller.AdminController.ApproveTransaction)
	g.GET("/history", controller.AdminController.FindHistoryCustomers)
	g.GET("/customer", controller.AdminController.FindCustomers)
	g.PUT("/customer", controller.AdminController.UpdateCustomer)
	g.DELETE("/customer", controller.AdminController.DeleteCustomer)
	g.PUT("/customer/point", controller.AdminController.UpdateCustomerPoint)
	g.GET("/stock", controller.AdminController.StockProduct)
	g.PUT("/stock", controller.AdminController.UpdateStock)
	g.GET("/historystore", controller.AdminController.HistoryStore)
	g.DELETE("/store", controller.AdminController.DeleteStore)
	g.GET("/store", controller.AdminController.GetStore)
	g.PUT("/store", controller.AdminController.UpdateStore)
	s := c.Group("/store")
	s.POST("", controller.CustomerController.RegisterStore)
	s.POST("/login", controller.StoreController.LoginStore)
	s.POST("/poin", controller.StoreController.InputPoinStore, middleware.StoreSetupAuthenticationJWT())
}
