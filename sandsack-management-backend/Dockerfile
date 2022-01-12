FROM golang:1.13.1-alpine3.10 as prepare

WORKDIR /source

COPY go.mod .
COPY go.sum .

RUN go mod download

#
# STAGE 2: build
#
FROM prepare AS build

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/server -v ./main.go

#
# STAGE 3: run
#
FROM scratch as run

COPY --from=build /source/bin/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]