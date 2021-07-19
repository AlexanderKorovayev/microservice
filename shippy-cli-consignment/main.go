// shippy/shippy-cli-consignment/main.go
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/AlexanderKorovayev/microservice/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "consignment:50051" //"host.docker.internal:50051" 127.0.0.1:50051 consignment:50051
	defaultFilename = "consignment.json"
	token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoiZW1haWwiLCJwYXNzd29yZCI6InBhc3N3b3JkIn0sImV4cCI6MTUwMDAsImlzcyI6Im1pY3Jvc2VydmljZS5zZXJ2aWNlLnVzZXIifQ.VjwVUwg687y-ztrpw7fiuvFvo1h_4nn2bK3hep7cx0A"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
