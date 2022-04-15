package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	albm, er := albumsByArtist("John Coltrane")
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(albm)

	firstData, ferr := albumByID(1)
	if ferr != nil {
		fmt.Println(ferr)
		return
	}
	fmt.Println(firstData)

	arijit := Album{
		Title:  "tere liye",
		Artist: "Arijit Singh",
		Price:  87.9,
	}
	i, aerr := addAlbum(arijit)
	if aerr != nil {
		fmt.Println(aerr, "error adding data")
		return
	}
	fmt.Println(i, "written in db")
}
