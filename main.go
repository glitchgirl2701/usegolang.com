package main

import (
  "net/http"
  "fmt"
  "usegolang.com/controllers"
  "usegolang.com/models"
  "github.com/gorilla/mux"
)

const (
  host =      "localhost"
  port =      5432
  user =      "kaylathomsen"
  password =  "broadway"
  dbname =    "usegolang_dev"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  us, err := models.NewUserService(psqlInfo)
    if err != nil {
      panic(err)
    }
  defer us.Close()
  //us.DestructiveReset()

  staticC := controllers.NewStatic()
  usersC := controllers.NewUsers(us)

  r := mux.NewRouter ()
  r.Handle("/", staticC.Home).Methods("GET")
  r.Handle("/friends", staticC.Friends).Methods("GET")
  r.Handle("/login", usersC.LoginView).Methods("GET")
  r.HandleFunc("/login", usersC.Login).Methods("POST")
  r.HandleFunc("/signup", usersC.New).Methods("GET")
  r.HandleFunc("/signup", usersC.Create).Methods("POST")
  http.ListenAndServe(":3000", r)
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}
