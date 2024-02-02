package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/xid"
)

func NewStore() (*pgxpool.Pool, func(), error) {
	urlOfDb := "postgres://postgres:Postgres@localhost:5432/postgres"
	dbpool, err := pgxpool.New(context.Background(), urlOfDb)
	if err != nil {
		return nil, nil, fmt.Errorf("pgxpool.New errored)) %w", err)
	}

	var greeting string
	if err = dbpool.QueryRow(context.Background(), "select 'postgres'").Scan(&greeting); err != nil {
		return nil, nil, fmt.Errorf("dbpool.QueryRow errored)) %w", err)
	}

	return dbpool, dbpool.Close, err

}

func CreateTable(pool *pgxpool.Pool) error {
	var table = `
        CREATE TABLE IF NOT EXISTS NOTES(
			id		varchar(20) PRIMARY KEY,
			title    varchar(80) NOT NULL,         
			content  text NOT NULL
        )`
	dbres, err := pool.Exec(context.Background(), table)
	fmt.Printf("value of db %v", dbres)
	return err
}

func GetAllNotes(pool *pgxpool.Pool) []Note {
	rows, err := pool.Query(context.Background(), "select * from NOTES")
	if err != nil {
		log.Printf("error getting data from db: %v", err)
	}
	defer rows.Close()

	notes := []Note{}

	for rows.Next() {
		n := Note{}
		err := rows.Scan(&n.ID, &n.Title, &n.Content)
		if err != nil {
			fmt.Println(err)
			continue
		}
		notes = append(notes, n)
	}
	return notes
}

func GetNoteById(pool *pgxpool.Pool, id string) Note {
	row := pool.QueryRow(context.Background(), "select * from NOTES where id = $1", id)
	n := Note{}
	row.Scan(&n.ID, &n.Title, &n.Content)
	return n
}

func AddNote(pool *pgxpool.Pool, note Note) error {
	guid := xid.New()
	id, title, content := guid, note.Title, note.Content

	result, err := pool.Exec(context.Background(), "insert into NOTES (id, title, content) values ($1, $2, $3)", id, title, content)

	if err != nil {
		fmt.Print("insert exec failed: %v", err)
		return err
	}
	fmt.Println(result.RowsAffected())
	return err
}
