package pg

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	db *pgxpool.Pool
}

var (
	PgClient *PostgresDB
	pgError  error
	pgOnce   sync.Once
)

func ConnectPostgres(ctx context.Context, connString string) (*PostgresDB, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			pgError = fmt.Errorf("unable to create connection pool: %w", err)
		}

		PgClient = &PostgresDB{db}
	})

	return PgClient, pgError
}

func (pg *PostgresDB) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *PostgresDB) Close() {
	pg.db.Close()
}

// type PostgresDB struct {
// 	db *pgxpool.Pool
// }

// var (
// 	pgInstance *PostgresDB
// 	pgError    error
// 	pgOnce     sync.Once
// )

// func NewPG(ctx context.Context, connString string) (*PostgresDB, error) {
// 	pgOnce.Do(func() {
// 		db, err := pgxpool.New(ctx, connString)
// 		if err != nil {
// 			pgError = fmt.Errorf("unable to create connection pool: %w", err)
// 		}

// 		pgInstance = &PostgresDB{db}
// 	})

// 	return pgInstance, pgError
// }

// func (pg *PostgresDB) Ping(ctx context.Context) error {
// 	return pg.db.Ping(ctx)
// }

// func (pg *PostgresDB) Close() {
// 	pg.db.Close()
// }
