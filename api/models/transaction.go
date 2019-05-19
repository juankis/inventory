package models

//Transaction struct
type Transaction struct {
	ID              int     `json:"id,omitempty" db:"id"`
	Quantity        int     `json:"quantity" db:"quantity" binding:"required,number"`
	Movement        int     `json:"movement,omitempty" db:"movement" binding:"required,number"`
	UserCreator     *int    `json:"user_creator" db:"user_creator" binding:"required,number"`
	UserConfirm     *int    `json:"user_confirm" db:"user_confirm"`
	UserCreatorName *string `json:"user_creator_name" db:"user_creator_name"`
	UserConfirmName *string `json:"user_confirm_name" db:"user_confirm_name"`
	ProductId       int     `json:"product_id" db:"product_id"`
	ProductName     *string `json:"product_name" db:"product_name"`
	StoreIDFrom     int     `json:"store_id_from" db:"store_id_from"`
	StoreIDTo       int     `json:"store_id_to" db:"store_id_to"`
	CreatedAt       string  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       string  `json:"updated_at,omitempty" db:"updated_at"`
}

//TransactionResponse struct
type TransactionResponse struct {
	Quantity string  `json:"quantity"`
	Movement string  `json:"movement"`
	Creator  *string `json:"user_creator_name"`
	Product  *string `json:"product_name"`
}

//Response struct
type Response struct {
	Data Transaction `json:"data"`
}

//ResponseRegitries struct
type ResponseRegitries struct {
	Data []Transaction `json:"data"`
}

//ConfirmTransaction struct
type ConfirmTransaction struct {
	TransactionID int  `json:"transaction_id"`
	Confirm       bool `json:"confirm"`
	ConfirmUserID int  `json:"confirm_user_id"`
}

//TransactionRequest struct
type TransactionRequest struct {
	Quantity    string `json:"quantity" binding:"required"`
	Movement    string `json:"movement" binding:"required"`
	UserCreator string `json:"user_creator" binding:"required"`
	ProductId   string `json:"product_id" binding:"required"`
	StoreIDFrom string `json:"store_id_from" binding:"required"`
	StoreIDTo   string `json:"store_id_to" binding:"required"`
}
