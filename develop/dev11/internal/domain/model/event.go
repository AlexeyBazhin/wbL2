package model

import (
	"time"
)

type (
	Event struct {
		ID          int     `json:"id"`
		Date        time.Time `json:"date"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	}

	UserEvent struct {
		ShortUser `json:"user"`
		Event     `json:"event"`
	}
)
