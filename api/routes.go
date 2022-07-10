package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/customer"
	"api-redeem-point/api/middleware"
	"api-redeem-point/api/store"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	CustomerController *customer.Controller
	AdminController    *admin.Controller
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
	g.GET("", controller.AdminController.Dashboard, middleware.AdminSetupAuthenticationJWT())
	g.GET("/:id", controller.AdminController.FindAdminByID, middleware.AdminSetupAuthenticationJWT())
	g.GET("/transaction/pending", controller.AdminController.TransactionPending, middleware.AdminSetupAuthenticationJWT())
	g.POST("/approve/:idtransaction", controller.AdminController.ApproveTransaction, middleware.AdminSetupAuthenticationJWT())
	g.GET("/history", controller.AdminController.FindHistoryCustomers, middleware.AdminSetupAuthenticationJWT())
	g.GET("/customer", controller.AdminController.FindCustomers, middleware.AdminSetupAuthenticationJWT())
	g.PUT("/customer", controller.AdminController.UpdateCustomer, middleware.AdminSetupAuthenticationJWT())
	g.DELETE("/customer", controller.AdminController.DeleteCustomer, middleware.AdminSetupAuthenticationJWT())
	g.PUT("/customer/point", controller.AdminController.UpdateCustomerPoint, middleware.AdminSetupAuthenticationJWT())
	g.GET("/stock", controller.AdminController.StockProduct, middleware.AdminSetupAuthenticationJWT())
	g.PUT("/stock", controller.AdminController.UpdateStock, middleware.AdminSetupAuthenticationJWT())
	g.GET("/historystore", controller.AdminController.HistoryStore, middleware.AdminSetupAuthenticationJWT())
	g.DELETE("/store", controller.AdminController.DeleteStore, middleware.AdminSetupAuthenticationJWT())
	g.GET("/store", controller.AdminController.GetStore, middleware.AdminSetupAuthenticationJWT())
	g.PUT("/store", controller.AdminController.UpdateStore, middleware.AdminSetupAuthenticationJWT())
	//store
	s := c.Group("/store")
	s.POST("", controller.CustomerController.RegisterStore)
	s.POST("/login", controller.StoreController.LoginStore)
	s.POST("/poin", controller.StoreController.InputPoinStore, middleware.StoreSetupAuthenticationJWT())
}
