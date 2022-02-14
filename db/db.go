package db

import (
	"context"

	"github.com/essayZW/hpcmanager/config"
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

// TransicationConn 返回给定ctx中的事务连接，没有则新建一个
func (db *DB) TransicationConn(ctx context.Context) (*sqlx.Tx, error) {
	conn := ctx.Value(transactionConnCtxKey)
	if conn == nil {
		return db.conn.BeginTxx(ctx, nil)
	}
	if tx, ok := conn.(*sqlx.Tx); ok {
		return tx, nil
	}
	return db.conn.BeginTxx(ctx, nil)
}

// TransicationHandler 事务处理的函数定义
type TransicationHandler func(context.Context, ...interface{}) (interface{}, error)

// Transication 在事务中进行函数的执行，如果执行的函数出现panic或者返回error，则进行回滚，否则自动进行提交
// 该函数会返回do函数的返回值，因此do函数必须为第一个返回值为data，第二个返回值为error
func (db *DB) Transication(ctx context.Context, do TransicationHandler, param ...interface{}) (doReturnValue interface{}, doError error) {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, transactionConnCtxKey, tx)

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
		if doError != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	doReturnValue, doError = do(ctx, param...)
	return doReturnValue, doError
}

// NewDB 创建一个新的数据库连接
func NewDB() (*DB, error) {
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
	return &DB{
		conn: db,
	}, nil
}
