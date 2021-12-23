package main

import (
	"blog-with-mongo-grpc/blog/blogpb"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
	blogpb.BlogServiceServer
	// this also works and is suggested in more places
	// blogpb.UnimplementedBlogServiceServer
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"auther_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog Service Started")
	fmt.Println("Connecting to mongoDB")

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_SERVER")))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", os.Getenv("LISTEN_TCP"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(grpcServer, &server{})

	go func() {
		fmt.Println("Starting Server...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	fmt.Println("Stopping the Server..")
	grpcServer.Stop()
	fmt.Println("Closing the Server.")
	lis.Close()
	fmt.Println("Closing MongoDB Connection.")
	client.Disconnect(context.TODO())
	fmt.Println("End of program.")
}
