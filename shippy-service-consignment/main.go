package main

import (
	"context"
	"errors"
	"log"
	"net"

	// Import the generated protobuf code
	pb "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/proto/vessel"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type VesselRepository struct {
	vessels []*pb.Vesselt
}

// Create a new consignment
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type vesselService struct {
	repo repository
	pb.UnimplementedShippingServiceServer
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification) (*pb.Response, error) {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return nil, err
	}

	// Set the vessel as part of the response message type
	return &pb.Response{Created: true, Vessel: vessel}, nil
}

func main() {

	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceServer(s, &vesselService{repo, pb.UnimplementedShippingServiceServer{}})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
