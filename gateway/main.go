package main

import (
	"log"
	"net/http"

	"github.com/AkshachRd/pitch-backend/common"
	pb "github.com/AkshachRd/pitch-backend/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
	cardServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.NewClient(cardServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to gRPC server: ", err)
	}
	defer conn.Close()

	log.Println("Listening cards on " + httpAddr)

	c := pb.NewCardServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Println("Starting server on " + httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}