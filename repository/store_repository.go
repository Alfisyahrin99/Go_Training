package repository

import (
	"database/sql"
	"godb/model"
	"godb/utils"
)

type StoreRepository interface {
	GetAll() ([]model.Store, error)
}

type storeRepository struct {
	db *sql.DB
}

func (st *storeRepository) GetAll() ([]model.Store, error) {
	rows, err := st.db.Query(utils.SELECT_STORE_LIST)
	if err != nil {
		return nil, err
	}
	// Tempat untuk menampung slice dengan tipe data model store
	var stores []model.Store
	for rows.Next() {
		// untuk mengambil setiap record dari Store
		var store model.Store
		// scan/masukan nilai hasil dari query kedalam pointer variable store
		err := rows.Scan(
			&store.Id,
			&store.SiupNumber,
			&store.Name,
		)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	return stores, nil
}

func NewStoreRepository(dbConnect *sql.DB) StoreRepository {
	return &storeRepository{
		db: dbConnect,
	}
}
