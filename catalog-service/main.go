package main

import (
	"log"

	// Import the generated protobuf code
	pb "github.com/jinpeng/gcommerce/catalog-service/proto/catalog"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// CatalogRepository interface for the service
type CatalogRepository interface {
	Create(*pb.Product) (*pb.Product, error)
	GetAll() []*pb.Product
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	products []*pb.Product
}

// Create create new product and add it to the list
func (repo *Repository) Create(product *pb.Product) (*pb.Product, error) {
	updated := append(repo.products, product)
	repo.products = updated
	return product, nil
}

// GetAll get all products
func (repo *Repository) GetAll() []*pb.Product {
	return repo.products
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo CatalogRepository
}

// Createproduct - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateProduct(ctx context.Context, req *pb.Product, res *pb.Response) error {

	// Save our product
	product, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Product = product
	return nil
}

func (s *service) GetProducts(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	products := s.repo.GetAll()
	res.Products = products
	return nil
}

func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.catalog"),
		micro.Version("latest"),
	)
	srv.Init()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterCatalogServiceHandler(srv.Server(), &service{repo})

	// Register reflection service on gRPC server.
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
