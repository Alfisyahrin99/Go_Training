package repository

import (
	"database/sql"
	"godb/model"
	"godb/utils"
)

type ProdukRepository interface {
	GetAll() ([]model.Product, error)
}

type produkRepository struct {
	db *sql.DB
}

func (pr *produkRepository) GetAll() ([]model.Product, error) {
	rows, err := pr.db.Query(utils.SELECT_PRODUK_LIST)

	if err != nil {
		return nil, err
	}

	var produks []model.Product
	for rows.Next() {
		var produk model.Product

		err := rows.Scan(
			&produk.Id,
			&produk.Nama,
			&produk.Price,
			&produk.Stock,
			&produk.Create_ad,
			&produk.Store_id,
		)

		if err != nil {

			return nil, err
		}

		produks = append(produks, produk)

	}
	return produks, nil
}

func NewProdukRepository(dbConnect *sql.DB) ProdukRepository {
	return &produkRepository{
		db: dbConnect,
	}
}
