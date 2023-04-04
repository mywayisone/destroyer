package server

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mywayisone/gRPC-project/destroyer/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server represents the gRPC server
type Server struct{
	 pb.UnimplementedDestroyerServer
}

// AcquireTargets is a gRPC method that acquires targets
func (s *Server) AcquireTargets(ctx context.Context, req *pb.TargetsRequest) (*empty.Empty, error) {
	for _, target := range req.Targets {
		log.Printf("Acquiring target: ID=%s, Message=%s, CreatedOn=%v", target.Id, target.Message, target.CreatedOn.AsTime())
		// Add code to acquire targets here
	}
	return &empty.Empty{}, nil
}

// ListTargets is a gRPC method that lists targets acquired from the database
func (s *Server) ListTargets(ctx context.Context, req *empty.Empty) (*pb.TargetsResponse, error) {
	log.Println("Listing targets")
	// Add code to list targets from the database here
	return &pb.TargetsResponse{}, nil
}

// StartServer starts the gRPC server
func StartServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterDestroyerServer(srv, &Server{})
	log.Println("Starting server...")
	return srv.Serve(lis)
}

// Helper function to convert a time.Time to a google.protobuf.Timestamp
func toTimestampProto(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
