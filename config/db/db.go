package db

import (
	"context"
	"errors"
	"fmt"
	"mereetmoi/pkg/vip"
	"mereetmoi/pkg/zlog"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	ErrTx        = " error transaction"
	ErrCom       = " error commit transaction"
	ErrRollback  = " error to rollback"
	ExecRollback = " execute rollback"
	source       = " auth db package "
)

var (
	ErrToRollback   = errors.New(" | error to execute rollback")
	ExecToRollback  = errors.New(" | execute rollback")
	ErrToCommit     = errors.New(" | error to commit")
	ErrNoAffected   = errors.New(" | error no row err no affected")
	ErrUserNotFound = errors.New(" | error user not found")
)

type DbService struct {
	DB  *pgxpool.Pool
	TX  pgx.Tx
	Err error
}

func NewDbService(ctx context.Context) *DbService {
	var errServDB error
	DB, errDB := New(ctx).DB()
	TX, errTX := New(ctx).TX()
	if errDB != nil {
		errServDB = errDB
	}
	if errTX != nil {
		errServDB = errTX
	}

	return &DbService{
		DB:  DB,
		TX:  TX,
		Err: errServDB,
	}
}

type PgConfig struct {
	Ctx context.Context
}

func New(ctx context.Context) *PgConfig {
	return &PgConfig{
		Ctx: ctx,
	}
}

var newErr error

func (p *PgConfig) TX() (pgx.Tx, error) {
	DB, err := p.DB()
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	TX, errTX := DB.Begin(p.Ctx)
	if errTX != nil {
		zlog.Error(err)
		return nil, errTX
	}
	return TX, nil
}

func (p *PgConfig) DB() (*pgxpool.Pool, error) {
	pool, errDB := p.Conn()
	if errDB != nil {
		zlog.Error(errDB)
		newErr = errDB
	}
	acquire, err := pool.Acquire(p.Ctx)
	if err != nil {
		zlog.Error(err)
		return nil, err
	}
	defer acquire.Release()

	return pool, newErr
}

func (p *PgConfig) Conn() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(p.Ctx, 10*time.Second)
	defer cancel()
	vipConf, errVip := vip.New().App()
	if errVip != nil {
		return nil, errVip
	}

	dbURL, errURL := p.dbUrl()
	if errURL != nil {
		return nil, errURL
	}
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		zlog.Error(err)
		newErr = err
	}

	config.MaxConns = int32(vipConf.DbMaxCon)
	config.MaxConnLifetime = time.Duration(vipConf.DbMaxLifetime)

	pool, errPool := pgxpool.NewWithConfig(ctx, config)
	if errPool != nil {
		zlog.Error(err)
		newErr = errPool
	}
	_, errConnect := pool.Exec(ctx, ";")
	if errConnect != nil {
		zlog.Error(err)
		newErr = errConnect
	}
	return pool, newErr
}

func (p *PgConfig) dbUrl() (string, error) {
	vipConf, errVip := vip.New().App()
	if errVip != nil {
		return "", errVip
	}

	dbUser := vipConf.DbUsername
	dbPass := vipConf.DbPassword
	dbHost := vipConf.DbHost
	dbPort := strconv.Itoa(vipConf.DbPort)
	dbName := vipConf.DbName
	dbSchema := vipConf.DbSchema
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		dbSchema)
	return url, nil
}

func RollbackTX(ctx context.Context, tx pgx.Tx) error {
	errRoll := tx.Rollback(ctx)
	if errRoll != nil {
		return errRoll
	}
	return nil
}

func ExecCommit(ctx context.Context, TX pgx.Tx) error {
	errComm := TX.Commit(ctx)
	if errComm != nil {
		return errors.Join(ErrToCommit, errComm)
	}
	return nil
}

func RollCommit(err error, ctx context.Context, TX pgx.Tx) error {
	if err != nil {
		errRoll := RollbackTX(ctx, TX)
		if errRoll != nil {
			return errors.Join(ErrToRollback, errRoll)
		}
		return errors.Join(ExecToRollback, err)
	}

	errComm := TX.Commit(ctx)
	if errComm != nil {
		return errors.Join(ErrToCommit, errComm)
	}
	return nil
}

func ErrExec(err error, exec pgconn.CommandTag) error {
	if err != nil {
		return err
	}
	if exec.RowsAffected() == 0 {
		return ErrNoAffected
	}
	return nil
}

func Ilike(key string) string {
	return "%" + key + "%"
}
