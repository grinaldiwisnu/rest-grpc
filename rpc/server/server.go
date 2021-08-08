package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"rest-grpc/api/config"
	"rest-grpc/api/model"
	"rest-grpc/api/repository"
	"rest-grpc/api/service"
	"rest-grpc/rpc/proto"
)

var ProductService service.ProductService

type server struct {
	proto.ProductServiceServer
}

func (s *server) CreateProduct(_ context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	request := model.CreateProductRequest{
		Id:       uuid.New().String(),
		Name:     req.Product.GetName(),
		Price:    req.Product.GetPrice(),
		Quantity: req.Product.GetQuantity(),
	}

	res, err := ProductService.Create(request)
	if err != nil {
		return &proto.CreateProductResponse{}, err
	}

	fmt.Printf("create product : %v \n", res)
	return &proto.CreateProductResponse{
		Product: &proto.Product{
			Id:       res.Id,
			Name:     res.Name,
			Price:    res.Price,
			Quantity: res.Quantity,
		},
	}, nil
}

func (s *server) GetProduct(_ context.Context, req *proto.GetProductRequest) (*proto.GetProductResponse, error) {
	res, err := ProductService.Get(req.ProductId)
	if err != nil {
		return &proto.GetProductResponse{}, err
	}

	fmt.Printf("get product : %v \n", res)

	return &proto.GetProductResponse{
		Product: &proto.Product{
			Id:       res.Id,
			Name:     res.Name,
			Price:    res.Price,
			Quantity: res.Quantity,
		},
	}, nil
}

func (s *server) DeleteProduct(_ context.Context, req *proto.DeleteProductRequest) (*proto.DeleteProductResponse, error) {
	res, err := ProductService.Delete(req.ProductId)
	if err != nil {
		return &proto.DeleteProductResponse{}, err
	}

	fmt.Printf("delete product : %v \n", res)

	return &proto.DeleteProductResponse{
		ProductId: res,
	}, nil
}

func (s *server) ListProduct(_ context.Context, _ *proto.ListProductRequest) (*proto.ListProductResponse, error) {
	res, err := ProductService.List()
	if err != nil {
		return &proto.ListProductResponse{}, err
	}

	var products []*proto.Product
	for _, product := range res {
		products = append(products, &proto.Product{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}

	fmt.Printf("get all product : %v \n", products)
	return &proto.ListProductResponse{
		Product: products,
	}, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db := config.NewDBConn()

	productRepo := repository.NewProductRepository(db)
	ProductService = service.NewProductService(&productRepo)

	fmt.Println("Product service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: #{err}")
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	proto.RegisterProductServiceServer(s, &server{})

	reflection.Register(s)

	go func() {
		fmt.Println("Starting server ...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: #{err}")
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	fmt.Println("Stopping the server")
	s.Stop()
}
