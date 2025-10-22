package bisnis

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/storage"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	jaegerLog "bisnis-be/pkg/log"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		fsdb *firestore.Client
		s    *storage.Client
		rdb  *redis.Client
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
// getOneStockProduct  = "GetOneStockProduct"
// qGetOneStockProduct = `SELECT stock_id, stock_code, stock_name, stock_pack,stock_qty, stock_price, stock_last_update, stock_update_by FROM stock WHERE stock_code = ? AND (? = "" OR stock_name LIKE ?) and (? = "" OR stock_id = ?)`
)

var (
	readStmt = []statement{
		// {getOneStockProduct, qGetOneStockProduct},
	}
	insertStmt = []statement{}
	updateStmt = []statement{}
	deleteStmt = []statement{}
)

// New ...
// func New(db *sqlx.DB, fsdb *db.Client, fs *storage.Client, rdb *redis.Client, tracer opentracing.Tracer, logger jaegerLog.Factory) *Data {
func New(db *sqlx.DB, fsdb *firestore.Client, fs *storage.Client, rdb *redis.Client, tracer opentracing.Tracer, logger jaegerLog.Factory) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)
	d := &Data{
		db:     db,
		fsdb:   fsdb,
		s:      fs,
		rdb:    rdb,
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
