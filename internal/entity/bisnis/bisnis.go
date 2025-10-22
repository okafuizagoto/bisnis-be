package bisnis

import "time"

type Transaction struct {
	TransID   int       `db:"trans_id" json:"trans_id"`
	AgentID   string    `db:"agent_id" json:"agent_id"`
	ProductID string    `db:"product_id" json:"product_id"`
	Nama      string    `db:"nama" json:"nama"`
	Usia      int       `db:"usia" json:"usia"`
	Premium   float64   `db:"premium" json:"premium"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
type AddTransaction struct {
	AgentID   string  `db:"agent_id" json:"agent_id"`
	ProductID string  `db:"product_id" json:"product_id"`
	Nama      string  `db:"nama" json:"nama"`
	Usia      int     `db:"usia" json:"usia"`
	Premium   float64 `db:"premium" json:"premium"`
}

type UpdateTransaction struct {
	TransID   int     `db:"trans_id" json:"trans_id"`
	AgentID   string  `db:"agent_id" json:"agent_id"`
	ProductID string  `db:"product_id" json:"product_id"`
	Nama      string  `db:"nama" json:"nama"`
	Usia      int     `db:"usia" json:"usia"`
	Premium   float64 `db:"premium" json:"premium"`
}

type DeleteTransaction struct {
	AgentID string `db:"agent_id" json:"agent_id"`
	TransID int    `db:"trans_id" json:"trans_id"`
}

type TransactionResp struct {
	TransID string `db:"trans_id" json:"TransId"`
}
