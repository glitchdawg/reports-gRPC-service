package server

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	pb "github.com/glitchdawg/reportservice/proto" 
)


type ReportStore struct {
	mu      sync.Mutex
	Reports map[string]string 
}

type ReportServiceServer struct {
	pb.UnimplementedReportServiceServer
	Store *ReportStore
}

func (s *ReportServiceServer) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
	reportID := uuid.New().String()

	
	s.Store.mu.Lock()
	s.Store.Reports[reportID] = req.UserId
	s.Store.mu.Unlock()

	log.Printf("[%s] Generated report for user %s (report ID: %s)", time.Now().Format(time.RFC3339), req.UserId, reportID)

	return &pb.GenerateReportResponse{ReportId: reportID}, nil
}

func (s *ReportServiceServer) HealthCheck(ctx context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("[%s] HealthCheck called", time.Now().Format(time.RFC3339))
	return &pb.HealthCheckResponse{Status: "SERVING"}, nil
}
