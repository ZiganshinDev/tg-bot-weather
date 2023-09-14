package open_weather

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
	addrs = flag.String("addrs", "localhost:50052", "the address to connect to")
)

func GetCityCoordinates(city string) (float64, float64) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCityServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCityCoordinates(ctx, &pb.CityRequest{CityName: city})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	return r.GetLatitude(), r.GetLongitude()
}
