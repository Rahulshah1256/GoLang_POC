package main

import (
	"context"
	"log"
	"net"

	pb "CRUD_gRPC/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type userService struct{}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Implement user creation logic using PostgreSQL or any other database
	// Return the created user ID in the CreateUserResponse
	return &pb.CreateUserResponse{
		Id: 123,
	}, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Implement logic to retrieve user from the database based on the provided ID
	// Return the user's name and age in the GetUserResponse
	return &pb.GetUserResponse{
		Name: "John Doe",
		Age:  30,
	}, nil
}

func main() {
	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, &userService{})
		log.Println("gRPC server started on port 50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start HTTP server using Gin framework
	router := gin.Default()

	router.GET("/users/:id", func(c *gin.Context) {
		// Extract the ID from the request parameters
		//id := c.Param("id")

		// Call the gRPC GetUser endpoint to retrieve user details
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()

		client := pb.NewUserServiceClient(conn)
		resp, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: 123})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"name": resp.Name,
			"age":  resp.Age,
		})
	})

	log.Println("HTTP server started on port 8080")
	router.Run(":8080")
}
