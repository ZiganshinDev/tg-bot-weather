package client

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/ZiganshinDev/tg-bot-weather/services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type Client struct {
	conn *grpc.ClientConn
}

func New() (*Client, error) {
	const op = "service.client.New"

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) GetWeather(city string) string {
	cl := pb.NewApiGatewayServiceClient(c.conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := cl.GetUserCity(ctx, &pb.UserCityRequest{CityName: city})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	return r.GetWeather()
}
