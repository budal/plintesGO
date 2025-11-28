FROM golang:1.23-alpine AS development

RUN apk add --no-cache git curl build-base bash

ENV GOBIN=/usr/local/bin
RUN go install github.com/cosmtrek/air@v1.52.0

WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

COPY . .

CMD ["air"]


# -----------------------------------------


FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o server ./cmd/api


# -----------------------------------------


FROM alpine:3.19

RUN adduser -D -u 1000 plintes

WORKDIR /app

COPY --from=builder /app/server .

USER plintes

EXPOSE 3000

CMD ["./server"]
