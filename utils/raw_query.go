package utils

const (
	INSERT_CUSTOMER      = "INSERT INTO mst_customer (name,balance) values ($1,$2)"
	SELECT_CUSTOMER_LIST = "SELECT * FROM mst_customer"
	SELECT_CUSTOMER_ID   = "SELECT * FROM mst_customer where id = $1"
	UPDATE_CUSTOMER      = "UPDATE mst_customer set name=$1, balance=$2 where id=$3"
	DELETE_CUSTOMER      = "DELETE FROM mst_customer where id = $1"

	//STORE
	SELECT_STORE_LIST = "SELECT * FROM mst_store"

	//PRODUCT
	INSERT_PRODUCT         = "INSERT INTO product (id,name,price,stock) values (:id,:name,:price,:stock)"
	SELECT_PRODUCT_LIST    = "SELECT * FROM product"
	SELECT_PRODUCT_STOREID = "SELECT * FROM product where store_id = $1"
)
