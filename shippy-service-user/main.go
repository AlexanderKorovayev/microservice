package main

import (
	"log"
	"net"

	core "github.com/AlexanderKorovayev/microservice/shippy-service-user/core"
	pb "github.com/AlexanderKorovayev/microservice/shippy-service-user/proto/user"
	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

func main() {

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := core.CreateConnection()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := core.NewPostgresRepository(db)
	tokenService := core.TokenService{repo}

	pb.RegisterUserServiceServer(s, &core.Handler{tokenService, pb.UnimplementedUserServiceServer{}})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
