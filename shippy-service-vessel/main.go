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
	port = ":50052"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	GetAll() []*pb.Vessel
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// GetAll vessels
func (repo *VesselRepository) GetAll() []*pb.Vessel {
	return repo.vessels
}

// Our grpc service handler
type service struct {
	repo Repository
	pb.UnimplementedVesselServiceServer
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification) (*pb.Response, error) {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return nil, err
	}
	vessels := s.repo.GetAll()
	// Set the vessel as part of the response message type
	return &pb.Response{Vessel: vessel, Vessels: vessels}, nil
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
	pb.RegisterVesselServiceServer(s, &service{repo, pb.UnimplementedVesselServiceServer{}})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
