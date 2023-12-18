package main

import "abstractFactory/notification"

func main() {
	notificationFactory, err := notification.GetNotificationFactory(notification.EmailNotificationType)
	if err != nil {
		panic(err)
	}
	notification.SendNotification(notificationFactory)
	notification.GetMethod(notificationFactory)

	notificationFactory2, err := notification.GetNotificationFactory(notification.SMSNotificationType)
	if err != nil {
		panic(err)
	}
	notification.SendNotification(notificationFactory2)
	notification.GetMethod(notificationFactory2)
}
