package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ZiganshinDev/tg-bot-weather/services/proto"
	yandex_weather "github.com/ZiganshinDev/tg-bot-weather/services/yandex-weather/internal/service/yandex-weather"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

type server struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *server) GetWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherReply, error) {
	yandex, err := yandex_weather.New()
	if err != nil {
		log.Fatalln("failed to init yandex-weather")
	}

	weather, err := yandex.GetWeather(in.GetLatitude(), in.GetLongitude())
	if err != nil {
		log.Printf("failed to get weather, %v", err)
	}

	log.Printf("Recived: %v, %v", in.GetLatitude(), in.GetLongitude())
	return &pb.WeatherReply{Conditions: weather}, nil
}

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
