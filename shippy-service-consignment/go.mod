module github.com/AlexanderKorovayev/microservice/shippy-service-consignment

go 1.15

//replace github.com/AlexanderKorovayev/microservice/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/AlexanderKorovayev/microservice/shippy-service-vessel v0.0.0-20210629062719-65a08fd11324
	github.com/aws/aws-sdk-go v1.38.69 // indirect
	github.com/klauspost/compress v1.13.1 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.5.3
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	google.golang.org/genproto v0.0.0-20210629135825-364e77e5a69d // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
)
