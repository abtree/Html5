package main

import (
	_ "HtmlBase/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("start server on 8078")
	err := http.ListenAndServe(":8078", nil)
	if err != nil {
		log.Fatal("ListenAndServe: 8078 ", err)
	}
}
