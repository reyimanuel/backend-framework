package service

import (
	"backend/contract"
	"backend/dto"
	"backend/model"
	"backend/pkg/errs"
	"backend/pkg/helpers"
	"fmt"
	"net/http"
)

type EventService struct {
	EventRepository contract.EventRepository
}

func ImplEventService(repo *contract.Repository) contract.EventService {
	return &EventService{
		EventRepository: repo.EventRepository,
	}
}

func (t *EventService) GetAllEvent() (*dto.EventResponse, error) {
	event, err := t.EventRepository.GetAllEvent()
	if err != nil {
		return nil, errs.NotFound("no members found")
	}

	data := []dto.EventData{}
	for _, event := range event {
		data = append(data, dto.EventData{
			EventName:        event.EventName,
			EventDate:        event.EventDate,
			EventTime:        event.EventTime,
			EventLocation:    event.EventLocation,
			EventDescription: event.EventDescription,
			EventOrganizer:   event.EventOrganizer,
			EventStatus:      event.EventStatus,
			CreatedAt:        event.CreatedAt,
			UpdatedAt:        event.UpdatedAt,
		})
	}

	response := &dto.EventResponse{
		StatusCode: http.StatusOK,
		Message:    "Data retrieved successfully",
		Data:       data,
	}

	return response, nil
}
func (t *EventService) GetEventByID(id uint64) (*dto.EventResponse, error) {
	event, err := t.EventRepository.GetEventByID(id)
	if err != nil {
		return nil, errs.NotFound("member not found")
	}

	data := dto.EventData{
		EventName:        event.EventName,
		EventDate:        event.EventDate,
		EventTime:        event.EventTime,
		EventLocation:    event.EventLocation,
		EventDescription: event.EventDescription,
		EventOrganizer:   event.EventOrganizer,
		EventStatus:      event.EventStatus,
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}

	response := &dto.EventResponse{
		StatusCode: http.StatusOK,
		Message:    "Data retrieved successfully",
		Data:       []dto.EventData{data},
	}

	return response, nil
}

func (t *EventService) CreateEvent(payload *dto.EventRequest) (*dto.EventResponse, error) {
	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, validPayload
	}

	event := &model.Event{
		EventName:        payload.EventName,
		EventDate:        payload.EventDate,
		EventTime:        payload.EventTime,
		EventLocation:    payload.EventLocation,
		EventDescription: payload.EventDescription,
		EventOrganizer:   payload.EventOrganizer,
		EventStatus:      payload.EventStatus,
	}

	newEvent, err := t.EventRepository.CreateEvent(event)
	if err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	response := &dto.EventResponse{
		StatusCode: http.StatusCreated,
		Message:    "member created successfully",
		Data: []dto.EventData{
			{
				EventName:        newEvent.EventName,
				EventDate:        newEvent.EventDate,
				EventTime:        newEvent.EventTime,
				EventLocation:    newEvent.EventLocation,
				EventDescription: newEvent.EventDescription,
				EventOrganizer:   newEvent.EventOrganizer,
				EventStatus:      newEvent.EventStatus,
				CreatedAt:        newEvent.CreatedAt,
				UpdatedAt:        newEvent.UpdatedAt,
			},
		},
	}
	return response, nil

}
