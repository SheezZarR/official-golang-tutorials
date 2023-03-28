package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v5"

	"databases/tutorial/utils"
	"databases/tutorial/gotypes"
)


var db *pgx.Conn


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
	
	albums, err := utils.AlbumsByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	album, err := utils.AlbumById(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", album)

	albId, err := utils.AddAlbum(db, gotypes.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Newly added album id: %v\n", albId)
}	
