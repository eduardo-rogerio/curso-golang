package main

import (
  "database/sql"
  "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
  result, err := db.Exec(sql)
  if err != nil {
    panic(err)
  }
  return result
}

func main() {
  cfg := mysql.Config{
    User: "homestead",
    Passwd: "secret",
    Net: "tcp",
    Addr: "127.0.0.1:33060",
  }
  db, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    panic(err)
  }
  defer db.Close()
  exec(db, "create database if not exists cursogo")
}
