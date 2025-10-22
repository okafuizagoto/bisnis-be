package product

import "time"

type Product struct {
	ProductID   string    `db:"product_id" json:"product_id"`
	ProductName string    `db:"product_name" json:"product_name"`
	Premium     float64   `db:"premium" json:"premium"`
	Active      bool      `db:"active" json:"active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
