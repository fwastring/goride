.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
