package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dev11/internal/domain/model"
	"dev11/internal/errors"

	"github.com/oklog/run"
	"go.uber.org/zap"
)

type (
	errResponse struct {
		ErrField string `json:"error"`
	}
	resultResponse struct {
		Result interface{} `json:"result"`
	}
	Server struct {
		*http.Server
		logger *zap.SugaredLogger
		svc    Service

		bindAddr string
	}

	Service interface {
		userService
		eventService
	}

	userService interface {
		CreateUser(ctx context.Context, userReq *CreateUserReq) (*model.User, *errors.MyErr)
	}
	eventService interface {
		CreateEvent(ctx context.Context, eventReq *CreateEventReq) (*model.UserEvent, *errors.MyErr)
		UpdateEvent(ctx context.Context, eventReq *UpdateEventReq) (*model.UserEvent, *errors.MyErr)
		DeleteEvent(ctx context.Context, eventId int) *errors.MyErr
		GetUserEvents(ctx context.Context, userId int, date time.Time, rangeType string) ([]model.UserEvent, *errors.MyErr)
	}
	OptionFunc func(s *Server)
)

func New(opts ...OptionFunc) *Server {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/register", s.register)
	mux.HandleFunc("/create_event", s.createEvent)
	mux.HandleFunc("/update_event", s.updateEvent)
	mux.HandleFunc("/delete_event", s.deleteEvent)
	mux.HandleFunc("/events_for_day", s.eventsForDay)
	mux.HandleFunc("/events_for_week", s.eventsForWeek)
	mux.HandleFunc("/events_for_month", s.eventsForMonth)

	s.Server = &http.Server{
		Addr:         s.bindAddr,
		Handler:      mux,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
	return s
}

func (s *Server) Run(g *run.Group) {
	g.Add(func() error {
		s.logger.Info("[http-server] started")
		s.logger.Info(fmt.Sprintf("listening on %v", s.Addr))
		return s.ListenAndServe()
	}, func(err error) {
		s.logger.Error("[http-server] terminated", zap.Error(err))

		ctx, cancel := context.WithTimeout(context.Background(), 30)
		defer cancel()

		s.logger.Error("[http-server] stopped", zap.Error(s.Shutdown(ctx)))
	})
}

func WithLogger(logger *zap.SugaredLogger) OptionFunc {
	return func(s *Server) {
		s.logger = logger
	}
}

func WithBindAddress(bindAddr string) OptionFunc {
	return func(s *Server) {
		s.bindAddr = bindAddr
	}
}

func WithService(svc Service) OptionFunc {
	return func(s *Server) {
		s.svc = svc
	}
}
