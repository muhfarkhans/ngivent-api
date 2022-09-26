package transaction

import "ngevent-api/user"

type CreateTransactionInput struct {
	EventId    int    `json:"event_id" binding:"required"`
	OrderId    string `json:"order_id"`
	Amount     int    `json:"amount"`
	Qty        int    `json:"qty" binding:"required"`
	Code       int    `json:"code"`
	PaymentUrl string `json:"payment_url"`
	User       user.User
}

type ParamsGetTransaction struct {
	EventId int    `form:"event_id"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	Q       string `form:"q"`
	User    user.User
}

type GetTransactionDetailInput struct {
	Id int `uri:"id" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
