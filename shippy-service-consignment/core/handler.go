package core

import (
	"context"
	"errors"
	"log"

	// Import the generated protobuf code
	pb "github.com/AlexanderKorovayev/microservice/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/proto/vessel"
)

type Handler struct {
	Repository
	VesselClient vesselProto.VesselServiceClient
	pb.UnimplementedShippingServiceServer
}

func (s *Handler) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.VesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if vesselResponse == nil {
		return nil, errors.New("error fetching vessel, returned nil")
	}

	if err != nil {
		return nil, err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	//consignment, err := s.repo.Create(req) возможно это рабочий варинат
	err = s.Repository.Create(ctx, MarshalConsignment(req))

	if err != nil {
		return nil, err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true}, nil
}

func (s *Handler) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments, err := s.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Consignments: UnmarshalConsignmentCollection(consignments)}, nil
}
