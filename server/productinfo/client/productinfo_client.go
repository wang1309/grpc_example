package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "productinfo/service/ecommerce"
	"time"
)

const address = "localhost:50051"

func main()  {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "Apple iphone 11"
	description := "Meet people"

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}

	log.Printf("Product ID: %s added successfully", r.Value)
}
