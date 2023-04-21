package handlers

import (
	"context"
	"log"

	pb "github.com/RafaelRochaS/above-account-service/above_account_service"
)

func (s *Server) Healthcheck(ctx context.Context, input *pb.HealthcheckRequest) (*pb.HealthcheckReply, error) {
	log.Printf("HealthcheckHandler :: new request")
	return &pb.HealthcheckReply{
		Status: "ok",
	}, nil
}
