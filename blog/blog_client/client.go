package main

import (
	"blog-with-mongo-grpc/blog/blogpb"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial(os.Getenv("LISTEN_CLIENT"), opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Printf("Created the blog")

	blog := &blogpb.Blog{
		AuthorId: "Pathomphong",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println("Blog has been created: %v", createBlogRes)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)

	// doUnaryWithDeadline(c, 5*time.Second) // should complete
	// doUnaryWithDeadline(c, 1*time.Second) // should timeout
}
