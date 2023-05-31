package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dev11/internal/api"
	"dev11/internal/domain/model"
	"dev11/internal/errors"
)

func (s *service) CreateEvent(ctx context.Context, eventReq *api.CreateEventReq) (*model.UserEvent, *errors.MyErr) {
	event := &model.Event{
		Name:        eventReq.Name,
		Description: eventReq.Description,
	}
	if eventReq.Date.Before(time.Now()) {
		return nil, &errors.MyErr{
			Code: http.StatusServiceUnavailable,
			Err:  fmt.Errorf("invalid event date (before now)"),
		}
	}
	event.Date = eventReq.Date

	if err := s.repo.InsertEvent(ctx, event); err != nil {
		return nil, err
	}

	user, err := s.repo.GetUser(ctx, eventReq.UserId)
	if err != nil {
		return nil, err
	}
	userEvent := &model.UserEvent{
		ShortUser: user.ShortUser,
		Event:     *event,
	}
	return userEvent, s.repo.InsertUserEvent(ctx, user.ID, event.ID)
}

func (s *service) UpdateEvent(ctx context.Context, eventReq *api.UpdateEventReq) (*model.UserEvent, *errors.MyErr) {
	event := &model.Event{
		ID:          eventReq.Id,
		Name:        eventReq.Name,
		Description: eventReq.Description,
	}
	if eventReq.Date.Before(time.Now()) {
		return nil, &errors.MyErr{
			Code: http.StatusServiceUnavailable,
			Err:  fmt.Errorf("invalid event date (before now)"),
		}
	}
	event.Date = eventReq.Date

	if err := s.repo.UpdateEvent(ctx, event); err != nil {
		return nil, err
	}

	return s.repo.GetUserEvent(ctx, event.ID)
}

func (s *service) DeleteEvent(ctx context.Context, eventId int) *errors.MyErr {
	return s.repo.DeleteEvent(ctx, eventId)
}

func (s *service) GetUserEvents(ctx context.Context, userId int, date time.Time, rangeType string) ([]model.UserEvent, *errors.MyErr) {
	var start, end time.Time

	switch rangeType {
	case "day":
		start = date
		end = date.AddDate(0, 0, 1)
	case "week":
		start = date.AddDate(0, 0, -1*int(date.Weekday()))
		end = date.AddDate(0, 0, 7-int(date.Weekday()))
	case "month":
		start = date.AddDate(0, 0, -1*date.Day()+1)
		end = date.AddDate(0, 1, -1*date.Day()+1)
	}

	return s.repo.GetUserEvents(ctx, userId, start, end)
}
