package types

import "fmt"

type EmailNotification struct {
}

func (*EmailNotification) SendNotification() {
	fmt.Println("Sending Notification by Email")
}

func (*EmailNotification) GetSender() ISender {
	return &EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "ProtonMail"
}
