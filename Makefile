build:
	docker-compose up --build

test:
	go test ./internal/user-management/service -coverprofile=coverage.out


.PHONY: build