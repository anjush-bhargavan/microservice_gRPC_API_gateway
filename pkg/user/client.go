package user

import (
	"log"

	pb "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (pb.UserServiceClient, error) {
	grpc, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing client on port 8080")
		return nil, err
	}
	log.Printf("Successfully connected to client on port 8080")
	return pb.NewUserServiceClient(grpc), nil
}
