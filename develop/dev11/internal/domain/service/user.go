package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dev11/internal/api"
	"dev11/internal/domain/model"
	"dev11/internal/errors"

	"golang.org/x/crypto/bcrypt"
)

func hashPass(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}

func (s *service) CreateUser(ctx context.Context, userReq *api.CreateUserReq) (*model.User, *errors.MyErr) {
	user := &model.User{
		ShortUser: model.ShortUser{
			FirstName: userReq.FirstName,
			LastName:  userReq.LastName,
			Email:     userReq.Email,
		},
		HashedPassword: hashPass(userReq.Password),
	}
	if userReq.BirthDate.Before(time.Now().AddDate(-100, 0, 0)) {
		return nil, &errors.MyErr{
			Code: http.StatusServiceUnavailable,
			Err:  fmt.Errorf("invalid birth date"),
		}
	}
	user.BirthDate = userReq.BirthDate

	return user, s.repo.InsertUser(ctx, user)
}
