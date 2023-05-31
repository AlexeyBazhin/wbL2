package domain

import (
	"context"
	"time"

	"dev11/internal/domain/model"
	"dev11/internal/errors"
)

type (
	Repository interface {
		userRepo
		eventRepo
	}
	userRepo interface {
		InsertUser(ctx context.Context, user *model.User) *errors.MyErr
		GetUser(ctx context.Context, id int) (*model.User, *errors.MyErr)
	}
	eventRepo interface {
		InsertEvent(ctx context.Context, event *model.Event) *errors.MyErr
		InsertUserEvent(ctx context.Context, userId, eventId int) *errors.MyErr
		UpdateEvent(ctx context.Context, event *model.Event) *errors.MyErr
		DeleteEvent(ctx context.Context, eventId int) *errors.MyErr
		GetUserEvent(ctx context.Context, eventId int) (*model.UserEvent, *errors.MyErr)
		GetUserEvents(ctx context.Context, userId int, start time.Time, end time.Time) ([]model.UserEvent, *errors.MyErr)
	}
)
