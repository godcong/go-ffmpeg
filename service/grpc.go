package service

import (
	"context"
	"fmt"
	"github.com/godcong/node-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// RestServer ...
type GRPCServer struct {
	Port   string
	server *grpc.Server
}

// RemoteDownload ...
func (s *GRPCServer) RemoteDownload(ctx context.Context, p *proto.RemoteDownloadRequest) (*proto.ServiceReply, error) {
	log.Printf("Received: %v", p.String())
	return Result(nil), nil
}

// Status ...
func (s *GRPCServer) Status(ctx context.Context, p *proto.StatusRequest) (*proto.ServiceReply, error) {
	log.Printf("Received: %v", p.String())
	return Result(nil), nil
}

// Result ...
func Result(detail *proto.ReplyDetail) *proto.ServiceReply {
	return &proto.ServiceReply{
		Code:                 0,
		Message:              "",
		Detail:               detail,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}

func NewGRPCServer() *GRPCServer {
	port := config.GRPC.Port
	if port == "" {
		port = ":7781"
	}
	return &GRPCServer{
		Port: port,
	}
}

// Start ...
func (s *GRPCServer) Start() {
	if !config.GRPC.Enable {
		return
	}
	log.Println("starting grpc")
	s.server = grpc.NewServer()

	go func() {
		lis, err := net.Listen("tcp", s.Port)
		if err != nil {
			panic(fmt.Sprintf("failed to listen: %v", err))
		}

		proto.RegisterNodeServiceServer(s.server, s)
		// Register reflection service on gRPC server.
		reflection.Register(s.server)
		log.Printf("Listening and serving TCP on %s\n", s.Port)
		if err := s.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

// Stop ...
func (s *GRPCServer) Stop() {
	s.server.Stop()
}
