package controllers

import (
	"html/template"
	"net/http"
)

type BaseController struct {
}

func (*BaseController) Init(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../views/main.html")
	t.Execute(w, nil)
}

var BaseCtr = &BaseController{}
