package main

import (
  "fmt"
  "log"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

type usuario struct {
  id int
  nome string
}

func main() {
  db, err := sql.Open("mysql", "homestead:secret@tcp(localhost:33060)/cursogo")  
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  
  rows, _ := db.Query("select id, nome from usuarios where id > ?", 3)
  defer rows.Close()

  for rows.Next() {
    var u usuario 
    rows.Scan(&u.id, &u.nome)
    fmt.Println(u)
  }
}
