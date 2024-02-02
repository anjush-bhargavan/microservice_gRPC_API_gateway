package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	dom "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/DOM"
	userpb "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

// UserLoginHandler bind the data and pass it with response
func UserLoginHandler(c *gin.Context, client userpb.UserServiceClient, role string) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var user dom.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.UserLogin(ctxt, &userpb.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
		Role:     role,
	})
	if err != nil {
		log.Printf("error logging in user %v err: %v", user.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in successfully", user.Email),
		"data":    response,
	})
}

// UserSignupHandler bind the data and pass it with response
func UserSignupHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var user dom.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.UserSignup(ctxt, &userpb.SignupRequest{
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Printf("error signing up user %v err: %v", user.UserName, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v created successfully", user.UserName),
		"data":    response,
	})
}
