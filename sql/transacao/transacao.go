package main

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql","homestead:secret@tcp(localhost:33060)/cursogo")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  tx, _ := db.Begin()
  stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

  stmt.Exec(2000, "Bia")
  stmt.Exec(2001, "Carlos")
  _, err = stmt.Exec(1, "Tiago")

  if err != nil {
    tx.Rollback()
    log.Fatal(err)
  }
}
