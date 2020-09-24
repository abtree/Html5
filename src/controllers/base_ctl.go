package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type SellItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

type BaseController struct {
	sellItems []*SellItem //销售的商品
}

func init() {
	BaseCtr.prepare()
}

func (ctr *BaseController) prepare() {
	ctr.sellItems = append(ctr.sellItems, &SellItem{
		ID:   1,
		Name: "测试商品1",
		Img:  "/views/imgs/XSLB_LB01.png",
	}, &SellItem{
		ID:   2,
		Name: "测试商品2",
		Img:  "/views/imgs/XSLB_LB03.png",
	}, &SellItem{
		ID:   3,
		Name: "测试商品3",
		Img:  "/views/imgs/XSLB_LB04.png",
	}, &SellItem{
		ID:   4,
		Name: "测试商品4",
		Img:  "/views/imgs/XSLB_LB05.png",
	}, &SellItem{
		ID:   5,
		Name: "测试商品5",
		Img:  "/views/imgs/XSLB_LB06.png",
	}, &SellItem{
		ID:   6,
		Name: "测试商品6",
		Img:  "/views/imgs/XSLB_LB07.png",
	}, &SellItem{
		ID:   7,
		Name: "测试商品7",
		Img:  "/views/imgs/XSLB_LB08.png",
	}, &SellItem{
		ID:   8,
		Name: "测试商品8",
		Img:  "/views/imgs/XSLB_LB09.png",
	}, &SellItem{
		ID:   9,
		Name: "测试商品9",
		Img:  "/views/imgs/XSLB_LB10.png",
	}, &SellItem{
		ID:   10,
		Name: "测试商品10",
		Img:  "/views/imgs/XSLB_LB11.png",
	}, &SellItem{
		ID:   11,
		Name: "测试商品11",
		Img:  "/views/imgs/XSLB_LB12.png",
	}, &SellItem{
		ID:   12,
		Name: "测试商品12",
		Img:  "/views/imgs/XSLB_LB13.png",
	})
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

func (ctr *BaseController) AjAxSellItem(w http.ResponseWriter, r *http.Request) {
	if byts, err := json.Marshal(ctr.sellItems); err == nil {
		w.Write(byts)
	} else {
		w.Write([]byte(err.Error()))
	}
}

var BaseCtr = &BaseController{
	sellItems: make([]*SellItem, 0, 12),
}
