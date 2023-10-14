package handlergrpc_test

import (
	"context"
	handlergrpc "library-exercise/handler/grpc"
	"library-exercise/mocks"
	"library-exercise/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestBookHandler_GetAllBooks(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return all books", func(t *testing.T) {
		// Create a mock service
		uu := mocks.UserUsecase{}
		uh := handlergrpc.NewUserHandler(&uu)

		// Create a gRPC server with the mock service
		server := grpc.NewServer()
		pb.RegisterAuthServer(server, uh)

		// Create a gRPC client for testing
		clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			t.Fatalf("Failed to create client connection: %v", err)
		}
		defer clientConn.Close()

		grpcClient := pb.NewAuthClient(clientConn)

		// Call the gRPC method
		response, err := grpcClient.Login(context.Background(), &pb.LoginRequest{
			Email:    "satria@mail.com",
			Password: "password",
		})

		// Assert the results
		assert.Equal(t, pb.LoginResponse{Message: "success", Token: "initoken"}, response)
		assert.Equal(t, nil, err)
	})
}
