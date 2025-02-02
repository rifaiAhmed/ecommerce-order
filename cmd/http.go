package cmd

import (
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/api"
	"ecommerce-order/internal/interfaces"
	"ecommerce-order/internal/repository"
	"ecommerce-order/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()
	healthcheckAPI := &api.HealthcheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthcheckAPI.Healthcheck)

	orderV1 := e.Group("orders/v1")
	orderV1.POST("", d.OrderAPI.CreateOrder, d.MiddlewareValidateAuth)
	orderV1.PUT("/:id/status", d.OrderAPI.UpdateOrderStatus, d.MiddlewareValidateAuth)
	orderV1.PUT("/in/:id/status", d.OrderAPI.UpdateOrderStatus)
	orderV1.GET("/:id", d.OrderAPI.GetOrderDetail, d.MiddlewareValidateAuth)
	orderV1.GET("", d.OrderAPI.GetOrderList, d.MiddlewareValidateAuth)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthcheckAPI

	OrderAPI interfaces.IOrderAPI
}

func dependencyInject() Dependency {
	external := &external.External{}

	orderRepo := &repository.OrderRepo{
		DB: helpers.DB,
	}

	orderSvc := &services.OrderService{
		OrderRepo: orderRepo,
		External:  external,
	}

	orderAPI := &api.OrderAPI{
		OrderService: orderSvc,
	}

	return Dependency{
		External:       external,
		HealthcheckAPI: &api.HealthcheckAPI{},

		OrderAPI: orderAPI,
	}
}
