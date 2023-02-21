package repository

import (
	"database/sql"
	"godb/model"
	"godb/utils"
)

type StoreRepository interface {
	GetAll() ([]model.StoreProducts, error)
}

type storeRepository struct {
	db *sql.DB
}

func (st *storeRepository) GetAll() ([]model.StoreProducts, error) {
	rows, err := st.db.Query(utils.SELECT_STORE_LIST)

	if err != nil {
		return nil, err
	}

	var stores []model.StoreProducts
	for rows.Next() {
		var store model.StoreProducts

		var siupNumber string
		err := rows.Scan(
			&store.Id,
			&siupNumber,
			&store.StoreName,
		)

		var products []model.Product
		rows, errProd := st.db.Query(utils.SELECT_PRODUCT_STOREID, store.Id)
		if errProd != nil {
			return nil, err
		}

		for rows.Next() {
			var product model.Product
			err := rows.Scan(
				&product.Id,
				&product.Nama,
				&product.Price,
				&product.Stock,
				&product.Create_ad,
				&product.Store_id,
			)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}

		store.Products = products
		stores = append(stores, store)

	}
	return stores, nil
}

func NewStoreRepository(dbConnect *sql.DB) StoreRepository {
	return &storeRepository{
		db: dbConnect,
	}
}
