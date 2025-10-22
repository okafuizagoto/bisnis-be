package agent

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	jaegerLog "bisnis-be/pkg/log"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt *map[string]*sqlx.Stmt

		tracer opentracing.Tracer
		logger jaegerLog.Factory
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

// Tambahkan query di dalam const
// getAllUser = "GetAllUser"
// qGetAllUser = "SELECT * FROM users"
const (
	// getGoldUser  = "GetGoldUser"
	// qGetGoldUser = `SELECT gold_email,gold_password,gold_nama,gold_nomorhp,gold_nomorkartu,gold_cvv,gold_expireddate,gold_namapemegangkartu FROM data_peserta`

	getAgentID  = "GetAgentID"
	qGetAgentID = `select agent_id, agent_name, password, active, created_at from agen where agent_id = ?`

	validationAgentPassword  = "ValidationAgentPassword"
	qValidationAgentPassword = `select agent_id, agent_name, password, active, created_at from agen where agent_id = ? and password = ?`

	insertAgent  = "InsertAgent"
	qInsertAgent = `INSERT INTO bisnis.agen
(agent_id, agent_name, password, active, created_at) 
select ?, ?, ?, CASE WHEN 'Y' = ? THEN 1 ELSE 0 END, NOW()`

	getAgentIDName  = "GetAgentIDName"
	qGetAgentIDName = `select agent_id, agent_name, password, active, created_at from agen where agent_id = ? or agent_name like ?`

	deleteAgent  = "DeleteAgent"
	qDeleteAgent = `DELETE FROM agen where agent_id = ?`

	updateAgent  = "UpdateAgent"
	qUpdateAgent = `UPDATE agen set agent_name = ? where agent_id = ?`

	getAgentName  = "GetAgentName"
	qGetAgentName = `select agent_id, agent_name, password, active, created_at from agen where agent_name like ?`
)

var (
	readStmt = []statement{
		{getAgentID, qGetAgentID},
		{validationAgentPassword, qValidationAgentPassword},
		{getAgentIDName, qGetAgentIDName},
		{getAgentName, qGetAgentName},
	}
	insertStmt = []statement{
		{insertAgent, qInsertAgent},
	}
	updateStmt = []statement{
		{updateAgent, qUpdateAgent},
	}
	deleteStmt = []statement{
		{deleteAgent, qDeleteAgent},
	}
)

// New ...
func New(db *sqlx.DB, tracer opentracing.Tracer, logger jaegerLog.Factory) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)
	d := &Data{
		db:     db,
		tracer: tracer,
		logger: logger,
		stmt:   &stmts,
	}

	d.InitStmt()
	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	*d.stmt = stmts
}
