package business

import (
	"context"
	"github.com/Jungle20m/electricity/component"
	grabModel "github.com/Jungle20m/electricity/internal/module/grab/model"
	grabStorage "github.com/Jungle20m/electricity/internal/module/grab/storage"
)

//type CreateServiceStorageInterface interface {
//	InsertService(ctx context.Context, service grabModel.ServiceCreate) error
//	WithTransaction(ctx context.Context, fn func(storage CreateServiceStorageInterface) error) error
//}

type createServiceBusiness struct {
	appCtx  component.AppContext
	storage *grabStorage.Storage
}

func NewCreateServiceBusiness(appCtx component.AppContext, storage *grabStorage.Storage) *createServiceBusiness {
	return &createServiceBusiness{
		appCtx:  appCtx,
		storage: storage,
	}
}

func (business *createServiceBusiness) CreateNewService(ctx context.Context, serviceCreate grabModel.ServiceCreate) error {
	err := business.storage.WithTx(ctx, func(storage *grabStorage.Storage) error {
		serviceCreate.Active = 1
		err := storage.InsertService(ctx, serviceCreate)
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
