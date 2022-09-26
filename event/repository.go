package event

import "gorm.io/gorm"

type Repository interface {
	Save(event Event) (Event, error)
	First(id int) (Event, error)
	Update(event Event) (Event, error)
	Delete(event Event) (Event, error)
	FindById(id int) (Event, error)
	FindAll() ([]Event, error)
	SaveImage(eventImage EventImage) (EventImage, error)
	MarkAllImagesAsNonPrimary(eventId int) (bool, error)
	MarkImagePrimaryAsPrimary(imageId int) (bool, error)
	MarkImagePrimaryAsNonPrimary(imageId int) (bool, error)
	DeleteImage(eventImage EventImage) (EventImage, error)
	FindImageById(id int) (EventImage, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(event Event) (Event, error) {
	err := r.db.Create(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) First(id int) (Event, error) {
	var event Event
	err := r.db.Where("id = ?", id).Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) Update(event Event) (Event, error) {
	err := r.db.Save(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) Delete(event Event) (Event, error) {
	err := r.db.Delete(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) FindById(id int) (Event, error) {
	var event Event
	err := r.db.Preload("EventImages").Where("id = ?", id).Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event
	err := r.db.Preload("EventImages", "event_images.is_primary = 1").Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *repository) SaveImage(eventImage EventImage) (EventImage, error) {
	err := r.db.Create(&eventImage).Error
	if err != nil {
		return eventImage, err
	}

	return eventImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(eventId int) (bool, error) {
	err := r.db.Model(&EventImage{}).Where("event_id = ?", eventId).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) MarkImagePrimaryAsPrimary(imageId int) (bool, error) {
	err := r.db.Model(&EventImage{}).Where("id = ?", imageId).Update("is_primary", true).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) MarkImagePrimaryAsNonPrimary(imageId int) (bool, error) {
	err := r.db.Model(&EventImage{}).Where("id = ?", imageId).Update("is_primary", true).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) DeleteImage(eventImage EventImage) (EventImage, error) {
	err := r.db.Delete(&eventImage).Error
	if err != nil {
		return eventImage, err
	}

	return eventImage, nil
}

func (r *repository) FindImageById(id int) (EventImage, error) {
	var eventImage EventImage
	err := r.db.Where("id = ?", id).Find(&eventImage).Error
	if err != nil {
		return eventImage, err
	}

	return eventImage, nil
}
