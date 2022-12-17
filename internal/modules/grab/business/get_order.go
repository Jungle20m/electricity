package business

import (
	"context"
	"fmt"
	"github.com/Jungle20m/electricity/common"
)

type GetOrderStorageInterface interface {
}

type getOrderBusiness struct {
	appCtx  common.AppContext
	storage GetOrderStorageInterface
}

func NewGetOrderBusiness(appCtx common.AppContext, storage GetOrderStorageInterface) *getOrderBusiness {
	return &getOrderBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (business *getOrderBusiness) GetCustomerOrders(ctx context.Context) {
	fmt.Println("get order business")
}
