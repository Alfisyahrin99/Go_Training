package usecase

import (
	"godb/model"
	"godb/repository"
)

type StoreUseCase interface {
	GetAllStore() ([]model.StoreProducts, error)
}

type storeUseCase struct {
	storeRepo repository.StoreRepository
}

func (st *storeUseCase) GetAllStore() ([]model.StoreProducts, error) {
	return st.storeRepo.GetAll()
}

func NewStoreUseCase(storeRepository repository.StoreRepository) StoreUseCase {
	return &storeUseCase{
		storeRepo: storeRepository,
	}
}
