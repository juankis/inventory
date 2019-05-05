package models

//Transaction struct
type Transaction struct {
	ID              int     `json:"id,omitempty" db:"id"`
	Quantity        int     `json:"quantity" db:"quantity" binding:"required"`
	Movement        int     `json:"movement,omitempty" db:"movement" binding:"required"`
	UserCreator     *int    `json:"user_creator" db:"user_creator" binding:"required"`
	UserConfirm     *int    `json:"user_confirm" db:"user_confirm"`
	UserCreatorName *string `json:"user_creator_name" db:"user_creator_name"`
	UserConfirmName *string `json:"user_confirm_name" db:"user_confirm_name"`
	ProductId       int     `json:"product_id" db:"product_id"`
	ProductName     *string `json:"product_name" db:"product_name"`
	CreatedAt       string  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       string  `json:"updated_at,omitempty" db:"updated_at"`
}

//Response struct
type Response struct {
	Data Transaction `json:"data"`
}

//ResponseRegitries struct
type ResponseRegitries struct {
	Data []Transaction `json:"data"`
}
