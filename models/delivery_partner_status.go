package models

type DeliveryPartnerStatus struct {
	DeliveryPartnerID     uint    `json:"deliveryID"`
	DeliveryPartnerStatus bool    `json:"status"`
	Lat                   float64 `json:"latitude"`
	Lng                   float64 `json:"longitude"`
}
