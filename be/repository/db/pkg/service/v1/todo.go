package v1

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/golang/protobuf/ptypes"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	v1 "github.com/mvp-eXpress/3g-todo-fullstack/be/repository/db/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type todoServiceServer struct {
	coll *mongo.Collection
}

// NewToDoServiceServer creates Todo grpc service
func NewToDoServiceServer(coll *mongo.Collection) v1.TodoServiceServer {
	return &todoServiceServer{coll: coll}
}

func (s *todoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.Todo, error) {
	// res, err := s.coll.InsertOne(ctx, &v1.Todo{
	// 	Id:          "1",
	// 	Title:       "Title A",
	// 	Description: "Description A",
	// })

	t := time.Now()
	createdAt, err := ptypes.TimestampProto(t)
	if err != nil {
		log.Fatalf("failed to convert golang Time to protobuf Timestamp: %#v", err)
	}

	req.Todo.IsCompleted = false
	req.Todo.IsCompletedWrapped = &wrappers.BoolValue{Value: false}
	req.Todo.CreatedAtWrapped = createdAt
	res, err := s.coll.InsertOne(ctx, req.Todo)
	if err != nil {
		log.Fatalf("insert data into collection <todo.todo>: %#v", err)
	}
	id := res.InsertedID
	log.Printf("inserted new item with id=%v successfully", id)

	idString := id.(primitive.ObjectID).Hex()
	log.Printf("idString=%v ", idString)

	resp := req.Todo
	resp.Id = idString

	return resp, nil
}

func (s *todoServiceServer) Get(ctx context.Context, req *v1.GetRequest) (*v1.Todo, error) {
	var out *v1.Todo
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Printf("ObjectIDFromHex failed for %v with error: %v", req.Id, err)
	}
	filter := bson.D{{Key: "_id", Value: objID}}

	err = s.coll.FindOne(ctx, filter).Decode(&out)
	if err != nil {
		fmt.Printf("failed to read data (id=%v) from collection <todo.todo>: %#v", req.Id, err)
	}

	return out, err
}

func (s *todoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.Todo, error) {
	var out *v1.Todo
	objID, err := primitive.ObjectIDFromHex(req.Todo.Id)
	if err != nil {
		fmt.Printf("ObjectIDFromHex failed for %v with error: %v", req.Todo.Id, err)
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	err = s.coll.FindOneAndReplace(ctx, filter, req.Todo).Decode(&out)

	return out, err

}

func (s *todoServiceServer) List(ctx context.Context, req *v1.Empty) (*v1.TodoList, error) {
	list := []*v1.Todo{}
	td := &v1.Todo{
		Id:          "3",
		Title:       "Title from list",
		Description: &wrappers.StringValue{Value: "Description from list"},
	}
	list = append(list, td)
	return &v1.TodoList{
		Todos: list,
	}, nil
}

func (s *todoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.Empty, error) {
	return &v1.Empty{}, nil
}
