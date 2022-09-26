package transaction

import (
	"ngevent-api/event"
	"ngevent-api/user"
	"time"
)

type Transaction struct {
	Id         int         `json:"id"`
	EventId    int         `json:"event_id"`
	UserId     int         `json:"user_id"`
	OrderId    string      `json:"order_id"`
	Amount     int         `json:"amount"`
	Qty        int         `json:"qty"`
	Status     string      `json:"status"`
	Code       int         `json:"code"`
	PaymentUrl string      `json:"payment_url"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Event      event.Event `gorm:"ForeignKey:EventId"`
	User       user.User   `gorm:"ForeignKey:UserId"`
}
