package v1

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	v1 "github.com/mvp-eXpress/3g-todo-fullstack/core-internal/be/repository/db/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type mongoOpServiceServer struct {
	coll *mongo.Collection
}

// NewMongoOpServiceServer creates mongoOP grpc service
func NewMongoOpServiceServer(coll *mongo.Collection) v1.MongoOPServiceServer {
	return &mongoOpServiceServer{coll: coll}
}

func (s *mongoOpServiceServer) ListCollections(ctx context.Context, req *v1.Empty) (*v1.ListCollectionsResponse, error) {
	list := []*v1.Collection{}
	// f:=interface{}
	// fo := options.Find()
	cur, err := s.coll.Find(ctx, bson.D{}, nil)
	if err != nil {
		// log.Fatalf("ListCollections: %#v", err)
		// return nil, err

		return nil, fmt.Errorf("readTasks: couldn't list all to-dos: %v", err)

	}

	log.Printf("curr: %v", cur)
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		myColl := &v1.Collection{}
		err := cur.Decode(myColl)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
		list = append(list, myColl)

	}

	return &v1.ListCollectionsResponse{
		Collections: list,
	}, nil
}

func (s *mongoOpServiceServer) TestWorld(ctx context.Context, req *v1.TestWorldRequest) (*v1.TestWorldResponse, error) {
	log.Println("TestWorld")
	fmt.Printf(req.String())
	return &v1.TestWorldResponse{
		Msg: "HEEELLOOO",
	}, nil
}

func (s *mongoOpServiceServer) CreateCollection(ctx context.Context, req *v1.CreateCollectionRequest) (*v1.CreateCollectionResponse, error) {
	log.Printf(req.String())
	res, err := s.coll.InsertOne(ctx, req.Collection)
	if err != nil {
		log.Fatalf("insert data into collection <todo.todo>: %#v", err)
		return nil, err
	}
	id := res.InsertedID
	idString := id.(primitive.ObjectID).Hex()
	resp := &v1.CreateCollectionResponse{
		Collection: req.Collection,
	}
	resp.Collection.Id = idString
	return resp, nil
}
