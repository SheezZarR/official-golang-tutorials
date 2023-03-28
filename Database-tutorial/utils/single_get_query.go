package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"databases/tutorial/gotypes"
)

// Queries for the single album id
func AlbumById(dbConnection *pgx.Conn, id int64) (gotypes.Album, error) {
	var alb gotypes.Album

	row := dbConnection.QueryRow(context.Background(), "SELECT * FROM album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == pgx.ErrNoRows {
			return alb, fmt.Errorf("albumById %d: no such album", id)
		}

		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}

	return alb, nil
}
