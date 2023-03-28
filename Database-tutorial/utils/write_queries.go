package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"databases/tutorial/gotypes"
)


// Writes an album
func AddAlbum(dbConnection *pgx.Conn, alb gotypes.Album) (int64, error) {
	var id int64 = 0
	
	if err := dbConnection.QueryRow(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", alb.Title, alb.Artist, alb.Price).Scan(&id); err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}