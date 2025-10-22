package agent

import "time"

type LoginAgent struct {
	AgentID       string `db:"agent_id" json:"agent_id"`
	AgentPassword string `db:"password" json:"password"`
}

type Agent struct {
	AgentID       string    `db:"agent_id" json:"agent_id"`
	AgentName     string    `db:"agent_name" json:"agent_name"`
	AgentPassword string    `db:"password" json:"password"`
	Active        bool      `db:"active" json:"active"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type AgentRequest struct {
	AgentID       string    `db:"agent_id" json:"agent_id"`
	AgentName     string    `db:"agent_name" json:"agent_name"`
	AgentPassword string    `db:"password" json:"password"`
	Active        string      `db:"active" json:"active"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type AgentResp struct {
	AgentID       string    `db:"agent_id" json:"AgentId"`
	AgentName     string    `db:"agent_name" json:"AgentName"`
}

type ResponseLogin struct {
	AgentID    string `db:"agent_id" json:"AgentId"`
	AgentName  string `db:"agent_name" json:"AgentName"`
	AgentToken string `json:"AccessToken"`
}
