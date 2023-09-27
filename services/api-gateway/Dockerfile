FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./api-gateway-app  cmd/api-gateway/main.go

FROM alpine

COPY --from=builder /app/api-gateway-app /

EXPOSE 50051

CMD ["/api-gateway-app"]