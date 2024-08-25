package main

import (
	"context"
	"fmt"
	"github.com/JubaerHossain/rootx/pkg/core/app"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize the application
	application, err := app.StartApp()
	if err != nil {
		log.Fatalf("❌ Failed to start application: %v", err)
	}

	fmt.Println(application)

	// Initialize HTTP server
	// httpServer := InitHTTPServer(application)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	// inventoryRepo := postgres.NewPostgresInventoryRepository(application.DB)
	// inventoryService := service.NewInventoryService(inventoryRepo)
	// inventoryServer := grpc.NewInventoryServer(inventoryService)
	// pb.RegisterInventoryServiceServer(grpcServer, inventoryServer)

	grpcPort := ":50053"

	go func() {
		if err := StartGRPCServer(grpcServer, grpcPort); err != nil {
			log.Printf("❌ %v", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(grpcServer, 5*time.Second)
}

func StartGRPCServer(grpcServer *grpc.Server, port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Failed to listen on port %s: %v", port, err)
	}
	log.Printf("🛰️ gRPC server running on port %s", port)
	return grpcServer.Serve(listener)
}

func gracefulShutdown(server *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("⚙️ Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Could not gracefully shutdown the server: %v", err)
	}

	log.Printf("✅ Server gracefully stopped")
}
