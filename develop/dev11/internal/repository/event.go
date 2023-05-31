package repository

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dev11/internal/domain/model"
	"dev11/internal/errors"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r *Repository) InsertEvent(ctx context.Context, event *model.Event) *errors.MyErr {
	row := r.sq.Insert("events").
		Columns("event_date",
			"name", "description").
		Values(event.Date,
			event.Name, event.Description).
		Suffix("RETURNING \"id\"").
		QueryRowContext(ctx)

	if err := row.Scan(&event.ID); err != nil {
		return &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while scanning sql row: %w", err),
		}
	}
	return nil
}

func (r *Repository) InsertUserEvent(ctx context.Context, userId int, eventId int) *errors.MyErr {
	if _, err := r.sq.Insert("user_events").
		Columns("user_id", "event_id").
		Values(userId, eventId).
		ExecContext(ctx); err != nil {
		return &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while scanning sql row: %w", err),
		}
	}
	return nil
}

func (r *Repository) UpdateEvent(ctx context.Context, event *model.Event) *errors.MyErr {
	if _, err := r.sq.Update("events").
		SetMap(map[string]interface{}{
			"name":        event.Name,
			"description": event.Description,
			"event_date":  event.Date,
		}).Where(sq.Eq{"id": event.ID}).
		ExecContext(ctx); err != nil {
		return &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while updating event: %w", err),
		}
	}
	return nil
}

func (r *Repository) GetUserEvent(ctx context.Context, eventId int) (*model.UserEvent, *errors.MyErr) {
	row := r.sq.Select("u.id",
		"u.email", "u.first_name", "u.last_name",
		"u.birth_date", "e.id", "e.event_date", "e.name", "e.description").
		From("events e").
		Join("user_events ue ON e.id = ue.event_id").
		Join("users u ON u.id = ue.user_id").
		Where(sq.Eq{"e.id": eventId}).
		QueryRowContext(ctx)

	userEvent := &model.UserEvent{
		ShortUser: model.ShortUser{},
		Event:     model.Event{},
	}
	if err := row.Scan(&userEvent.ShortUser.ID, &userEvent.Email,
		&userEvent.FirstName, &userEvent.LastName,
		&userEvent.BirthDate, &userEvent.Event.ID,
		&userEvent.Date, &userEvent.Name, &userEvent.Description,
	); err != nil {
		return nil, &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while scanning sql row: %w", err),
		}
	}
	return userEvent, nil
}

func (r *Repository) GetUserEvents(ctx context.Context, userId int, start time.Time, end time.Time) ([]model.UserEvent, *errors.MyErr) {
	rows, err := r.sq.Select("u.id",
		"u.email", "u.first_name", "u.last_name",
		"u.birth_date", "e.id", "e.event_date", "e.name", "e.description").
		From("users u").
		Join("user_events ue ON u.id = ue.user_id").
		Join("events e ON e.id = ue.event_id").
		Where(sq.And{sq.Eq{"u.id": userId}, sq.GtOrEq{"e.event_date": start}, sq.Lt{"e.event_date": end}}).
		QueryContext(ctx)
	if err != nil {
		return nil, &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while performing sql request: %w", err),
		}
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			r.logger.Error("error while closing sql rows", zap.Error(err))
		}
	}()

	userEvents := make([]model.UserEvent, 0)
	for rows.Next() {
		userEvent := model.UserEvent{}
		if err = rows.Scan(&userEvent.ShortUser.ID, &userEvent.Email,
			&userEvent.FirstName, &userEvent.LastName,
			&userEvent.BirthDate, &userEvent.Event.ID,
			&userEvent.Date, &userEvent.Name, &userEvent.Description,
		); err != nil {
			return nil, &errors.MyErr{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf("error while scanning sql row: %w", err),
			}
		}
		userEvents = append(userEvents, userEvent)
	}
	return userEvents, nil
}

func (r *Repository) DeleteEvent(ctx context.Context, eventId int) *errors.MyErr {
	if _, err := r.sq.Delete("events").
		Where(sq.Eq{"id": eventId}).ExecContext(ctx); err != nil {
		return &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while deleting event: %w", err),
		}
	}
	return nil
}
