package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	v1 "github.com/mvp-eXpress/3g-todo-fullstack/be/repository/db/pkg/api/v1"
	"google.golang.org/grpc"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)

	pfx := t.Format(time.RFC3339Nano)

	// Call Create
	req1 := v1.CreateRequest{

		Todo: &v1.Todo{
			Title:       "title (" + pfx + ")",
			Description: &wrappers.StringValue{Value: "description (" + pfx + ")"},
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	// id := res1.Id
	// Call Create

	// Read
	// req2 := &v1.GetRequest{
	// 	Id: "aaa",
	// }
	// res2, err := c.Get(ctx, req2)
	// if err != nil {
	// 	log.Fatalf("Read failed: %v", err)
	// }
	// log.Printf("Read result: <%+v>\n\n", res2)
	// Read

	// Update
	// req3 := v1.UpdateRequest{

	// 	Todo: &v1.Todo{
	// 		Id:          "ble id",
	// 		Title:       "ble title",
	// 		Description: &wrappers.StringValue{Value: "ble desc"},
	// 	},
	// }
	// res3, err := c.Update(ctx, &req3)
	// if err != nil {
	// 	log.Fatalf("Update failed: %v", err)
	// }
	// log.Printf("Update result: <%+v>\n\n", res3)
	// Update

	// Call ReadAll
	// req4 := v1.Empty{}
	// res4, err := c.List(ctx, &req4)
	// if err != nil {
	// 	log.Fatalf("ReadAll failed: %v", err)
	// }
	// log.Printf("ReadAll result: <%+v>\n\n", res4)
	// Call ReadAll

	// Delete
	// req5 := v1.DeleteRequest{

	// 	Id: id,
	// }
	// res5, err := c.Delete(ctx, &req5)
	// if err != nil {
	// 	log.Fatalf("Delete failed: %v", err)
	// }
	// log.Printf("Delete result: <%+v>\n\n", res5)
	// Delete

}
