package repository

import (
	"backend/contract"
	"backend/model"
	"fmt"
	"strings"

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

func (t *EventRepository) GetAllEvent(search, status, sort string) ([]model.Event, error) {
	var events []model.Event
	query := t.db.Model(&model.Event{})

	if search != "" {
		query = query.Where("event_name ILIKE ? OR event_description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("event_status = ?", status)
	}

	if sort != "" {
		parts := strings.Split(sort, ":")
		if len(parts) == 2 {
			column := parts[0]
			order := strings.ToUpper(parts[1])
			if order == "ASC" || order == "DESC" {
				query = query.Order(fmt.Sprintf("%s %s", column, order))
			}
		}
	}

	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
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
