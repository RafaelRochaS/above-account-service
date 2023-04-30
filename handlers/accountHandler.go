package handlers

import (
	"context"
	"log"

	pb "github.com/RafaelRochaS/above-account-service/above_account_service"
	"github.com/RafaelRochaS/above-account-service/models"
	"github.com/RafaelRochaS/above-account-service/repository"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServer
	dbConn repository.IDBConn
}

func RegisterAccountService(grpcServer *grpc.Server, dbConn repository.IDBConn) {
	server := &Server{
		dbConn: dbConn,
	}
	pb.RegisterUserServer(grpcServer, server)
}

func (s *Server) CreateUser(ctx context.Context, input *pb.UserCreateRequest) (*pb.UserCreateReply, error) {
	log.Printf("AccountHandler :: CreateUser request: %v", input)

	err := s.dbConn.Create(&models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Address:   input.Address,
		Age:       int(input.Age),
	})

	if err != nil {
		return &pb.UserCreateReply{
			Status: "failed",
		}, err
	}

	return &pb.UserCreateReply{
		Status: "success",
	}, nil
}
