package main

import (
	"context"
	"log"
	"net"

	// Import the generated protobuf code
	core "github.com/AlexanderKorovayev/microservice/shippy-service-consignment/core"
	pb "github.com/AlexanderKorovayev/microservice/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/proto/vessel"
	"google.golang.org/grpc"
)

const (
	defaultHost   = "mongodb://datastore:27017" //mongodb://127.0.0.1:27017 mongodb://datastore:27017
	port          = ":50051"
	vesselAddress = "vessel:50052" //"host.docker.internal:50052" 127.0.0.1:50052 vessel:50052
)

func main() {

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Set up a connection to the vessel server.
	conn, err := grpc.Dial(vesselAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client, err := core.CreateClient(context.Background(), defaultHost, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &core.MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselServiceClient(conn)

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceServer(s, &core.Handler{repository, vesselClient, pb.UnimplementedShippingServiceServer{}})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
