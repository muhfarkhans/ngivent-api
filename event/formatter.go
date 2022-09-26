package event

import (
	"time"
)

type EventFormatter struct {
	Id               int       `json:"id"`
	Title            string    `json:"title"`
	Slug             string    `json:"slug"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	DateStart        time.Time `json:"date_start"`
	DateEnd          time.Time `json:"date_end"`
	TypeEvent        string    `json:"type_event"`
	LocationLink     string    `json:"location_link"`
	LocationText     string    `json:"location_text"`
	Quota            int       `json:"quota"`
	Price            int       `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ImagePath        string    `json:"image"`
}

type EventImageFormatter struct {
	Id        int    `json:"id"`
	Image     string `json:"image"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatEvent(event Event) EventFormatter {
	formatter := EventFormatter{}
	formatter.Id = event.Id
	formatter.Title = event.Title
	formatter.Slug = event.Slug
	formatter.ShortDescription = event.ShortDescription
	formatter.Description = event.Description
	formatter.DateStart = event.DateStart
	formatter.DateEnd = event.DateEnd
	formatter.TypeEvent = event.TypeEvent
	formatter.LocationLink = event.LocationLink
	formatter.LocationText = event.LocationText
	formatter.Quota = event.Quota
	formatter.Price = event.Price
	formatter.UpdatedAt = event.UpdatedAt
	formatter.CreatedAt = event.CreatedAt
	formatter.ImagePath = ""

	if len(event.EventImages) > 0 {
		formatter.ImagePath = event.EventImages[0].Image
	}

	return formatter
}

func FormatEvents(events []Event) []EventFormatter {
	eventsFormatter := []EventFormatter{}

	for _, event := range events {
		eventFormatter := FormatEvent(event)
		eventsFormatter = append(eventsFormatter, eventFormatter)
	}

	return eventsFormatter
}

type EventDetailFormatter struct {
	Id               int                   `json:"id"`
	Title            string                `json:"title"`
	Slug             string                `json:"slug"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	DateStart        time.Time             `json:"date_start"`
	DateEnd          time.Time             `json:"date_end"`
	TypeEvent        string                `json:"type_event"`
	LocationLink     string                `json:"location_link"`
	LocationText     string                `json:"location_text"`
	Quota            int                   `json:"quota"`
	Price            int                   `json:"price"`
	CreatedAt        time.Time             `json:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at"`
	ImagePath        string                `json:"image"`
	Images           []EventImageFormatter `json:"images"`
}

func FormatEventDetail(event Event) EventDetailFormatter {
	eventDetailFormatter := EventDetailFormatter{}
	eventDetailFormatter.Id = event.Id
	eventDetailFormatter.Title = event.Title
	eventDetailFormatter.Slug = event.Slug
	eventDetailFormatter.ShortDescription = event.ShortDescription
	eventDetailFormatter.Description = event.Description
	eventDetailFormatter.DateStart = event.DateStart
	eventDetailFormatter.DateEnd = event.DateEnd
	eventDetailFormatter.TypeEvent = event.TypeEvent
	eventDetailFormatter.LocationLink = event.LocationLink
	eventDetailFormatter.LocationText = event.LocationText
	eventDetailFormatter.Quota = event.Quota
	eventDetailFormatter.Price = event.Price
	eventDetailFormatter.UpdatedAt = event.UpdatedAt
	eventDetailFormatter.CreatedAt = event.CreatedAt

	if len(event.EventImages) > 0 {
		eventDetailFormatter.ImagePath = event.EventImages[0].Image
	}

	images := []EventImageFormatter{}

	for _, image := range event.EventImages {
		eventImageFormatter := EventImageFormatter{}
		eventImageFormatter.Id = image.Id
		eventImageFormatter.Image = image.Image

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		eventImageFormatter.IsPrimary = isPrimary
		images = append(images, eventImageFormatter)
	}
	eventDetailFormatter.Images = images

	return eventDetailFormatter
}
