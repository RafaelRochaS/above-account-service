package handlers

import (
	"context"
	"log"

	pb "github.com/RafaelRochaS/above-account-service/above_account_service"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServer
}

func RegisterAccountService(server *grpc.Server) {
	pb.RegisterUserServer(server, &Server{})
}

func (s *Server) CreateUser(ctx context.Context, input *pb.UserCreateRequest) (*pb.UserCreateReply, error) {
	log.Printf("AccountHandler :: CreateUser request: %v", input)

	return &pb.UserCreateReply{
		Status: "success",
	}, nil
}
