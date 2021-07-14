build:
	$(eval GIT_BRANCH=$(shell git rev-parse --short HEAD))
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-X main.version=$(GIT_BRANCH)" -o .bin/csv2csv-linux .
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags="-X main.version=$(GIT_BRANCH)" -o .bin/csv2csv-mac .

test:
	go test -v -race ./...

lint:
	golangci-lint run ./...