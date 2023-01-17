# start from go 1.17
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

# install go module for building the application
RUN go mod download

# copy all source file to docker for building
COPY . .

# run go generate to make sure the generated file from open api specification is always up to date
RUN go generate ./...

# building the go application
RUN go build -o bin/app cmd/app/main.go

# expose 3000 port for accessing
EXPOSE 3000

# binary for executing go application when docker container start
CMD ["./bin/app"]
