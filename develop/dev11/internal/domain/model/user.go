package model

import (
	"time"
)

type (
	User struct {
		ShortUser
		HashedPassword string `json:"hashedPassword"`
	}

	ShortUser struct {
		ID        int       `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		BirthDate time.Time `json:"birthDate"`
		Email     string    `json:"email"`
	}
)
