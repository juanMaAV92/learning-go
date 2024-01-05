package notification

import (
	. "abstractFactory/notification/types"
	"errors"
	"fmt"
)

const (
	SMSNotificationType   = "SMS"
	EmailNotificationType = "Email"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case SMSNotificationType:
		return &SMSNotification{}, nil
	case EmailNotificationType:
		return &EmailNotification{}, nil
	default:
		return nil, errors.New("unsupported notification type")
	}
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func GetMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}
