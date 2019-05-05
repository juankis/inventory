package dao

import (
	"fmt"

	"github.com/juankis/inventory/api/models"
)

//InsertTransaction insert
func InsertTransaction(transaction *models.Transaction) error {
	res, err := Db.NamedExec(`INSERT INTO transaction (quantity, movement, user_creator, product_id) VALUES (:quantity, :movement, :user_creator, :product_id)`, &transaction)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error inserting transaction to the database, transaction: %v\n", transaction), err, nil)
		return err
	}
	id, _ := res.LastInsertId()
	transaction.ID = int(id)
	return nil
}

//GetTransactionAll Returns error and transaction if exist
func GetTransactionAll() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := Db.Select(&transaction, "SELECT t.*, p.name as `product_name`, u.name as `user_creator_name` , u2.name as `user_confirm_name` FROM `transaction` as t LEFT JOIN `product` as p ON t.product_id = p.id LEFT JOIN `user` as u ON t.user_creator = u.id LEFT JOIN `user` as u2 ON t.user_confirm = u2.id  ORDER BY t.id ASC")

	if err != nil {
		fmt.Errorf("Error getting transaction from database %v \n", err)
		return transaction, err
	}
	return transaction, nil
}

//GetTransaction Returns error and transaction if exist
func GetTransaction(transaction *models.Transaction) error {
	err := Db.Get(transaction, "SELECT * FROM `transaction` where id = ? limit 1", transaction.ID)
	if err != nil {
		fmt.Errorf("Error getting transaction from database %v \n", err)
		return err
	}
	return nil
}

//UpdateTransaction update transaction
func UpdateTransaction(transaction *models.Transaction) error {
	fmt.Printf("update transaction : %+v", transaction)
	_, err := Db.NamedExec(`UPDATE transaction set quantity = :quantity,
											movement = :movement,
											user_creator= :user_creator,
											user_confirm = :user_confirm,
											product_id = :product_id,
											updated_at = CURRENT_TIMESTAMP
											where id = :id`, &transaction)
	if err != nil {
		fmt.Errorf("Error updating transaction in database", err)
		return err
	}
	return nil
}

//DeleteTransaction delete Transaction
func DeleteTransaction(transaction *models.Transaction) error {
	_, err := Db.Exec("DELETE FROM transaction WHERE id = ?", transaction.ID)
	if err != nil {
		fmt.Errorf("Error deleting transaction from database", err, nil)
		return err
	}
	return nil
}
