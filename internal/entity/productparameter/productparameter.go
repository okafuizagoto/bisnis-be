package productparameter

import "time"

type ProductParameter struct {
	ID             string    `db:"id" json:"id"`
	ProductID      string    `db:"product_id" json:"product_id"`
	ParameterName  string    `db:"parameter_name" json:"parameter_name"`
	ParameterValue string    `db:"parameter_value" json:"parameter_value"`
	Active         bool      `db:"active" json:"active"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
