/*
 * All rights reserved
 * Copyright © Diasoft
 * 2024
 * Address: 3/14, Polkovaya St., Moscow, 127018, Russian Federation
 * Tel.: +7 (495) 780 7575
 * Fax.: +7 (495) 780 7576
 * WEB: http://www.diasoft.com
 * Author: Dmitrii Kondratiev (dkondratiev@diasoft.ru)
 */
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// # url: ${DB_URL:jdbc:postgresql://devops-pg:5432/dkondratiev?currentSchema=ams}
	// # username: ${DB_LOGIN:micro}
	// # password: ${DB_PASSWORD:micro}
	conn, err := pgx.Connect(context.Background(), "postgresql://micro:micro@devops-pg:5432/dkondratiev?search_path=ams")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	// var weight int64
	var artist string
	// err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	err = conn.QueryRow(context.Background(), "select title, artist from album where id=$1", 3).Scan(&name, &artist)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	// fmt.Println(name, weight)
	fmt.Println(name, "=", artist)

	albums, err := albumsByArtist("John Coltrane", conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// albID, err := addAlbum(
	// 	Album{
	// 		Title:  "The Modern Sound of Betty Carter",
	// 		Artist: "Betty Carter",
	// 		Price:  49.99,
	// 	},
	// 	conn,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albID)

}

func albumsByArtist(name string, conn *pgx.Conn) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album
	// rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	rows, err := conn.Query(context.Background(), "SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func addAlbum(alb Album, conn *pgx.Conn) (int64, error) {
	// result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	result, err := conn.Exec(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	//     return 0, fmt.Errorf("addAlbum: %v", err)
	// }
	id := result.RowsAffected()
	return id, nil
}
