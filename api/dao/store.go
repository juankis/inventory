package dao

import (
	"fmt"

	"github.com/juankis/inventory/api/models"
)

//InsertStore insert
func InsertStore(store *models.Store) error {
	res, err := Db.NamedExec(`INSERT INTO store (name) VALUES (:name)`, &store)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error inserting store to the database, store: %v\n", store), err, nil)
		return err
	}
	id, _ := res.LastInsertId()
	store.ID = int(id)
	return nil
}

//GetStoreAll Returns error and store if exist
func GetStoreAll() ([]models.Store, error) {
	var store []models.Store
	err := Db.Select(&store, "SELECT * FROM `store` ORDER BY id ASC")
	if err != nil {
		fmt.Errorf("Error getting store from database %v \n", err)
		return store, err
	}
	return store, nil
}

//GetStore Returns error and store if exist
func GetStore(store *models.Store) error {
	err := Db.Get(store, "SELECT * FROM `store` where id = ? limit 1", store.ID)
	if err != nil {
		fmt.Errorf("Error getting store from database %v \n", err)
		return err
	}
	return nil
}

//UpdateStore update store
func UpdateStore(store *models.Store) error {
	fmt.Printf("update store : %+v", store)
	_, err := Db.NamedExec(`UPDATE store set name = :name
											where id = :id`, &store)
	if err != nil {
		fmt.Errorf("Error updating store in database", err)
		return err
	}
	return nil
}

//DeleteStore delete store
func DeleteStore(store *models.Store) error {
	_, err := Db.Exec("DELETE FROM store WHERE id = ?", store.ID)
	if err != nil {
		fmt.Errorf("Error deleting store from database", err, nil)
		return err
	}
	return nil
}
