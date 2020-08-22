package main

import (
	"log"
	"net/http"
	_ "HtmlBase/routes"
)

func main()  {
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe: 8082 ", err)
	} else {
		log.Println("start server on 8082")
	}
}
