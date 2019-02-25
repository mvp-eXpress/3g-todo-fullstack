package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	// "github.com/mongodb/mongo-go-driver/bson"

	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	// "github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mvp-eXpress/3g-todo-fullstack/core-internal/be/repository/db/pkg/protocol/grpc"
	v1 "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/be/repository/db/pkg/service/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	GRPCPort        string
	MongoHost       string
	MongoDB         string
	MongoCollection string
	MongoUser       string
	MongoPassword   string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "grpc port to connect to")
	flag.StringVar(&cfg.MongoHost, "mongo-host", "", "mongo hostname")
	flag.StringVar(&cfg.MongoDB, "mongo-db", "", "mongo database name")
	flag.StringVar(&cfg.MongoCollection, "mongo-collection", "", "mongo collection name")
	flag.StringVar(&cfg.MongoUser, "mongo-user", "", "mongo username")
	flag.StringVar(&cfg.MongoPassword, "mongo-password", "", "mongo password")
	flag.Parse()

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s",
		cfg.MongoUser,
		cfg.MongoPassword,
		cfg.MongoHost,
		cfg.MongoDB)

	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()
	// client, err := mongo.NewClient((options.Client().SetRegistry(reg)).ApplyURI(mongoURI))
	// client, err := mongo.NewClientWithOptions(mongoURI, &options.ClientOptions{
	// 	Registry: reg,
	// })
	client, err := mongo.NewClient(options.Client().
		ApplyURI(mongoURI).
		SetRegistry(reg),
	)
	if err != nil {
		log.Fatalf("failed to create new MongoDB client: %#v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect client
	if err = client.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %#v", err)
	}

	log.Printf("connected successfully")

	// Get collection from database
	coll := client.Database(cfg.MongoDB).Collection(cfg.MongoCollection)
	// coll := client.Database("todo").Collection("todo")

	v1API := v1.NewMongoOpServiceServer(coll)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)

}
