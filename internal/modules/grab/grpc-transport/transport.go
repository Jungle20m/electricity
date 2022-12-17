package grpc_transport

import (
	"github.com/Jungle20m/electricity/common"
	pb "github.com/Jungle20m/electricity/internal/modules/grab/grpc-transport/proto"
	"google.golang.org/grpc"
)

type Transportation struct {
	pb.UnimplementedOrderServer
	pb.UnimplementedServiceServer
	AppContext common.AppContext
}

func newTransportation(appCtx common.AppContext) *Transportation {
	return &Transportation{AppContext: appCtx}
}

func Register(server *grpc.Server, appCtx common.AppContext) error {
	transport := newTransportation(appCtx)
	pb.RegisterOrderServer(server, transport)
	pb.RegisterServiceServer(server, transport)
	return nil
}
