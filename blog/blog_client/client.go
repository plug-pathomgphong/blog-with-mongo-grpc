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

	blogId := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("reading the blog.")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "61c46457d7d796381a59580c"})
	if err2 != nil {
		fmt.Println("Error happened while reading2: %v", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogres, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Println("Error happened while reading: %v", readBlogErr)
	}
	fmt.Println("Blog was read: %v", readBlogres)

}
