package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/jmoiron/sqlx"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// key 存储到 context中的键值的类型（为了消除某个warning)
// https://stackoverflow.com/questions/40891345/fix-should-not-use-basic-type-string-as-key-in-context-withvalue-golint
type key string

// transactionConnCtxKey 事务连接在context中的键值
const transactionConnCtxKey key = "transactionConnCtxKey"

// DB 数据库操作对象，持有一个数据库连接
type DB struct {
	conn *sqlx.DB
}

// TransactionConn 返回给定ctx中的事务连接，没有则新建一个
func (db *DB) TransactionConn(ctx context.Context) (*sqlx.Tx, error) {
	conn := ctx.Value(transactionConnCtxKey)
	if conn == nil {
		// 说明不在事务当中
		return nil, errors.New("No transaction conn")
	}
	if tx, ok := conn.(*sqlx.Tx); ok {
		return tx, nil
	}
	return nil, errors.New("No transaction conn")
}

// TransactionHandler 事务处理的函数定义
type TransactionHandler func(context.Context, ...interface{}) (interface{}, error)

// Transaction 在事务中进行函数的执行，如果执行的函数出现panic或者返回error，则进行回滚，否则自动进行提交
// 该函数会返回do函数的返回值，因此do函数必须为第一个返回值为data，第二个返回值为error
func (db *DB) Transaction(ctx context.Context, do TransactionHandler, param ...interface{}) (doReturnValue interface{}, doError error) {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, transactionConnCtxKey, tx)

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			logger.Debug("Transaction rollback because panic error: ", err)
			panic(err)
		}
		if doError != nil {
			tx.Rollback()
			logger.Debug("Transaction rollback because do func error: ", doError)
			return
		}
		if err := tx.Commit(); err != nil {
			logger.Debug("Transaction commit error: ", err)
			tx.Rollback()
			return
		}
		logger.Debug("Transaction commit")
	}()
	doReturnValue, doError = do(ctx, param...)
	return doReturnValue, doError
}

// QueryRow 执行单行查询
func (db *DB) QueryRow(ctx context.Context, query string, params ...interface{}) (*sqlx.Row, error) {
	tx, err := db.TransactionConn(ctx)
	if err != nil {
		// 说明不在事务当中，应该直接使用conn执行
		return db.conn.QueryRowxContext(ctx, query, params...), nil
	}
	return tx.QueryRowxContext(ctx, query, params...), nil
}

// Query 执行查询语句
func (db *DB) Query(ctx context.Context, query string, params ...interface{}) (*sqlx.Rows, error) {
	tx, err := db.TransactionConn(ctx)
	if err != nil {
		return db.conn.QueryxContext(ctx, query, params...)
	}
	return tx.QueryxContext(ctx, query, params...)
}

// Exec 执行某个SQL语句
func (db *DB) Exec(ctx context.Context, query string, params ...interface{}) (sql.Result, error) {
	tx, err := db.TransactionConn(ctx)
	if err != nil {
		return db.conn.ExecContext(ctx, query, params...)
	}
	return tx.ExecContext(ctx, query, params...)
}

// defaultDB 默认已经创建的DB
var defaultDB *DB

// NewDB 创建一个新的数据库连接
func NewDB() (*DB, error) {
	if defaultDB != nil {
		return defaultDB, nil
	}
	dbConfig, err := config.LoadDatabase()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("mysql", dbConfig.Dsn())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	defaultDB = &DB{
		conn: db,
	}
	return defaultDB, nil
}

// Transaction 使用默认DB初始化进行事务操作
func Transaction(ctx context.Context, do TransactionHandler, params ...interface{}) (interface{}, error) {
	if defaultDB == nil {
		return nil, errors.New("Must call NewDB before use Transaction")
	}
	return defaultDB.Transaction(ctx, do, params...)
}
