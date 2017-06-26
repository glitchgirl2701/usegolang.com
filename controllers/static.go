package controllers

import ("usegolang.com/views")

func NewStatic() *Static{
  return &Static {
    Home: views.NewView("bootstrap", "views/static/home.gohtml"),
    Friends: views.NewView("bootstrap", "views/static/friends.gohtml")}
}

type Static struct {
  Home    *views.View
  Friends *views.View
}
