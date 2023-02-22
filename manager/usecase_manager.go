package manager

import "godb/usecase"

type UseCaseManager interface {
	ProductUseCase() usecase.ProductUseCase
	StoreUseCase() usecase.StoreUseCase
	Customer() usecase.CustomerUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}


func (u *useCaseManager) ProductUseCase() usercase.ProductUseCase {
	return usecase.NewStoreUseCase(
		u.repoManager.StoreRepository(),
		u.ProductUseCase(),
	)
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomeRepository())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
