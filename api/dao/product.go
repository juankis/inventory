package dao

import (
	"fmt"

	"github.com/juankis/inventory/api/models"
)

//InsertProduct insert
func InsertProduct(product *models.Product) error {
	res, err := Db.NamedExec(`INSERT INTO product (name) VALUES (:name)`, &product)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error inserting product to the database, product: %v\n", product), err, nil)
		return err
	}
	id, _ := res.LastInsertId()
	product.ID = int(id)
	return nil
}

//GetProductAll Returns error and product if exist
func GetProductAll() ([]models.Product, error) {
	var product []models.Product
	err := Db.Select(&product, "SELECT * FROM `product` ORDER BY id ASC")
	if err != nil {
		fmt.Errorf("Error getting product from database %v \n", err)
		return product, err
	}
	return product, nil
}

//GetProduct Returns error and product if exist
func GetProduct(product *models.Product) error {
	err := Db.Get(product, "SELECT * FROM `product` where id = ? limit 1", product.ID)
	if err != nil {
		fmt.Errorf("Error getting product from database %v \n", err)
		return err
	}
	return nil
}

//UpdateProduct update product
func UpdateProduct(product *models.Product) error {
	fmt.Printf("update product : %+v", product)
	_, err := Db.NamedExec(`UPDATE product set name = :name
											where id = :id`, &product)
	if err != nil {
		fmt.Errorf("Error updating product in database", err)
		return err
	}
	return nil
}

//DeleteProduct delete product
func DeleteProduct(product *models.Product) error {
	_, err := Db.Exec("DELETE FROM product WHERE id = ?", product.ID)
	if err != nil {
		fmt.Errorf("Error deleting product from database", err, nil)
		return err
	}
	return nil
}
