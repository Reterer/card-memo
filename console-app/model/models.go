package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(dataSourceName string) error {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}

	// Check tables
	{
		sqlQueue := `
			SELECT name
			FROM sqlite_schema
			WHERE 
				type ='table' AND 
				name NOT LIKE 'sqlite_%'`

		var tables []string
		rows, err := db.Query(sqlQueue)
		if err != nil {
			return err
		}
		for rows.Next() {
			var row string
			err = rows.Scan(&row)
			if err != nil {
				return err
			}
			tables = append(tables, row)
		}
		rows.Close()
		if len(tables) == 0 {
			return CreateTables()
		}
	}

	return db.Ping()
}

func Deinit() {
	if db != nil {
		db.Close()
	}
}

func CreateTables() error {
	tableGroups := `
		CREATE TABLE groups (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			short_desc TEXT,
			full_desc TEXT
		)`

	tableCards := `
		CREATE TABLE cards (
			id INTEGER PRIMARY KEY,
			group_id INTEGER FOREGIN KEY REFERENCES groups (id)
								ON DELETE CASCADE,
			title TEXT NOT NULL,
			short_desc TEXT,
			full_desc TEXT
		)`

	var err error
	_, err = db.Exec(tableGroups)
	if err != nil {
		return err
	}
	_, err = db.Exec(tableCards)
	if err != nil {
		return err
	}

	return nil
}
