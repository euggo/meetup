package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = db.Exec(`
	CREATE TABLE info (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		day VARCHAR(64) NULL,
		god VARCHAR(64) NULL
	);
	`)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = db.Exec(`
	INSERT INTO info
	(day, god)
	VALUES
	("Sunday", "Helios"),
	("Monday", "Selene"),
	("Tuesday", "Ares"),
	("Wednesday", "Hermes"),
	("Thursday", "Zeus"),
	("Friday", "Aphrodite"),
	("Saturday", "Kronos")
	`)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cleanup() {
	_ = os.Remove("./example.db")
}

// START OMIT
func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanup()

	rows, err := db.Query("SELECT * FROM info")
	if err != nil {
		log.Panicln(err)
	}

	for rows.Next() {
		id, day, god := 0, "", ""
		if err = rows.Scan(&id, &day, &god); err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%d: %s corresponds to %s\n", id, day, god)
	}
}

// END OMIT
