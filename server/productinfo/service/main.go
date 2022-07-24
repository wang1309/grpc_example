package main

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "productinfo/service/ecommerce"
)

const port = ":50051"

type Server struct {
	productMap map[string]*pb.Product
}


func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductId, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("generrate product id fault, err: %+v", err.Error())
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in
	return &pb.ProductId{Value: in.Id}, nil
}

func (s *Server) GetProduct(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, nil
	}

	return nil, fmt.Errorf("Product does not exist, id: %v", in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()


	pb.RegisterProductInfoServer(s, &Server{})

	log.Printf("Starting gRpc listener on port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to server: %v", err)
	}
}
