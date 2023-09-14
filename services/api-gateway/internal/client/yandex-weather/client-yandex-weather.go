package yandex_weather

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ZiganshinDev/tg-bot-weather/services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50053", "the address to connect to")
)

func GetWeather(lan float64, lon float64) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWeatherServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetWeather(ctx, &pb.WeatherRequest{Latitude: lan, Longitude: lon})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	return r.GetConditions()
}
