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
		log.Fatalf("could not connect: %v \n", err)
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
		log.Fatalf("Unexpected error: %v \n", err)
	}
	fmt.Println("Blog has been created: %v \n", createBlogRes)

	blogId := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("reading the blog.")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "61c46457d7d796381a59580c"})
	if err2 != nil {
		fmt.Println("Error happened while reading2: %v \n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogres, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Println("Error happened while reading: %v \n", readBlogErr)
	}
	fmt.Println("Blog was read: %v \n", readBlogres)

	// update blog
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Pathomphong Menthong",
		Title:    "My first blog (edited)",
		Content:  "Content of the first blog, Awesome!",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})

	if updateErr != nil {
		fmt.Printf("Error happened while updateing: %v \n", updateErr)
	}
	fmt.Printf("Blog was Updated: %v \n", updateRes)

	// Delete blog
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogId})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v \n", deleteRes)
}
