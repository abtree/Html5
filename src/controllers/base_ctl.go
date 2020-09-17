package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type BaseController struct {
}

func (*BaseController) Init(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../views/main.html")
	t.Execute(w, nil)
}

func (*BaseController) AjAxBase(w http.ResponseWriter, r *http.Request) {
	buff := bytes.NewBuffer([]byte("Hello "))
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	buff.WriteString(fmt.Sprintf("%v: fname: %v lname: %v", r.Method, fname, lname))
	w.Write(buff.Bytes())
}

var BaseCtr = &BaseController{}
