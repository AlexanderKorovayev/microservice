package main

import (
	"context"
	"log"
	"os"

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
	/*
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
		fmt.Println("created", rsp.User)
	*/
	getAll, err := client.GetAll(context.Background(), &proto.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	/*
		authResponse, err := client.Auth(context.TODO(), &proto.User{
			Email:    "email",
			Password: "password",
		})

		if err != nil {
			log.Fatalf("Could not authenticate user: %s error: %v\n", "email", err)
		}

		log.Printf("Your access token is: %s \n", authResponse.Token)
	*/

	// secret token is
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoiZW1haWwiLCJwYXNzd29yZCI6InBhc3N3b3JkIn0sImV4cCI6MTUwMDAsImlzcyI6Im1pY3Jvc2VydmljZS5zZXJ2aWNlLnVzZXIifQ.VjwVUwg687y-ztrpw7fiuvFvo1h_4nn2bK3hep7cx0A

	//https://pkg.go.dev/github.com/dgrijalva/jwt-go#StandardClaims.VerifyIssuer
	//https://stackoverflow.com/questions/61281636/how-to-access-jwt-sub-claims-using-go
	//надо разобраться почему не происходит раскодирование токена
	validateResponse, err := client.ValidateToken(context.Background(),
		&proto.Token{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoiZW1haWwiLCJwYXNzd29yZCI6InBhc3N3b3JkIn0sImV4cCI6MTUwMDB9.E7qfTlNFGT6lfNFJAX6pOtHtbTzSyaSGFWc5TJbod6E"})

	if err != nil {
		log.Println(err)
	}
	log.Println(validateResponse.Valid)

	// let's just exit because
	os.Exit(0)
}
