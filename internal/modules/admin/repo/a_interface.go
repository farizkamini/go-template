package repo_admin_auth

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Read struct {
	DB *pgxpool.Pool
}

func NewRead(db *pgxpool.Pool) *Read {
	return &Read{DB: db}
}

type Write struct {
	TX pgx.Tx
}

func NewWrite(tx pgx.Tx) *Write {
	return &Write{TX: tx}
}
