package main

type Config struct {
	Env      string `required:"true" default:"development" desc:"production, development"`
	DSN      string `required:"true" default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" desc:"DSN для соединения с базой данных"`
	BindAddr string `required:"true" default:":8080" split_words:"true" desc:"Адрес и порт входящих соединений"`
}
