package main

import (
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "homestead:secret@tcp(localhost:33060)/cursogo") 
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

  stmt.Exec("Uóxiton istive", 1)
  stmt.Exec("Sheristone Valeska", 2)
}
