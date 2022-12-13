package rpcs

import (
	"context"

	"github.com/thirumathikart/thirumathikart-order-service/generated/notification"
)

func NotificationRPC(
	notificationContent *notification.SingleNotificationRequest,
	client notification.NotificationServiceClient) (*notification.SingleNotificationResponse, error) {
	return client.SingleNotificationRPC(context.Background(), notificationContent)
}
