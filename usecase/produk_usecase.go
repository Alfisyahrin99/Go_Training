package usecase

import (
	"godb/model"
	"godb/repository"
)

type ProdukUseCase interface {
	GetAllProduk() ([]model.Product, error)
}

type produkUseCase struct {
	produkRepo repository.ProdukRepository
}

func (pr *produkUseCase) GetAllProduk() ([]model.Product, error) {
	return pr.produkRepo.GetAll()
}

func NewProdukUseCase(produkRepository repository.ProdukRepository) ProdukUseCase {
	return &produkUseCase{
		produkRepo: produkRepository,
	}
}
