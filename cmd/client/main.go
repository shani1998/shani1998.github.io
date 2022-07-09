package main

import (
	"context"
	"log"

	"github.com/shani1998/shani1998.github.io/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Addr = "127.0.0.1:50051"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	cl := pb.NewPortfolioClient(conn)

	skills, err := cl.ListSkills(context.Background(), &pb.SkillRequest{})
	if err != nil {
		log.Fatalf("failed to list skills: %v", err)
	}

	log.Printf("skills: %v", skills)
}
