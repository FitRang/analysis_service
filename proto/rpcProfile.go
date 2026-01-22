package proto

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProfileClient struct {
	client ProfileServiceClient
}

func NewProfileClient() *ProfileClient {
	addr := "profile-service:8081"

	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &ProfileClient{
		client: NewProfileServiceClient(conn),
	}
}
