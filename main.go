package main

import (
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"github.com/glitchdawg/reportservice/internal/server"
	pb "github.com/glitchdawg/reportservice/proto"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

var predefinedUsers = []string{"user-101", "user-202", "user-303"}

func main() {
	store := &server.ReportStore{
		Reports: make(map[string]string),
	}

	s := grpc.NewServer()
	reportServer := &server.ReportServiceServer{Store: store}
	pb.RegisterReportServiceServer(s, reportServer)
	reflection.Register(s)

	go startCronJob(reportServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startCronJob(s *server.ReportServiceServer) {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		for _, uid := range predefinedUsers {
			_, err := s.GenerateReport(nil, &pb.GenerateReportRequest{UserId: uid})
			if err != nil {
				log.Printf("Cron failed for user %s: %v", uid, err)
			}
		}
	})
	c.Start()
}
