package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ZiganshinDev/tg-bot-weather/services/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

type Server struct {
	s *grpc.Server
}

type server struct {
	pb.UnimplementedWeatherServiceServer
	w WeatherGetter
}

type WeatherGetter interface {
	GetWeather(float64, float64) (string, error)
}

func (s *server) GetWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherReply, error) {
	weather, err := s.w.GetWeather(in.GetLatitude(), in.GetLongitude())
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	log.Printf("Recived: %v, %v", in.GetLatitude(), in.GetLongitude())
	return &pb.WeatherReply{Conditions: weather}, nil
}

func New() *Server {
	s := grpc.NewServer()

	return &Server{s: s}
}

func (s *Server) Start(weather WeatherGetter) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterWeatherServiceServer(s.s, &server{w: weather})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
