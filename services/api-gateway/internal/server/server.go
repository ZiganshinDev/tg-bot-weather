package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	open_weather "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/internal/client/open-weather"
	yandex_weather "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/internal/client/yandex-weather"
	pb "github.com/ZiganshinDev/tg-bot-weather/services/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedApiGatewayServiceServer
}

func (s *server) GetUserCity(ctx context.Context, in *pb.UserCityRequest) (*pb.UserWeatherReply, error) {
	yandex_weather.GetWeather(open_weather.GetCityCoordinates(in.GetCityName()))

	log.Printf("Recived: %v", in.GetCityName())
	return &pb.UserWeatherReply{Weather: yandex_weather.GetWeather(open_weather.GetCityCoordinates(in.GetCityName()))}, nil
}

func New() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterApiGatewayServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
