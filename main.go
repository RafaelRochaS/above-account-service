package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RafaelRochaS/above-account-service/handlers"
	"github.com/RafaelRochaS/above-account-service/repository"
	"github.com/RafaelRochaS/above-account-service/utils"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting Account Service...")
	utils.LoadEnvs()
	startServer()
}

func startServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", utils.HOST, utils.PORT))
	if err != nil {
		log.Fatalf("ERROR AccountService :: %v", err)
	}

	server := grpc.NewServer()
	dbConn := &repository.DbConn{}
	if err = dbConn.InstantiateDbConn(); err != nil {
		log.Fatalf("AccountService :: failed to instantiate DB: %v", err)
	}

	handlers.RegisterAccountService(server, dbConn)
	log.Printf("AccountService :: listening at %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("AccountService :: failed to start serving: %v", err)
	}
}
