package controllers

import (
  "usegolang.com/views"
  "fmt"
  "net/http"
  "usegolang.com/models"
)

func NewUsers(us models.UserService) *Users {
  return &Users {
    NewView:      views.NewView("bootstrap", "views/users/new.gohtml"),
    LoginView:    views.NewView("bootstrap", "views/users/login.gohtml"),
    UserService:  us,
  }
}

type Users struct {
  NewView     *views.View
  LoginView   *views.View
  models.UserService
}

type SignupForm struct {
  Name      string  `schema:"name"`
  Email     string `schema:"email"`
  Password  string `schema:"password"`
}

type LoginForm struct {
  Email    string `schema:"email"`
  Password string `schema:"password"`
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
  form := LoginForm{}
  if err := parseForm(r, &form); err != nil {
    panic(err)
  }

  user, err := u.UserService.Authenticate(form.Email, form.Password)
  switch err {
  case models.ErrNotFound:
    fmt.Fprintln(w,  "Invalid email address.")
  case models.ErrInvalidPassword
    fmt.Fprintln(w, "Invalid password provided.")
  case nil:
    fmt.Fprintln(w, user)
  default:
    http.Error(w, err.Error(), http.StatusInterna;StatusInternalServerError)
  }
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
  if err := u.NewView.Render(w, nil); err != nil {
    panic(err)
  }
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
  form := SignupForm{}
  if err := parseForm(r, &form); err != nil {
    panic(err)
  }
  user := models.User{
    Name:       form.Name,
    Email:      form.Email,
    Password:   form.Password,
  }
  if err := u.UserService.Create(user); err != nil {
    panic(err)
  }
  fmt.Fprintln(w, user)
}
