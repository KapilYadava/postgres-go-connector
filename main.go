package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "10.96.78.123" /* Ip address of container (bridge network which enable communication
						 between docker containers) CMD Used - docker inspect container-name | grep IPAddress */
  port     = 1443
  user     = "postgres" // by default username
  password = "mysecretpassword" // by default password
  dbname   = "mytestdb" // CMD used: psql -U postgres (postgres is username here)
)

func main() {
  //time.Sleep(30 * time.Second)
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
	fmt.Println("Not connected!")
	panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
	fmt.Println("Not connected!")
    panic(err)
  }
  fmt.Println("Successfully connected!")
}