package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"dev11/internal/api"
	"dev11/internal/db"
	"dev11/internal/domain/service"
	"dev11/internal/repository"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/oklog/run"
	"go.uber.org/zap"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	var (
		ctx, ctxCancel = context.WithCancel(context.Background())
		cfg            = new(Config)
		logger         *zap.Logger
		err            error
	)
	defer ctxCancel()

	if cfg.Env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(fmt.Errorf("[main] Failed to initialize logger: %w", err))
	}
	sugaredLogger := logger.Sugar()

	if err = envconfig.Process("APP", cfg); err != nil {
		sugaredLogger.Fatal(err.Error())
	}

	conn, err := db.ConnectPostgreSQL(ctx, cfg.DSN)
	if err != nil {
		panic(fmt.Errorf("[main] Unable to open connection with DB: %w", err))
	}

	g := &run.Group{}

	repo := repository.NewRepository(conn, sugaredLogger)
	svc := service.NewService(repo)
	api.New(
		api.WithLogger(sugaredLogger),
		api.WithService(svc),
		api.WithBindAddress(cfg.BindAddr)).Run(g)

	ctx, cancel := context.WithCancel(context.Background())
	g.Add(func() error {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

		sugaredLogger.Info("[signal-watcher] Started")

		select {
		case sig := <-shutdown:
			return fmt.Errorf("[signal-watcher] Terminated with signal: %s", sig.String())
		case <-ctx.Done():
			return nil
		}
	}, func(err error) {
		cancel()
		sugaredLogger.Error("[signal-watcher] Gracefully shutdown application", zap.Error(err))
	})

	sugaredLogger.Error("[main] Successful shutdown", zap.Error(g.Run()))
}
