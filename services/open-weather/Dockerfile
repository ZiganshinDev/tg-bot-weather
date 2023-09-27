FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./location-app  cmd/open-weather/main.go

FROM alpine

COPY --from=builder /app/location-app /
COPY config.env /

EXPOSE 50052

CMD ["/location-app"]