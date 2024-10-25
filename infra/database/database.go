package database

import (
	"context"
	"fmt"

	"bookstore/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	sqlxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/jmoiron/sqlx"
)

type Block func(tx *sqlx.Tx) error

type PostgreSQL struct {
	db *sqlx.DB
}

func NewPostgreSQL(cfg *config.PostgresConfig) (*PostgreSQL, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
	db, err := sqlxtrace.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgreSQL{db: db}, nil
}

func (p *PostgreSQL) WithTransaction(ctx context.Context, block Block) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("can't start DB transaction: %w", err)
	}
	err = block(tx)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return fmt.Errorf("rollback fails: %w", err)
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("transaction commit fails: %w", err)
	}
	return nil
}

func (p *PostgreSQL) DB() *sqlx.DB {
	return p.db
}

func (p *PostgreSQL) ArrayAggregate(field, alias string) string {
	return fmt.Sprintf("ARRAY_AGG(%s ORDER BY %s) AS %s", field, field, alias)
}
