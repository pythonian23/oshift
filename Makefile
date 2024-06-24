CODE = *.go

.PHONY: all format test

all:
	$(MAKE) format
	$(MAKE) test

format: $(CODE) go.mod
	goimports -w $(CODE)
	gofmt -w $(CODE)

test: $(CODE)
	go test -cover

go.mod: $(CODE)
	go mod tidy
