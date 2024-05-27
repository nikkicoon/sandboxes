package main

import (
  "database/sql"
  "fmt"
  "strconv"
  _ "github.com/mattn/go-sqlite3"
)

func main() {
  // 1. open database.
  // sql.Open(): requires a driver, followed by 'what database'
  database, _ := sql.Open("sqlite3", "./nraboy.db")
  // Prepare(): this will be provided to the query statement.
  // takes a SQL query.
  statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
  // execute the statement.
  statement.Exec()
  // another query. parameterized query, "?" being the parameter
  statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
  // provide this exec the 2 parameters.
  statement.Exec("Nic", "Raboy")
  // query the database.
  rows, _ := database.Query("SELECT id, firstname, lastname FROM people")

  // places to store the results from  the query:
  var id int
  var firstname string
  var lastname string
  // loop over the rows
  // for every row that exist, we say rows.Scan(), and
  // we are going to store our columns for that row in our variables.
  // we have to list them in the order we selected them.
  // last we are printing them out.
  for rows.Next() {
    rows.Scan(&id, &firstname, &lastname)
    fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
  }
}
