package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gunnlaugurcalvi/grpc-stuff/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to dial, %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	fmt.Println("creating blog")
	blog := &blogpb.Blog{
		AuthorId: "Gulli Blog Artist",
		Title:    "my first blog",
		Content:  "Content stuf fokokokokok",
	}
	response, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("unexpected error, %v", err)
	}

	fmt.Printf("Blog has been created: %v\n", response)

	// READ blog

	fmt.Println("reading blog")
	blogID := response.GetBlog().GetId()
	_, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5e721a4ed8c149ff636fef31"})
	if err != nil {
		fmt.Printf("this is error %v\n", err)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	res, err := c.ReadBlog(context.Background(), readBlogReq)
	if err != nil {
		fmt.Printf("this is error %v", err)
	}

	fmt.Printf("response %v\n", res)

	// update blog
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Gulli Picasso",
		Title:    "advanced gulli ",
		Content:  "btc will moon",
	}
	updateRes, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})

	if err != nil {
		fmt.Printf("Error happened while updating: %v", err)
	}

	fmt.Printf("blog was updated: %v\n", updateRes)

	// delete blog

	fmt.Printf("Delete blog waiting 5sec... \n")

	time.Sleep(5000 * time.Millisecond)
	delRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Printf("Error while deleting blog, %v", err)
	}

	fmt.Printf("Successfully deleted blog with ID: %s, BLOG= %v\n", blogID, delRes)

	// List blog
	fmt.Printf("Listing blogs\n")

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling listBlog RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("somehing happend: %v", err)
		}

		fmt.Printf("blog: %s\n", res.GetBlog())
	}
}
