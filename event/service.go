package event

import (
	"errors"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateEvent(input CreateEventInput, pathEventImage string) (Event, error)
	GetEvents() ([]Event, error)
	FindEvent(input GetEventDetailInput) (Event, error)
	CreateEventImage(input GetEventDetailInput, pathEventImage string) (EventImage, error)
	UpdateEvent(input CreateEventInput, inputUri GetEventDetailInput) (Event, error)
	UpdateImagePrimary(input UpdatePrimaryImageInput, inputUri GetEventImageDetailInput) (Event, error)
	DeleteEvent(input GetEventDetailInput) (Event, error)
	DeleteEventImage(input GetEventImageDetailInput) (EventImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateEvent(input CreateEventInput, pathEventImage string) (Event, error) {
	event := Event{}
	event.Title = input.Title
	event.ShortDescription = input.ShortDescription
	event.Description = input.Description
	event.DateStart = input.DateStart
	event.DateEnd = input.DateEnd
	event.TypeEvent = input.TypeEvent
	event.LocationLink = input.LocationLink
	event.LocationText = input.LocationText
	event.Quota = input.Quota
	event.Price = input.Price

	titleSlug := slug.Make(input.Title)
	event.Slug = titleSlug

	eventImage := EventImage{}
	eventImage.IsPrimary = 1
	eventImage.Image = pathEventImage
	event.EventImages = []EventImage{eventImage}

	newEvent, err := s.repository.Save(event)
	if err != nil {
		return newEvent, err
	}

	return newEvent, nil
}

func (s *service) GetEvents() ([]Event, error) {
	events, err := s.repository.FindAll()
	if err != nil {
		return events, err
	}

	return events, nil
}

func (s *service) FindEvent(input GetEventDetailInput) (Event, error) {
	event, err := s.repository.FindById(input.Id)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *service) CreateEventImage(input GetEventDetailInput, pathEventImage string) (EventImage, error) {
	event, err := s.repository.FindById(input.Id)
	if err != nil {
		return EventImage{}, err
	}

	if event.Id == 0 {
		return EventImage{}, errors.New("event not found")
	}

	eventImage := EventImage{}
	eventImage.IsPrimary = 0
	eventImage.EventId = input.Id
	eventImage.Image = pathEventImage
	newEventImage, err := s.repository.SaveImage(eventImage)
	if err != nil {
		return newEventImage, err
	}

	return newEventImage, nil
}

func (s *service) UpdateEvent(input CreateEventInput, inputUri GetEventDetailInput) (Event, error) {
	event, err := s.repository.FindById(inputUri.Id)
	if err != nil {
		return event, err
	}

	if event.Id == 0 {
		return event, errors.New("event not found")
	}

	event.Title = input.Title
	event.ShortDescription = input.ShortDescription
	event.Description = input.Description
	event.DateStart = input.DateStart
	event.DateEnd = input.DateEnd
	event.TypeEvent = input.TypeEvent
	event.LocationLink = input.LocationLink
	event.LocationText = input.LocationText
	event.Quota = input.Quota
	event.Price = input.Price

	updatedEvent, err := s.repository.Update(event)
	if err != nil {
		return updatedEvent, err
	}

	return updatedEvent, nil
}

func (s *service) UpdateImagePrimary(input UpdatePrimaryImageInput, inputUri GetEventImageDetailInput) (Event, error) {
	event, err := s.repository.FindById(inputUri.IdEvent)
	if err != nil {
		return event, err
	}

	if event.Id == 0 {
		return event, errors.New("event not found")
	}

	imageFound := 0
	for _, image := range event.EventImages {
		if image.Id == inputUri.IdImage {
			imageFound = 1
		}
	}

	if imageFound == 0 {
		return event, errors.New("image event not found")
	}

	if input.IsPrimary {
		_, err := s.repository.MarkAllImagesAsNonPrimary(inputUri.IdEvent)
		if err != nil {
			return event, err
		}

		_, err = s.repository.MarkImagePrimaryAsPrimary(inputUri.IdImage)
		if err != nil {
			return event, err
		}
	}

	event, err = s.repository.FindById(inputUri.IdEvent)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *service) DeleteEvent(input GetEventDetailInput) (Event, error) {
	event, err := s.repository.FindById(input.Id)
	if err != nil {
		return event, err
	}

	if event.Id == 0 {
		return event, errors.New("event not found")
	}

	_, err = s.repository.Delete(event)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *service) DeleteEventImage(input GetEventImageDetailInput) (EventImage, error) {
	eventImage, err := s.repository.FindImageById(input.IdImage)
	if err != nil {
		return eventImage, err
	}

	if eventImage.Id == 0 {
		return eventImage, errors.New("image not found")
	}

	if eventImage.IsPrimary == 1 {
		return eventImage, errors.New("image is primary")
	}

	_, err = s.repository.DeleteImage(eventImage)
	if err != nil {
		return eventImage, err
	}

	return eventImage, nil
}
