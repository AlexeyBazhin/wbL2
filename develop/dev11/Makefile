GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
OUT := bin/dev11
OUT_PATH=$(shell pwd)/bin/$(GOOS)_$(GOARCH)

clean:
	rm -rf ./bin/*
.PHONY: clean

clean.bin: ## remove $(OUT_PATH) directory
	rm -rf $(OUT_PATH)
.PHONY: clean.bin

build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $(OUT) ./cmd/dev11
.PHONY: build

gen:
	go generate ./...
.PHONY: gen

dbuild:
	docker compose -f ../dev11/docker-compose.yml build --no-cache
.PHONY: dbuild

dstart:
	docker compose -f ../dev11/docker-compose.yml up --force-recreate
.PHONY: dstart

run: dbuild dstart
.PHONY: run

lint: ## run linters for project
	$(OUT_PATH)/golangci-lint run
.PHONY: lint

prepare: clean install.tools ## performs steps needed before first build
.PHONY: prepare