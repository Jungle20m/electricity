package business

import (
	"context"
	"fmt"
	"github.com/Jungle20m/electricity/component"
)

type GetOrderStorageInterface interface {
}

type getOrderBusiness struct {
	appCtx  component.AppContext
	storage GetOrderStorageInterface
}

func NewGetOrderBusiness(appCtx component.AppContext, storage GetOrderStorageInterface) *getOrderBusiness {
	return &getOrderBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (business *getOrderBusiness) GetCustomerOrders(ctx context.Context) {
	fmt.Println("get order business")
}
