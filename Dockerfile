FROM golang:1.16.0-alpine3.13 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /app

COPY . ./

RUN go mod download

#RUN CGO_ENABLED=1 go build -ldflags '-linkmode external -w -extldflags "-static"' -o main .
RUN GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine:3.13

WORKDIR /app


COPY --from=builder /app/main /app/main

CMD ["/app/main"]

