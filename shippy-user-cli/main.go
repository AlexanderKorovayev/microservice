package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/AlexanderKorovayev/microservice/shippy-service-user/proto/user"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50053" //"host.docker.internal:50053" 127.0.0.1:50053 user:50053
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewUserServiceClient(conn)

	ctx := context.Background()
	user := &proto.User{
		Name:     "name",
		Email:    "email",
		Company:  "company",
		Password: "password",
	}

	rsp, err := client.Create(ctx, user)

	if err != nil {
		log.Println(err)
	}

	// print the response
	fmt.Println("Response: ", rsp.User)
}
