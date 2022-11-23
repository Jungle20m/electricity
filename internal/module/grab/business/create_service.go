package business

import (
	"context"
	"github.com/Jungle20m/electricity/component"
	grabModel "github.com/Jungle20m/electricity/internal/module/grab/model"
)

type CreateServiceStorageInterface interface {
	WithTx(ctx context.Context, fn func(c context.Context) error) error
	InsertService(ctx context.Context, service grabModel.ServiceCreate) error
}

type createServiceBusiness struct {
	appCtx  component.AppContext
	storage CreateServiceStorageInterface
}

func NewCreateServiceBusiness(appCtx component.AppContext, storage CreateServiceStorageInterface) *createServiceBusiness {
	return &createServiceBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (business *createServiceBusiness) CreateNewService(ctx context.Context, serviceCreate grabModel.ServiceCreate) error {
	err := business.storage.WithTx(ctx, func(txContext context.Context) error {
		serviceCreate.Active = 1
		err := business.storage.InsertService(txContext, serviceCreate)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
