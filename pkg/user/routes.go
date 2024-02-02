package user

import (
	"log"

	"github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/user/handler"
	pb "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type User struct {
	client pb.UserServiceClient
}

func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c,u.client,"user")
}

func (u *User) Signup(c *gin.Context) {
	handler.UserSignupHandler(c,u.client)
}

func NewUserRoute(c *gin.Engine)  {
	client , err := ClientDial()
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v",err.Error())
	}

	userHandler := User {
		client: client,
	}
	apiUser := c.Group("/api/user")
	{
		apiUser.POST("/login",userHandler.Login)
		apiUser.POST("/signup",userHandler.Signup)
	}
}