package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	mydb "github.com/mrtyunjaygr8/pass-n-go/db"
)

func main() {
  fmt.Println("yo")
  dsn := "host=localhost user=mgr8 password=dr0w.Ssap dbname=passngo port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    fmt.Println("in error")
    log.Fatal(err)
  }

  fmt.Println("not in error")
  fmt.Println(db)

  db.AutoMigrate(&mydb.User{})

  u := mydb.User{Username: "test-user", Password: "test-password"}
  create_result := u.CreateUser(db)
  if create_result != nil {
    log.Println(create_result)
  }

  var users []mydb.User
  find_result := db.Find(&users)
  fmt.Println(find_result.RowsAffected)
  fmt.Println(find_result.Error)

  for _, x := range users {
    fmt.Println(x)
  }
}
