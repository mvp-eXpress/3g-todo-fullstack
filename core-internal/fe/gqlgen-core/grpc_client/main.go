package grpc_client

import (
	"context"
	"log"
	"time"

	v1 "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/be/repository/db/pkg/api/v1"
	"google.golang.org/grpc"
)

var c v1.MongoOPServiceClient
var ctx context.Context

func main() {
	conn, err := grpc.Dial("http://localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c = v1.NewMongoOPServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

func CreateCollection(input *v1.CreateCollectionRequest) (*v1.Collection, error) {

}
