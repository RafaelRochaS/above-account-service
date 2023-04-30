package handlers

import (
	"context"
	"log"

	pb "github.com/RafaelRochaS/above-account-service/above_account_service"
)

func (s *Server) Healthcheck(ctx context.Context, input *pb.HealthcheckRequest) (*pb.HealthcheckReply, error) {
	log.Printf("HealthcheckHandler :: new request")

	if err := s.dbConn.CheckConn(); err != nil {
		return &pb.HealthcheckReply{
			Status: "DB connection failed",
		}, err
	}

	return &pb.HealthcheckReply{
		Status: "ok",
	}, nil
}
