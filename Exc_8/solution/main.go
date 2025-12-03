package main

import (
	"exc8/client"
	"exc8/server"
	"log"
	"time"
)

func main() {
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	time.Sleep(1 * time.Second)
	// todo start client
	grpcClient, err := client.NewGrpcClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Then, run the logic
	if err := grpcClient.Run(); err != nil {
		log.Fatalf("Client error: %v", err)
	}
	println("Orders complete!")
}
