package usecase

import (
	"godb/model"
	"godb/repository"
)

type StoreUseCase interface {
	GetAllStore() ([]model.Store, error)
	GetAllStoreWithProducts() ([]model.StoreProducts, error)
}

type storeUseCase struct {
	storeRepo repository.StoreRepository
	productUc ProductUseCase
}

func (st *storeUseCase) GetAllStore() ([]model.Store, error) {
	return st.storeRepo.GetAll()
}

func (st *storeUseCase) GetAllStoreWithProducts() ([]model.StoreProducts, error) {
	var storeProducts []model.StoreProducts
	stores, _ := st.GetAllStore()
	for _, store := range stores {
		var storeProduct model.StoreProducts
		products, _ := st.productUc.GetProductByStoreId(store.Id)
		storeProduct.Products = products
		storeProducts = append(storeProducts, storeProduct)
	}
	return storeProducts, nil
}

func NewStoreUseCase(storeRepository repository.StoreRepository, productUseCase ProductUseCase) StoreUseCase {
	return &storeUseCase{
		storeRepo: storeRepository,
		productUc: productUseCase,
	}
}
