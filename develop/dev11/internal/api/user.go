package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	CreateUserReq struct {
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		BirthDate time.Time `json:"birthDate"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
	}
)

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/register")
	userReq := &CreateUserReq{}
	if err := json.NewDecoder(r.Body).Decode(userReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errResponse{err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}
	defer r.Body.Close()

	user, err := s.svc.CreateUser(r.Context(), userReq)
	if err != nil {
		w.WriteHeader(err.Code)
		if err := json.NewEncoder(w).Encode(errResponse{err.Err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resultResponse{user}); err != nil {
		s.logger.Error(err)
	}
}
