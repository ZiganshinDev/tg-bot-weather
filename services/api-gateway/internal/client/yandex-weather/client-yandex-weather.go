package yandex_weather

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/adapter/weatherclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "weather-app:50053", "the address to connect to")
)

type Client struct {
	conn *grpc.ClientConn
}

func New() (*Client, error) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) GetWeather(lan float64, lon float64) string {
	cl := pb.NewWeatherServiceClient(c.conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cl.GetWeather(ctx, &pb.WeatherRequest{Latitude: lan, Longitude: lon})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	return r.GetConditions()
}
