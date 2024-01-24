package main

import (
 _ "github.com/go-sql-driver/mysql"
  "fmt"
  "database/sql"
)

func main() {
  db, err := sql.Open("mysql","homestead:secret@tcp(localhost:33060)/cursogo")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
  stmt.Exec("Maria")
  stmt.Exec("João")

  res, _ := stmt.Exec("Pedro")

  id, _ := res.LastInsertId()
  fmt.Println(id)

  linhas, _ := res.RowsAffected()
  fmt.Println(linhas)
}
