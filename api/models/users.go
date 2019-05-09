package models

//User struct
type User struct {
	ID        int    `json:"id,omitempty" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	User      string `json:"user" db:"user" binding:"required"`
	Password  string `json:"password" db:"password" binding:"required"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

//LoginRequest struct
type LoginRequest struct {
	User     string `json:"user" binding:"required" schema:"user"`
	Password string `json:"password" binding:"required" schema:"password"`
}
