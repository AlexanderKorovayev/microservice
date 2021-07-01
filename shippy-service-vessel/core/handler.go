package core

import (
	"context"

	pb "github.com/AlexanderKorovayev/microservice/shippy-service-vessel/proto/vessel"
)

type Handler struct {
	Repository
	pb.UnimplementedVesselServiceServer
}

func (s *Handler) FindAvailable(ctx context.Context, req *pb.Specification) (*pb.Response, error) {
	// Find the next available vessel
	vessel, err := s.Repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Vessel: UnmarshalVessel(vessel)}, nil
}

func (s *Handler) Create(ctx context.Context, req *pb.Vessel) (*pb.Response, error) {
	if err := s.Repository.Create(ctx, MarshalVessel(req)); err != nil {
		return nil, err
	}
	return &pb.Response{Created: true}, nil
}
