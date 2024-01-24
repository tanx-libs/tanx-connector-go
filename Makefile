all:
	go test -v -cover -bench=. -benchmem ./client

.PHONY: all