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
		//设置跨域允许
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		w.Header().Add("Access-Control-Max-Age", "3600")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Access-Token,Authorization,ybg")
		f(w, r)
	}
}

func init() {
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("../views"))))

	http.HandleFunc("/", logPanics(controllers.BaseCtr.Init))
	http.HandleFunc("/ajax/base", logPanics(controllers.BaseCtr.AjAxBase))
	http.HandleFunc("/ajax/sellitem", logPanics(controllers.BaseCtr.AjAxSellItem))
}
