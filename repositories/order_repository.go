package repositories

import (
	"log"

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
		AddressId uint,
	) error

	UpdateOrderStatus(
		orderID uint,
		orderStatus schemas.OrderStatus,
	) error

	FindCustomer(
		orderID uint,
	) (uint, error)

	GetDeliveryPartners() ([]schemas.DeliveryPartner, error)

	GetOrder(
		orderID uint,
	) (schemas.Order, error)

	AssignDeliveryPartner(
		orderID uint,
		DeliveryPartnerID uint,
	) error

	UpdateDelvieryPartnerStatus(
		deliveryPartnerID uint,
		deliveryPartnerStatus bool,
		Lat float64,
		Lng float64,
	) error
	FetchOrderByDeliveryPartner(
		SellerID uint,
	) ([]schemas.Order, error)

	FetchOrderByCustomer(
		SellerID uint,
	) ([]schemas.Order, error)

	FetchOrderBySeller(
		SellerID uint,
	) ([]schemas.Order, error)
	
	FetchOrderItemsByOrder(
		OrderID uint,
	) ([]schemas.OrderItem, error)

}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (or *orderRepository) CreateOrder(
	user *user.User,
	productInfo *products.GetProductsResponse,
	requestedItems []models.CreateOrderItem,
	AddressId uint,
) error {

	log.Println(".............-.",user)
	order := schemas.Order{
		CustomerID:        uint(user.UserId),
		CustomerAddressID: AddressId,
		OrderStatus:       schemas.BuyerOrdered,
	}
	log.Println(".............-.",order)

	if err := or.db.Create(&order).Error; err != nil {
		return err
	}
	quantityFromID := utils.QuantityFromID(requestedItems)

	orderItems := []schemas.OrderItem{}
	log.Println(".............-.",productInfo.Products)

	for _, product := range productInfo.Products {
		log.Println("..............+",product)
		orderItem := schemas.OrderItem{
			OrderID:  1,
			Name:     product.ProductTitle,
			Price:    uint(product.ProductPrice),
			Quantity: quantityFromID[uint(product.ProductId)],
		}
		orderItems = append(orderItems, orderItem)
	}
	log.Println("...............",orderItems)

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

func (or *orderRepository) GetDeliveryPartners() ([]schemas.DeliveryPartner, error) {
	var deliveryPartners []schemas.DeliveryPartner
	query := or.db.Find(&deliveryPartners)
	return deliveryPartners, query.Error
}

func (or *orderRepository) GetOrder(
	orderID uint,
) (schemas.Order, error) {
	var order schemas.Order
	res := or.db.Where("id = ?", orderID).Find(&order)
	return order, res.Error
}

func (or *orderRepository) AssignDeliveryPartner(
	orderID uint,
	DeliveryPartnerID uint,
) error {
	order := schemas.Order{
		DeliveryPartnerID: DeliveryPartnerID,
		OrderStatus:       schemas.DeliveryPartnerAssigned,
	}
	if err :=
		or.db.Model(&schemas.Order{}).Where("id = ?", orderID).Updates(order).Error; err != nil {
		return err
	}
	return nil
}
func (or *orderRepository) UpdateDelvieryPartnerStatus(
	deliveryPartnerID uint,
	deliveryPartnerStatus bool,
	Lat float64,
	Lng float64) error {
	if err :=
		or.db.Model(&schemas.DeliveryPartner{}).Where("delivery_partner_id = ?", deliveryPartnerID).Updates(map[string]interface{}{"status": deliveryPartnerStatus, "Lat": Lat, "Lng": Lng}).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepository) FetchOrderItemsByOrder(
	OrderID uint,
) ([]schemas.OrderItem, error) {
	var orders []schemas.OrderItem
	err:=or.db.Find(&orders).Where("order_id = ?",OrderID).Preload("Order").Error	
	return orders,err
}


func (or *orderRepository) FetchOrderBySeller(
	SellerID uint,
) ([]schemas.Order, error) {
	var orders []schemas.Order
	err:=or.db.Find(&orders).Where("seller_id = ?",SellerID).Error	
	return orders,err
}

func (or *orderRepository) FetchOrderByDeliveryPartner(
	DeliveryPartnerID uint,
) ([]schemas.Order, error) {
	var orders []schemas.Order
	err:=or.db.Find(&orders).Where("delivery_partner_id = ?",DeliveryPartnerID).Error	
	return orders,err
}

func (or *orderRepository) FetchOrderByCustomer(
	CustomerID uint,
) ([]schemas.Order, error) {
	var orders []schemas.Order
	log.Println(CustomerID)
	err:=or.db.Where("customer_id = ?",CustomerID).Find(&orders).Error	
	log.Println(err)
	return orders,err
}

