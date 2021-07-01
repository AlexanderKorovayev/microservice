package main

import (
	"context"
	"log"
	"net"

	// Import the generated protobuf code
	core "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/core"
	pb "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/proto/vessel"
	"google.golang.org/grpc"
)

const (
	defaultHost = "mongodb://127.0.0.1:27017"
	port        = ":50052"
)

func main() {

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// create vessel collections
	client, err := core.CreateClient(context.Background(), defaultHost, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("vessel").Collection("vessels")

	repository := &core.MongoRepository{vesselCollection}

	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	for _, v := range vessels {
		repository.Create(context.TODO(), core.MarshalVessel(v))
	}
	/*
		пробный поиск
		type Trainer struct {
			Name string
			Age  int
			City string
		}
		ash := Trainer{"Ash", 10, "Pallet Town"}
		insertResult, err := vesselCollection.InsertOne(context.TODO(), ash)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Inserted a single document: ", insertResult.InsertedID)
		filter := bson.D{{"name", "Ash"}}
		var result Trainer
		err = vesselCollection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Found a single document: %+v\n", result)
	*/

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterVesselServiceServer(s, &core.Handler{repository, pb.UnimplementedVesselServiceServer{}})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
