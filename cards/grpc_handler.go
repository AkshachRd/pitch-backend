package main

import (
	"context"
	"log"

	pb "github.com/AkshachRd/pitch-backend/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedCardServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) pb.CardServiceServer {
	handler := &grpcHandler{}
	pb.RegisterCardServiceServer(grpcServer, handler)
	return handler
}

func (h *grpcHandler) CreateCard(ctx context.Context, p *pb.CreateCardRequest) (*pb.Card, error) {
	log.Printf("New card created %v", p)
	card := pb.Card{Id: "1", FrontSide: "New card", BackSide: "Новая карточка"}
	return &card, nil
}
