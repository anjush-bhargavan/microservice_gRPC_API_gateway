package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	dom "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/DOM"
	adminpb "github.com/anjush-bhargavan/microservice_gRPC_API_gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func AdminLoginHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var admin dom.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx := context.Background()

	response, err := client.AdminLogin(ctx, &adminpb.AdminRequest{
		Username: admin.Username,
		Password: admin.Password,
	})
	if err != nil {
		log.Printf("error logging in admin %v err: %v", admin.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", admin.Username),
		"data":    response,
	})
}

func CreateUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var user dom.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding .. user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}

	ctx := context.Background()

	response, err := client.CreateUserFn(ctx, &adminpb.User{
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     "user",
	})
	if err != nil {
		log.Printf("error creating user %v err: %v", user.UserName, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v user created succesfully", user.UserName),
		"data":    response,
	})
}

func EditUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var user dom.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}

	ctx := context.Background()

	response, err := client.EditUserFn(ctx, &adminpb.User{
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     "user",
	})
	if err != nil {
		log.Printf("error editing user %v err: %v", user.UserName, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "binding error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v user edited succesfully", user.UserName),
		"data":    response,
	})
}

func FindAllUsersHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctx := context.Background()

	response, err := client.FindAllUsersFn(ctx, &adminpb.NoParam{})
	if err != nil {
		log.Printf("error fetching users err: %v", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "error fetching users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": ("users fetched succesfully"),
		"data":    response,
	})
}

func DeleteUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	response, err := client.DeleteUserFn(ctx, &adminpb.UserID{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error deleting user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "user deleted successfully",
		"data":    response,
	})
}

func FindUserByEmailHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	email := c.Query("email")

	ctx := context.Background()
	response, err := client.FindUserByEmailFn(ctx, &adminpb.UserRequest{
		Email: email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error finding user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "user found",
		"data":    response,
	})
}

func FindUserByIDHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	response, err := client.FindUserByIDFn(ctx, &adminpb.UserID{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error finding user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "user found",
		"data":    response,
	})
}
