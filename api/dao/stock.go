package dao

import (
	"fmt"

	"github.com/juankis/inventory/api/models"
)

//InsertStock insert
func InsertStock(stock *models.Stock) error {
	res, err := Db.NamedExec(`INSERT INTO stock (product_id, store_id, stock) VALUES (:product_id, :store_id, :stock)`, &stock)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error inserting stock to the database, stock: %v\n", stock), err, nil)
		return err
	}
	id, _ := res.LastInsertId()
	stock.ID = int(id)
	fmt.Printf("stock in insertStock(): %+v\n", stock)
	return nil
}

//GetStockAll Returns error and stock if exist
func GetStockAll() ([]models.Stock, error) {
	var stock []models.Stock
	err := Db.Select(&stock, "SELECT id, product_id, store_id, stock FROM `stock` ORDER BY id ASC")
	if err != nil {
		fmt.Errorf("Error getting stock from database %v \n", err)
		return stock, err
	}
	return stock, nil
}

//GetStockByID Returns error and stock if exist
func GetStockByID(stock *models.Stock) error {
	err := Db.Get(stock, "SELECT * FROM `stock` where id = ? limit 1", stock.ID)
	if err != nil {
		fmt.Errorf("Error getting stock from database %v \n", err)
		return err
	}
	return nil
}

//GetStock Returns error and stock if exist
func GetStock(stock *models.Stock) error {
	err := Db.Get(stock, "SELECT * FROM `stock` where store_id = ? and product_id = ? limit 1", stock.StoreID, stock.ProductID)
	if err != nil {
		err = InsertStock(stock)
		fmt.Printf("Stock in getStock(): %+v\n", stock)
		if err != nil {
			return err
		}
	}
	return nil
}

//UpdateStock update stock
func UpdateStock(transaction *models.Transaction) error {
	stockFrom := &models.Stock{ProductID: transaction.ProductId, StoreID: transaction.StoreIDFrom}
	stockTo := &models.Stock{ProductID: transaction.ProductId, StoreID: transaction.StoreIDTo}
	err := GetStock(stockFrom)
	fmt.Printf("Stock : %+v \n", stockFrom)
	if err != nil {
		fmt.Errorf("Error getting stock from database %v \n", err)
		return err
	}
	err = GetStock(stockTo)
	fmt.Printf("Stock : %+v \n", stockTo)
	if err != nil {
		fmt.Errorf("Error getting stock from database %v \n", err)
		return err
	}

	stockFrom.Stock = stockFrom.Stock - transaction.Quantity
	stockTo.Stock = stockTo.Stock + transaction.Quantity

	_, err = Db.NamedExec(`UPDATE stock set stock = :stock where id = :id`, &stockFrom)
	if err != nil {
		fmt.Errorf("Error updating stock in database", err)
		return err
	}

	_, err = Db.NamedExec(`UPDATE stock set stock = :stock where id = :id`, &stockTo)
	if err != nil {
		fmt.Errorf("Error updating stock in database", err)
		return err
	}

	return nil
}
