package dao

import (
	"fmt"

	"github.com/juankis/inventory/api/models"
)

//InsertUser insert
func InsertUser(user *models.User) error {
	res, err := Db.NamedExec(`INSERT INTO user (name, user, password) VALUES (:name, :user, :password)`, &user)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error inserting user to the database, user: %v\n", user), err, nil)
		return err
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return nil
}

//GetUserAll Returns error and user if exist
func GetUserAll() ([]models.User, error) {
	var user []models.User
	err := Db.Select(&user, "SELECT * FROM `user` ORDER BY id ASC")
	if err != nil {
		fmt.Errorf("Error getting user from database %v \n", err)
		return user, err
	}
	return user, nil
}

//GetUser Returns error and user if exist
func GetUser(user *models.User) error {
	err := Db.Get(user, "SELECT * FROM `user` where id = ? limit 1", user.ID)
	if err != nil {
		fmt.Errorf("Error getting user from database %v \n", err)
		return err
	}
	return nil
}

//UpdateUser update user
func UpdateUser(user *models.User) error {
	fmt.Printf("update user : %+v", user)
	_, err := Db.NamedExec(`UPDATE user set name = :name,
											user = :user,
											password = :password
											where id = :id`, &user)
	if err != nil {
		fmt.Errorf("Error updating user in database", err)
		return err
	}
	return nil
}

//DeleteUser delete user
func DeleteUser(user *models.User) error {
	_, err := Db.Exec("DELETE FROM user WHERE id = ?", user.ID)
	if err != nil {
		fmt.Errorf("Error deleting user from database", err, nil)
		return err
	}
	return nil
}

//Login insert
func Login(user *models.LoginRequest) bool {
	query := fmt.Sprintf("SELECT * FROM `user` where user = '%s' and password = '%s' limit 1", user.User, user.Password)
	return RowExists(query)
}
