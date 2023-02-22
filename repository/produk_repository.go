package repository

import (
	"database/sql"
	"godb/model"
	"godb/utils"
)

type ProductRepository interface {
	GetAll() ([]model.Product, error)
	GetByStoreId(id string) ([]model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) GetAll() ([]model.Product, error) {
	rows, err := p.db.Query(utils.SELECT_PRODUCT_LIST)
	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var product model.Product
		rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.StoreId,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) GetByStoreId(id string) ([]model.Product, error) {
	rows, err := p.db.Query(utils.SELECT_PRODUCT_STOREID, id)
	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var product model.Product
		rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.StoreId,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
