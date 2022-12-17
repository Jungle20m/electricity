package business

import (
	"context"
	"github.com/Jungle20m/electricity/common"
	productModel "github.com/Jungle20m/electricity/internal/modules/product/model"
)

type GetProductStorageInterface interface {
	GetProductByID(ctx context.Context, id int) (*productModel.Product, error)
}

type getProductBusiness struct {
	storage GetProductStorageInterface
	appCtx  common.AppContext
}

func NewGetProductBusiness(storage GetProductStorageInterface, appCtx common.AppContext) *getProductBusiness {
	return &getProductBusiness{
		storage: storage,
		appCtx:  appCtx,
	}
}

func (business *getProductBusiness) GetProductByID(ctx context.Context, id int) (*productModel.Product, error) {
	return business.storage.GetProductByID(ctx, id)
}
