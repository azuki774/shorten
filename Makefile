COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''
SHELL=/bin/bash

.PHONY: build bin bin-linux-amd64 test clean start

build:
	docker build -t shorten-redirector:latest -f build/Dockerfile-redirector .
	docker build -t shorten-register:latest -f build/Dockerfile-register .

bin:
	go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\"" -o build/bin/ ./...

bin-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\"" -o build/bin/ ./...

start:
	docker compose -f deployment/compose-local-db.yml up -d
	docker compose -f deployment/compose-local-svc.yml up -d

stop:
	docker compose -f deployment/compose-local-svc.yml down
	docker compose -f deployment/compose-local-db.ym down

restart:
	docker compose -f deployment/compose-local-svc.yml down
	docker compose -f deployment/compose-local-svc.yml up -d

test:
	gofmt -l .
	go vet ./...
	staticcheck ./...
	go test -v ./...  | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

clean:
	docker compose -f deployment/compose-local-svc.yml down
	docker compose -f deployment/compose-local-db.yml down
	rm -rf build/bin/shorten-redirector
	rm -rf build/bin/shorten-register
