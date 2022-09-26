package event

import "time"

type Event struct {
	Id               int          `json:"id"`
	Title            string       `json:"title"`
	Slug             string       `json:"slug"`
	ShortDescription string       `json:"short_description"`
	Description      string       `json:"description"`
	DateStart        time.Time    `json:"date_start"`
	DateEnd          time.Time    `json:"date_end"`
	TypeEvent        string       `json:"type_event"`
	LocationLink     string       `json:"location_link"`
	LocationText     string       `json:"location_text"`
	Quota            int          `json:"quota"`
	Price            int          `json:"price"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	EventImages      []EventImage `gorm:"ForeignKey:EventId"`
}

type EventImage struct {
	Id        int       `json:"id"`
	EventId   int       `json:"event_id"`
	IsPrimary int       `json:"is_primary"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
