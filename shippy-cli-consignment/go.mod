module github.com/AlexanderKorovayev/microservice/shippy-cli-consignment

go 1.16

//replace github.com/AlexanderKorovayev/microservice/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/AlexanderKorovayev/microservice/shippy-service-consignment v0.0.0-20210616150543-2b85c343e061
	github.com/google/uuid v1.2.0 // indirect
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	github.com/miekg/dns v1.1.42 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	google.golang.org/genproto v0.0.0-20210614182748-5b3b54cad159 // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
