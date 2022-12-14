package repositories

import (
	"github.com/thirumathikart/thirumathikart-order-service/generated/products"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/schemas"
	"github.com/thirumathikart/thirumathikart-order-service/utils"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

type OrderRepository interface {
	CreateOrder(
		user *user.User,
		productInfo *products.GetProductsResponse,
		requestedItems []models.CreateOrderItem,
	) error

	UpdateOrderStatus(
		orderID uint,
		orderStatus schemas.OrderStatus,
	) error

	FindCustomer(
		orderID uint,
	) (uint, error)
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (or *orderRepository) CreateOrder(
	user *user.User,
	productInfo *products.GetProductsResponse,
	requestedItems []models.CreateOrderItem,
) error {

	order := schemas.Order{
		CustomerID:        uint(user.UserId),
		CustomerAddressID: uint(*user.AddressId),
		OrderStatus:       schemas.BuyerOrdered,
	}

	if err := or.db.Create(&order).Error; err != nil {
		return err
	}
	quantityFromID := utils.QuantityFromID(requestedItems)

	orderItems := []schemas.OrderItem{}

	for _, product := range productInfo.GetProducts() {
		orderItem := schemas.OrderItem{
			OrderID:  order.ID,
			Name:     product.ProductTitle,
			Price:    uint(product.ProductPrice),
			Quantity: quantityFromID[uint(product.ProductId)],
		}
		orderItems = append(orderItems, orderItem)
	}

	if err := or.db.Create(&orderItems).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) UpdateOrderStatus(
	orderID uint,
	orderStatus schemas.OrderStatus,
) error {

	if err :=
		or.db.Model(&schemas.Order{}).Where("id = ?", orderID).Update("order_status", orderStatus).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) FindCustomer(
	orderID uint,
) (uint, error) {
	var order schemas.Order
	query := or.db.Table("users").Select("customer_id").Where("name = ?", orderID).Scan(&order)
	if query.Error != nil {
		return 0, query.Error
	}
	return order.CustomerID, nil
}
