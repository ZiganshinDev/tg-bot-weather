FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./weather-app  cmd/yandex-weather/main.go

FROM alpine

COPY --from=builder /app/weather-app /
COPY config.env /

EXPOSE 50053

CMD ["/weather-app"]