package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	open_weather "github.com/ZiganshinDev/tg-bot-weather/services/open-weather/internal/service/open-weather"
	pb "github.com/ZiganshinDev/tg-bot-weather/services/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnimplementedCityServiceServer
}

func (s *server) GetCityCoordinates(ctx context.Context, in *pb.CityRequest) (*pb.CoordinatesReply, error) {
	client, err := open_weather.New()
	if err != nil {
		log.Fatal(err)
	}

	lat, lon, err := client.GetCoordinates(in.GetCityName())
	if err != nil {
		log.Println(err)
	}

	log.Printf("Recived: %v", in.GetCityName())
	return &pb.CoordinatesReply{Latitude: lat, Longitude: lon}, nil
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
	pb.RegisterCityServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
