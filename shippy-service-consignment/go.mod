module github.com/AlexanderKorovayev/microservice/shippy-service-consignment

go 1.15

//replace github.com/AlexanderKorovayev/microservice/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/AlexanderKorovayev/microservice/shippy-service-vessel v0.0.0-20210622061245-7198ff2709c9
	go.mongodb.org/mongo-driver v1.5.3 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
