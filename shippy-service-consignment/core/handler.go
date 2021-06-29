package core

import (
	"context"
	"errors"
	"log"

	// Import the generated protobuf code
	pb "github.com/AlexanderKorovayev/microservice/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if vesselResponse == nil {
		return nil, errors.New("error fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	//consignment, err := s.repo.Create(req) возможно это рабочий варинат
	consignment := s.repository.Create(ctx, MarshalConsignment(req))

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err, nil
	}
	return &pb.Response{Consignments: UnmarshalConsignmentCollection(consignments)}, nil
}
