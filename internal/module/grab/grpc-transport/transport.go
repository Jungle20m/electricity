package grpc_transport

import (
	"github.com/Jungle20m/electricity/component"
	pb "github.com/Jungle20m/electricity/internal/module/grab/grpc-transport/protoc"
	"google.golang.org/grpc"
)

type Transportation struct {
	pb.UnimplementedOrderServer
	AppContext component.AppContext
}

func NewTransportation(appCtx component.AppContext) *Transportation {
	return &Transportation{AppContext: appCtx}
}

func Register(server *grpc.Server, appCtx component.AppContext) error {
	transport := NewTransportation(appCtx)
	pb.RegisterOrderServer(server, transport)
	return nil
}
