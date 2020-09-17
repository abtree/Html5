package routes

import (
	"HtmlBase/controllers"
	"log"
	"net/http"
)

type HandleFnc func(http.ResponseWriter, *http.Request)

//处理异常的闭包封装函数
func logPanics(f HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		f(w, r)
	}
}

func init() {
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("../views"))))

	http.HandleFunc("/", logPanics(controllers.BaseCtr.Init))
	http.HandleFunc("/ajax/base", logPanics(controllers.BaseCtr.AjAxBase))
}
