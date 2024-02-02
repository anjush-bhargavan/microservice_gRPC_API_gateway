package main

import (
	"github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/admin"
	"github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/server"
	"github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/user"
)


func main () {
	server := server.Server()
	user.NewUserRoute(server.R)
	admin.NewAdminRoute(server.R)
	server.StartServer()
}