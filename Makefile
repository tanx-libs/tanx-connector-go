test:
	go test -v ./client

coverage:
	go test -cover ./client

all:
	go test -v -cover -bench=. -benchmem ./client

.PHONY: all