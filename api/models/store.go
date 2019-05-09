package models

//Store struct
type Store struct {
	ID        int    `json:"id,omitempty" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}
