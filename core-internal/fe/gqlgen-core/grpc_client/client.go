package grpc_client

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/be/repository/db/pkg/api/v1"
	gqlgen_core "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/fe/gqlgen-core"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service v1.MongoOPServiceClient
}

var c v1.MongoOPServiceClient
var ctx context.Context

func NewClient(url string) (*Client, error) {
	log.Printf("grpc_client::NewClient::url: %v", url)

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = v1.NewMongoOPServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) GetCollections(ctx context.Context) ([]*gqlgen_core.Collection, error) {
	res, err := c.service.ListCollections(ctx, &v1.Empty{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	collections := []*gqlgen_core.Collection{}
	for _, coll := range res.Collections {
		collections = append(collections, &gqlgen_core.Collection{
			ID:   coll.Id,
			Host: coll.Host,
			Name: coll.Name,
		})
	}
	fmt.Printf("grpc_client::GetCollections::Collections: %v", collections)
	return collections, nil
}

func (c *Client) CreateCollection(ctx context.Context, in gqlgen_core.CreateCollectionInput) (*gqlgen_core.Collection, error) {
	var fvs []*v1.FieldValue
	for _, inFv := range in.FieldValues {
		fvs = append(fvs, &v1.FieldValue{
			FieldName: *inFv.FieldName,
			IsIndexed: *inFv.IsIndexed,
			IsUnique:  *inFv.IsUnique,
		})
	}
	res, err := c.service.CreateCollection(ctx, &v1.CreateCollectionRequest{
		Collection: &v1.Collection{
			Name:        in.Name,
			Host:        in.Host,
			FieldValues: fvs,
		},
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &gqlgen_core.Collection{
		ID:   res.Collection.Id,
		Name: in.Name,
		Host: in.Host,
	}, nil
}
