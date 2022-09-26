package event

import (
	"time"
)

type CreateEventInput struct {
	Title            string    `form:"title" binding:"required"`
	ShortDescription string    `form:"short_description" binding:"required"`
	Description      string    `form:"description" binding:"required"`
	DateStart        time.Time `form:"date_start" binding:"required" time_format:"2006-01-02"`
	DateEnd          time.Time `form:"date_end" binding:"required" time_format:"2006-01-02"`
	TypeEvent        string    `form:"type_event" binding:"required"`
	LocationLink     string    `form:"location_link" binding:"required"`
	LocationText     string    `form:"location_text" binding:"required"`
	Quota            int       `form:"quota" binding:"required"`
	Price            int       `form:"price" binding:"required"`
}

type GetEventDetailInput struct {
	Id int `uri:"id" binding:"required"`
}

type GetEventImageDetailInput struct {
	IdEvent int `uri:"id_event" binding:"required"`
	IdImage int `uri:"id_image" binding:"required"`
}

type UpdatePrimaryImageInput struct {
	IsPrimary bool `form:"is_primary"`
}

type CreateEventImageInput struct {
	EventId   int  `form:"event_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
}
