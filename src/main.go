package main

import (
	_ "HtmlBase/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("start server on 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: 8081 ", err)
	}
}
