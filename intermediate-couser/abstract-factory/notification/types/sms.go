package types

import "fmt"

type SMSNotification struct {
}

func (*SMSNotification) SendNotification() {
	fmt.Println("Sending Notification by SMS")
}

func (*SMSNotification) GetSender() ISender {
	return &SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
