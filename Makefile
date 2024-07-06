BUILD = go build -o ./bin/main ./cmd/main/main.go
RUN = ./bin/main

.PHONY: build run gorun vendor  
init: dep dir	
build:
	$(BUILD)
run:
	$(RUN)
gorun: build run
dep:
	go get -u github.com/go-chi/chi/v5
	go get -u github.com/jackc/pgx/v5
	go get -u github.com/jackc/pgx/v5/pgxpool
	go get -u aidanwoods.dev/go-paseto
	go get -u github.com/anthonynsimon/bild
	go get -u github.com/go-playground/validator/v10
	go get -u github.com/gofrs/uuid/v5
	go get -u github.com/kolesa-team/go-webp
	go get -u github.com/oklog/ulid/v2
	go get -u github.com/rs/zerolog
	go get -u github.com/spf13/viper
	go get -u golang.org/x/crypto


