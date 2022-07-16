package main

import (
	"context"
	"html/template"
	"net/http"

	"github.com/shani1998/shani1998.github.io/pkg/pb"
	log "github.com/sirupsen/logrus"
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

	// index route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		skills, err := cl.ListSkills(context.Background(), &pb.SkillRequest{})
		if err != nil {
			log.Fatalf("failed to list skills: %v", err)
		}

		log.Printf("skills: %v", skills)

		t, err := template.ParseFiles(staticFilesDir + "/index.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := t.Execute(w, skills); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Errorf("failed to start portfolio svc, Reason %v", err)
	}

}
