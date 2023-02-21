package utils

const (
	INSERT_CUSTOMER      = "INSERT INTO mst_customer (name,balance) values ($1,$2)"
	SELECT_CUSTOMER_LIST = "SELECT * FROM mst_customer"

	SELECT_STORE_LIST      = "SELECT * FROM mst_store"
	SELECT_PRODUK_LIST     = "SELECT * FROM mst_product"
	SELECT_PRODUCT_STOREID = "SELECT * FROM mst_product where store_id= $1"
)
