
all: signer

#Build dir specified from terminal.
signer: gofmt lint test
	go build -o $(BUILDDIR)/$@ .

gofmt:
	go fmt .

lint:
	golangci-lint run

test:
	go test ./...
