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
	port = flag.Int("port", 50052, "The server port")
)

type Server struct {
	s *grpc.Server
}

type server struct {
	pb.UnimplementedCityServiceServer
	c CoordinatesGetter
}

type CoordinatesGetter interface {
	GetCoordinates(string) (float64, float64, error)
}

func (s *server) GetCityCoordinates(ctx context.Context, in *pb.CityRequest) (*pb.CoordinatesReply, error) {
	lat, lon, err := s.c.GetCoordinates(in.GetCityName())
	if err != nil {
		log.Println(err)
	}

	log.Printf("Recived: %v", in.GetCityName())
	return &pb.CoordinatesReply{Latitude: lat, Longitude: lon}, nil
}

func New() *Server {
	s := grpc.NewServer()

	return &Server{s: s}
}

func (s *Server) Start(c CoordinatesGetter) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterCityServiceServer(s.s, &server{c: c})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
