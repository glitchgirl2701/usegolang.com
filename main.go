package main

import (
  "net/http"

  "usegolang.com/controllers"

  "github.com/gorilla/mux"
)

func main() {
  staticC := controllers.NewStatic()
  usersC := controllers.NewUsers()

  r := mux.NewRouter ()
  r.Handle("/", staticC.Home).Methods("GET")
  r.Handle("/friends", staticC.Friends).Methods("GET")
  r.HandleFunc("/signup", usersC.New).Methods("GET")
  r.HandleFunc("/signup", usersC.Create).Methods("POST")
  http.ListenAndServe(":3000", r)
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}
