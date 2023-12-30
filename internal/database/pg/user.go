package pg

import (
	"context"
	"fmt"

	"github.com/ididie/ifidie_backend/internal/model"
	"github.com/jackc/pgx/v5"
)

func (pg *PostgresDB) GetUser(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT * FROM public."user"  WHERE id = $1 LIMIT 1`

	rows, _ := pg.db.Query(ctx, query, id)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])

	if err != nil {
		return nil, fmt.Errorf("unable to query user: %w", err)
	}
	defer rows.Close()

	if len(users) != 1 {
		return nil, fmt.Errorf("user not found. user.id: %d", id)
	}

	fmt.Printf("FOUND: %s", users[0].FirstName)

	return &users[0], nil
}
