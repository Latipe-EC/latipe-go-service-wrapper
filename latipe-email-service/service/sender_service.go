package service

import "email-service/data/dto"

type SenderEmailService interface {
	SendOrderEmail(message *dto.OrderMessage) error
	SendRegisterEmail(message *dto.UserRegisterMessage) error
	SendForgotPassword(message *dto.ForgotPasswordMessage) error
	SendDeliveryAccount(message *dto.DeliveryAccountMessage) error
	SendTakeoutPayment(message *dto.PaymentMessage) error
}
