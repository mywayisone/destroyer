package main

import (
	"log"

	"github.com/mywayisone/gRPC-project/destroyer/api/server"
)

func main() {
	address := ":50051"
	if err := server.StartServer(address); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}