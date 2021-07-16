build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .bin/top10-linux .
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o .bin/top10-mac .

test:
	go test -v -race ./...

lint:
	golangci-lint run ./...