package business

import (
	"context"
	"fmt"
	"github.com/Jungle20m/electricity/component"
	grabModel "github.com/Jungle20m/electricity/internal/module/grab/model"
	grabStorage "github.com/Jungle20m/electricity/internal/module/grab/storage"
)

type CreateServiceStorageInterface interface {
	InsertService(ctx context.Context, service grabModel.ServiceCreate) error
	WithTransaction(ctx context.Context, fn func(storage *grabStorage.Storage) error) error
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
	err := business.storage.WithTransaction(ctx, func(storage CreateServiceStorageInterface) error {
		serviceCreate.Active = 1

		storage.InsertService(ctx, serviceCreate)

		return nil
	})

	if err != nil {
		fmt.Printf("error to create service: %v", err)
	}

	//err := business.storage.InsertService(ctx, serviceCreate)
	//if err != nil {
	//	fmt.Printf("error to create service: %v", err)
	//}

	return nil
}
