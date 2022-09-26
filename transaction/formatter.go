package transaction

import (
	"ngevent-api/event"
	"ngevent-api/helper"
	"ngevent-api/user"
	"time"
)

type TransactionsFormatter struct {
	Pagination helper.Pagination      `json:"pagination"`
	Rows       []TransactionFormatter `json:"rows"`
}

type TransactionFormatter struct {
	Id         int       `json:"id"`
	EventId    int       `json:"event_id"`
	OrderId    string    `json:"order_id"`
	UserId     int       `json:"user_id"`
	UserName   string    `json:"user_name"`
	UserEmail  string    `json:"user_email"`
	Amount     int       `json:"amount"`
	Qty        int       `json:"qty"`
	Status     string    `json:"status"`
	Code       int       `json:"code"`
	PaymentUrl string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	EventTitle string    `json:"event_title"`
	EventPrice int       `json:"event_price"`
}

type DetailTransactionFormatter struct {
	Id         int         `json:"id"`
	EventId    int         `json:"event_id"`
	OrderId    string      `json:"order_id"`
	UserId     int         `json:"user_id"`
	Amount     int         `json:"amount"`
	Qty        int         `json:"qty"`
	Status     string      `json:"status"`
	Code       int         `json:"code"`
	PaymentUrl string      `json:"payment_url"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Event      event.Event `json:"event"`
	User       user.User   `json:"user"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.EventId = transaction.EventId
	formatter.OrderId = transaction.OrderId
	formatter.UserId = transaction.UserId
	formatter.UserName = transaction.User.Name
	formatter.UserEmail = transaction.User.Email
	formatter.Amount = transaction.Amount
	formatter.Qty = transaction.Qty
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentUrl = transaction.PaymentUrl
	formatter.CreatedAt = transaction.CreatedAt
	formatter.UpdatedAt = transaction.UpdatedAt
	formatter.EventTitle = transaction.Event.Title
	formatter.EventPrice = transaction.Event.Price

	return formatter
}

func FormatDetailTransaction(transaction Transaction) DetailTransactionFormatter {
	formatter := DetailTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.EventId = transaction.EventId
	formatter.OrderId = transaction.OrderId
	formatter.UserId = transaction.UserId
	formatter.Amount = transaction.Amount
	formatter.Qty = transaction.Qty
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentUrl = transaction.PaymentUrl
	formatter.CreatedAt = transaction.CreatedAt
	formatter.UpdatedAt = transaction.UpdatedAt
	formatter.Event = transaction.Event
	formatter.User = transaction.User

	return formatter
}

func FormatTransactions(trasnsactions []Transaction, paginate *helper.Pagination) TransactionsFormatter {

	transactionsFormatter := []TransactionFormatter{}
	for _, transaction := range trasnsactions {
		transactionFormatter := FormatTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	paginateFormatter := TransactionsFormatter{}
	paginateFormatter.Pagination = *paginate
	paginateFormatter.Rows = transactionsFormatter

	return paginateFormatter
}
