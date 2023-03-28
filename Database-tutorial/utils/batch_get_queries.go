package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"databases/tutorial/gotypes"
)

// albumsByArtist queries for albums that have the specified artist name.
func AlbumsByArtist(dbConnection *pgx.Conn, name string) ([]gotypes.Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []gotypes.Album

	rows, err := dbConnection.Query(context.Background(), "SELECT id, artist, price FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var album gotypes.Album

		if err := rows.Scan(&album.ID, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil

}