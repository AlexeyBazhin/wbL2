package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type (
	CreateEventReq struct {
		UserId      int       `json:"user_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Date        time.Time `json:"date"`
	}
	UpdateEventReq struct {
		Id          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Date        time.Time `json:"date"`
	}
)

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/create_event")
	eventReq := &CreateEventReq{}
	if err := json.NewDecoder(r.Body).Decode(eventReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errResponse{err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}
	defer r.Body.Close()

	userEvent, err := s.svc.CreateEvent(r.Context(), eventReq)
	if err != nil {
		w.WriteHeader(err.Code)
		if err := json.NewEncoder(w).Encode(errResponse{err.Err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resultResponse{userEvent}); err != nil {
		s.logger.Error(err)
	}
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/update_event")
	eventReq := &UpdateEventReq{}
	if err := json.NewDecoder(r.Body).Decode(eventReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errResponse{err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}
	defer r.Body.Close()

	userEvent, err := s.svc.UpdateEvent(r.Context(), eventReq)
	if err != nil {
		w.WriteHeader(err.Code)
		if err := json.NewEncoder(w).Encode(errResponse{err.Err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resultResponse{userEvent}); err != nil {
		s.logger.Error(err)
	}
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/delete_event")
	eventReq := &struct {
		EventId int `json:"event_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(eventReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errResponse{err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}
	defer r.Body.Close()

	if err := s.svc.DeleteEvent(r.Context(), eventReq.EventId); err != nil {
		w.WriteHeader(err.Code)
		if err := json.NewEncoder(w).Encode(errResponse{err.Err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resultResponse{"deleted"}); err != nil {
		s.logger.Error(err)
	}
}

func (s *Server) eventsForDay(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/events_for_day")
	s.eventsFor(w, r, "day")
}

func (s *Server) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/events_for_week")
	s.eventsFor(w, r, "week")
}

func (s *Server) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(r.Method, "/events_for_month")
	s.eventsFor(w, r, "month")
}

func (s *Server) eventsFor(w http.ResponseWriter, r *http.Request, rangeType string) {
	userId, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errResponse{err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	date, err := time.Parse(time.RFC3339, r.FormValue("date"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errResponse{err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	userEvents, getErr := s.svc.GetUserEvents(r.Context(), userId, date, rangeType)
	if getErr != nil {
		w.WriteHeader(getErr.Code)
		if err := json.NewEncoder(w).Encode(errResponse{getErr.Err.Error()}); err != nil {
			s.logger.Error(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resultResponse{userEvents}); err != nil {
		s.logger.Error(err)
	}
}
