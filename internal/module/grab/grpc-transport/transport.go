package grpc_transport

import (
	"github.com/Jungle20m/electricity/component"
	pb "github.com/Jungle20m/electricity/internal/module/grab/grpc-transport/proto"
	"google.golang.org/grpc"
)

type Transportation struct {
	pb.UnimplementedOrderServer
	pb.UnimplementedServiceServer
	AppContext component.AppContext
}

func newTransportation(appCtx component.AppContext) *Transportation {
	return &Transportation{AppContext: appCtx}
}

func Register(server *grpc.Server, appCtx component.AppContext) error {
	transport := newTransportation(appCtx)
	pb.RegisterOrderServer(server, transport)
	pb.RegisterServiceServer(server, transport)
	return nil
}
