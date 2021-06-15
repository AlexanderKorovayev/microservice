module github.com/AlexanderKorovayev/microservice/shippy-cli-consignment

go 1.16

replace github.com/AlexanderKorovayev/microservice/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/AlexanderKorovayev/microservice/shippy-service-consignment v0.0.0-20210610150148-174729434b99
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
	golang.org/x/net v0.0.0-20210610132358-84b48f89b13b // indirect
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	google.golang.org/genproto v0.0.0-20210611144927-798beca9d670 // indirect
	google.golang.org/grpc v1.38.0
)
