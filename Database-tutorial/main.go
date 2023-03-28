package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	// "database/sql"
)

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

var db *pgx.Conn

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album
	
    rows, err := db.Query(context.Background(), "SELECT id, artist, price FROM album WHERE artist = $1", name)
	if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
	defer rows.Close()

	for rows.Next() {
		var album Album
		
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


func albumById(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow(context.Background(), "SELECT * FROM album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == pgx.ErrNoRows {
			return alb, fmt.Errorf("albumById %d: no such album", id)
		}

		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}

	return alb, nil
}


func main() {
	fmt.Println("Hello databases!")
	var err error

	db, err = pgx.Connect(context.Background(), "postgres://postgres:1@0.0.0.0:5555/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection successful")
	fmt.Printf("Checking db connection: %p\n", db)
	
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumById(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", album)
}	
