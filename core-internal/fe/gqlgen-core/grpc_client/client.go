package grpc_client

import (
	"context"
	"log"

	v1 "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/be/repository/db/pkg/api/v1"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service v1.MongoOPServiceClient
}

var c v1.MongoOPServiceClient
var ctx context.Context

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()

	c = v1.NewMongoOPServiceClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

// func CreateCollection(input *v1.CreateCollectionRequest) (*v1.Collection, error) {

// }
