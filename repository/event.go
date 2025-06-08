package repository

import (
	"backend/contract"
	"backend/model"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func ImplEventRepository(db *gorm.DB) contract.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (t *EventRepository) GetAllEvent() ([]model.Event, error) {
	var event []model.Event
	if err := t.db.Find(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (t *EventRepository) GetEventByID(id uint64) (*model.Event, error) {
	var event model.Event
	if err := t.db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (t *EventRepository) CreateEvent(event *model.Event) (*model.Event, error) {
	if err := t.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (t *EventRepository) UpdateEvent(id uint64, event *model.Event) error {
	err := t.db.Model(event).Where("id = ?", id).Updates(event).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *EventRepository) DeleteEvent(id uint64) error {
	var event model.Event
	if err := t.db.Where("id = ?", id).Delete(&event).Error; err != nil {
		return err
	}
	return nil
}
