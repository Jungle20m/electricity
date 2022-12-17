package grpcserver

import (
	"fmt"
	"github.com/Jungle20m/electricity/component"
	grabTransport "github.com/Jungle20m/electricity/internal/module/grab/grpc-transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	AppContext component.AppContext
}

func NewServer(appCtx component.AppContext) *Server {
	return &Server{AppContext: appCtx}
}

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", "8001"))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	// router
	grabTransport.Register(grpcServer, s.AppContext)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
