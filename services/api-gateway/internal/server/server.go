package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ZiganshinDev/tg-bot-weather/services/api-gateway/adapter/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	s *grpc.Server
}

type server struct {
	pb.UnimplementedApiGatewayServiceServer
	w WeatherGetter
}

type WeatherGetter interface {
	GetWeather(string) string
}

func (s *server) GetWeather(city string) string {
	return s.w.GetWeather(city)
}

func (s *server) GetUserCity(ctx context.Context, in *pb.UserCityRequest) (*pb.UserWeatherReply, error) {
	city := in.GetCityName()

	log.Printf("Recived: %v", in.GetCityName())
	return &pb.UserWeatherReply{Weather: s.GetWeather(city)}, nil
}

func New() *Server {
	s := grpc.NewServer()
	return &Server{s: s}
}

func (s *Server) Run(weather WeatherGetter) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterApiGatewayServiceServer(s.s, &server{w: weather})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
