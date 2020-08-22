package main

import (
	_ "HtmlBase/routes"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe: 8082 ", err)
	} else {
		log.Println("start server on 8082")
	}
}
