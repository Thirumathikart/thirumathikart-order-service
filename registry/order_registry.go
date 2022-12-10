package registry

import (
	"github.com/thirumathikart/thirumathikart-order-service/controllers"
	"github.com/thirumathikart/thirumathikart-order-service/services"
)

func (r *registry) NewOrderController() controllers.OrderController {
	return controllers.NewOrderController(r.NewOrderService())
}

func (r *registry) NewOrderService() services.OrderService {
	return services.NewOrderService()
}
