package admin

import (
	"log"

	adminpb "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (adminpb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: 8083 ")
		return nil, err
	}
	log.Printf("succefully connected to admin server at port :8083")
	return adminpb.NewAdminServiceClient(grpc),nil
}