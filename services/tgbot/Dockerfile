FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./tgbot-app  cmd/tgbot/main.go

FROM alpine

COPY --from=builder /app/tgbot-app /
COPY config.env /

CMD ["/tgbot-app"]