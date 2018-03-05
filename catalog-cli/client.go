package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/jinpeng/gcommerce/catalog-service/proto/catalog"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:50051"
	defaultFilename = "catalog.json"
)

func parseFile(file string) (*pb.Product, error) {
	var product *pb.Product
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &product)
	return product, err
}

func main() {
	cmd.Init()

	client := pb.NewCatalogServiceClient("go.micro.srv.catalog", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	product, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateProduct(context.Background(), product)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
	log.Printf("Created: %v", r.Product)

	getAll, err := client.GetProducts(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list products: %v", err)
	}
	log.Print("GetProducts:")
	for _, v := range getAll.Products {
		log.Println(v)
	}
}
