package models

//Stock struct
type Stock struct {
	ID        int    `json:"id,omitempty" db:"id"`
	ProductID int    `json:"product_id" db:"product_id" binding:"required"`
	StoreID   int    `json:"store_id" db:"store_id" binding:"required"`
	Stock     int    `json:"stock" db:"stock" binding:"required"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}
