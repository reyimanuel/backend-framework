package service

import (
	"backend/config"
	"backend/contract"
	"backend/dto"
	"backend/model"
	"backend/pkg/errs"
	"backend/pkg/helpers"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type EventService struct {
	EventRepository contract.EventRepository
	baseURL         string
}

func ImplEventService(repo *contract.Repository) contract.EventService {
	return &EventService{
		EventRepository: repo.EventRepository,
		baseURL:         config.Get().BaseURL,
	}
}

func (t *EventService) GetAllEvent(search, status, sort string) (*dto.EventResponse, error) {
	event, err := t.EventRepository.GetAllEvent(search, status, sort)
	if err != nil {
		return nil, errs.NotFound("no members found")
	}

	data := []dto.EventData{}
	for _, event := range event {
		data = append(data, dto.EventData{
			ID:               event.ID,
			EventName:        event.EventName,
			EventDate:        event.EventDate.Format("2006-01-02"),
			EventTime:        event.EventTime,
			EventLocation:    event.EventLocation,
			EventDescription: event.EventDescription,
			EventOrganizer:   event.EventOrganizer,
			EventStatus:      event.EventStatus,
			EventImageURL:    fmt.Sprintf("%s%s", t.baseURL, event.EventImageURL),
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
		ID:               event.ID,
		EventName:        event.EventName,
		EventDate:        event.EventDate.Format("2006-01-02"),
		EventTime:        event.EventTime,
		EventLocation:    event.EventLocation,
		EventDescription: event.EventDescription,
		EventOrganizer:   event.EventOrganizer,
		EventStatus:      event.EventStatus,
		EventImageURL:    fmt.Sprintf("%s%s", t.baseURL, event.EventImageURL),
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

func (t *EventService) CreateEvent(ctx *gin.Context, payload *dto.EventRequest, file *multipart.FileHeader) (*dto.EventResponse, error) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, err
	}

	imageName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	imagePath := fmt.Sprintf("static/%s", imageName)
	if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}
	imageURL := fmt.Sprintf("/static/%s", imageName)

	eventDate, err := time.Parse("2006-01-02", payload.EventDate)
	if err != nil {
		return nil, fmt.Errorf("invalid event_date format, expected YYYY-MM-DD")
	}

	eventTime, err := time.Parse("15:04", payload.EventTime)
	if err != nil {
		return nil, fmt.Errorf("invalid event_time format, expected HH:MM")
	}

	event := &model.Event{
		EventName:        payload.EventName,
		EventDate:        eventDate,
		EventTime:        eventTime.Format("15:04:00"),
		EventLocation:    payload.EventLocation,
		EventDescription: payload.EventDescription,
		EventOrganizer:   payload.EventOrganizer,
		EventStatus:      payload.EventStatus,
		EventImageURL:    imageURL,
	}

	newEvent, err := t.EventRepository.CreateEvent(event)
	if err != nil {
		_ = os.Remove(imagePath)
		return nil, fmt.Errorf("failed to create event: %w", err)
	}

	response := &dto.EventResponse{
		StatusCode: http.StatusCreated,
		Message:    "Event created successfully",
		Data: []dto.EventData{
			{
				ID:               newEvent.ID,
				EventName:        newEvent.EventName,
				EventDate:        newEvent.EventDate.Format("2006-01-02"),
				EventTime:        newEvent.EventTime,
				EventLocation:    newEvent.EventLocation,
				EventDescription: newEvent.EventDescription,
				EventOrganizer:   newEvent.EventOrganizer,
				EventStatus:      newEvent.EventStatus,
				EventImageURL:    fmt.Sprintf("%s%s", t.baseURL, newEvent.EventImageURL),
				CreatedAt:        newEvent.CreatedAt,
				UpdatedAt:        newEvent.UpdatedAt,
			},
		},
	}

	return response, nil
}

func (t *EventService) UpdateEvent(ctx *gin.Context, id uint64, payload *dto.EventRequest, file *multipart.FileHeader) (*dto.EventResponse, error) {
	oldEvent, err := t.EventRepository.GetEventByID(id)
	if err != nil {
		return nil, errs.NotFound("Event not found")
	}

	var imageURL string = oldEvent.EventImageURL

	if file != nil {
		imageName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		imagePath := fmt.Sprintf("static/%s", imageName)

		if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
			return nil, errs.InternalServerError("Failed to save image")
		}

		imageURL = fmt.Sprintf("/static/%s", imageName)

		if oldEvent.EventImageURL != "" {
			oldImagePath := fmt.Sprintf("static/%s", filepath.Base(oldEvent.EventImageURL))
			_ = os.Remove(oldImagePath)
		}
	}

	var eventDate time.Time
	if strings.TrimSpace(payload.EventDate) != "" {
		eventDate, err = time.Parse("2006-01-02", strings.TrimSpace(payload.EventDate))
		if err != nil {
			return nil, fmt.Errorf("invalid event_date format, expected YYYY-MM-DD")
		}
	} else {
		// kosong, pakai oldEvent.EventDate
		eventDate = oldEvent.EventDate
	}

	var eventTime time.Time
	if strings.TrimSpace(payload.EventTime) != "" {
		eventTime, err = time.Parse("15:04", strings.TrimSpace(payload.EventTime))
		if err != nil {
			return nil, fmt.Errorf("invalid event_time format, expected HH:MM")
		}
	} else {
		eventTime, _ = time.Parse("15:04", oldEvent.EventTime)
	}

	oldEventTime, err := time.Parse("15:04:00", oldEvent.EventTime)
	if err != nil {
		return nil, fmt.Errorf("invalid old_event_time format: %w", err)
	}

	updatedEvent := &model.Event{
		EventName:        helpers.Choose(payload.EventName, oldEvent.EventName),
		EventDate:        helpers.ChooseTime(eventDate, oldEvent.EventDate),
		EventTime:        helpers.ChooseTime(eventTime, oldEventTime).Format("15:04:00"),
		EventLocation:    helpers.Choose(payload.EventLocation, oldEvent.EventLocation),
		EventDescription: helpers.Choose(payload.EventDescription, oldEvent.EventDescription),
		EventOrganizer:   helpers.Choose(payload.EventOrganizer, oldEvent.EventOrganizer),
		EventStatus:      helpers.Choose(payload.EventStatus, oldEvent.EventStatus),
		EventImageURL:    helpers.Choose(imageURL, oldEvent.EventImageURL),
	}

	if err := t.EventRepository.UpdateEvent(id, updatedEvent); err != nil {
		if imageURL != oldEvent.EventImageURL {
			_ = os.Remove(fmt.Sprintf("static/%s", filepath.Base(imageURL)))
		}
		return nil, fmt.Errorf("failed to update event: %w", err)
	}

	response := &dto.EventResponse{
		StatusCode: http.StatusOK,
		Message:    "Event updated successfully",
		Data: []dto.EventData{
			{
				ID:               id,
				EventName:        updatedEvent.EventName,
				EventDate:        updatedEvent.EventDate.Format("2006-01-02"),
				EventTime:        updatedEvent.EventTime,
				EventLocation:    updatedEvent.EventLocation,
				EventDescription: updatedEvent.EventDescription,
				EventOrganizer:   updatedEvent.EventOrganizer,
				EventStatus:      updatedEvent.EventStatus,
				EventImageURL:    fmt.Sprintf("%s%s", t.baseURL, updatedEvent.EventImageURL),
				CreatedAt:        updatedEvent.CreatedAt,
				UpdatedAt:        updatedEvent.UpdatedAt,
			},
		},
	}

	return response, nil
}

func (t *EventService) DeleteEvent(id uint64) (*dto.EventResponse, error) {
	err := t.EventRepository.DeleteEvent(id)
	if err != nil {
		return nil, errs.NotFound("Event not found")
	}

	response := &dto.EventResponse{
		StatusCode: http.StatusOK,
		Message:    "Event deleted successfully",
	}

	return response, nil
}
