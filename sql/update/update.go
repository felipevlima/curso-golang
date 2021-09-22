package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "felipe:mysql/gocurso")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uoxiton Isvele", 1)
	stmt.Exec("Sheiriston Amani", 2)

	// Delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
