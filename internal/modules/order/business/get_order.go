package business

import (
	"context"
	"github.com/Jungle20m/electricity/common"
	orderModel "github.com/Jungle20m/electricity/internal/modules/order/model"
)

type GetOrderStorageInterface interface {
	GetOrderByID(ctx context.Context, id int) (*orderModel.Order, error)
}

type getOrderBusiness struct {
	storage GetOrderStorageInterface
	appCtx  common.AppContext
}

func NewGetOrderBusiness(storage GetOrderStorageInterface, appCtx common.AppContext) *getOrderBusiness {
	return &getOrderBusiness{
		storage: storage,
		appCtx:  appCtx,
	}
}

func (business *getOrderBusiness) GetOrderByID(ctx context.Context, id int) (*orderModel.Order, error) {
	// get name
	return business.storage.GetOrderByID(ctx, id)
}
