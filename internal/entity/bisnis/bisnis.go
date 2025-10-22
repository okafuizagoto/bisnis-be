package bisnis

type AddTransaction struct {
	AgentID   string  `db:"agent_id" json:"agent_id"`
	ProductID string  `db:"product_id" json:"product_id"`
	Nama      string  `db:"nama" json:"nama"`
	Usia      int     `db:"usia" json:"usia"`
	Premium   float64 `db:"premium" json:"premium"`
}
