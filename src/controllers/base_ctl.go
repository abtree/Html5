package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type SellItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Content   string `json:"content"`
	Kind      int    `json:"kind"`
	Author    string `json:"author"`
	Flag      int    `json:"flag"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	IsDeleted bool   `json:"isDeleted"`
	UserId    int    `json:"userId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	IP        string `json:"ip"`
}

type BaseController struct {
	sellItems []*SellItem //销售的商品
	articles  []*Article  //文章
}

func init() {
	BaseCtr.prepare()
	BaseCtr.init_articles()
}

func (ctr *BaseController) init_articles() {
	now := time.Now().Unix()
	ctr.articles = append(ctr.articles, &Article{
		ID:        1,
		Title:     "Title1",
		Cover:     "assets/images/dog.png",
		Content:   "content",
		Kind:      1,
		Author:    "author",
		Flag:      1,
		CreatedAt: now,
		UpdatedAt: now,
		IsDeleted: false,
		UserId:    1,
		Name:      "name",
		Avatar:    "assets/images/dog.png",
		IP:        "127.0.0.1",
	}, &Article{
		ID:        2,
		Title:     "Title2",
		Cover:     "assets/images/dog.png",
		Content:   "content",
		Kind:      1,
		Author:    "author",
		Flag:      1,
		CreatedAt: now,
		UpdatedAt: now,
		IsDeleted: false,
		UserId:    1,
		Name:      "name",
		Avatar:    "assets/images/dog.png",
		IP:        "127.0.0.1",
	})
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

func (ctr *BaseController) ArticleListRandom(w http.ResponseWriter, r *http.Request) {
	if byts, err := json.Marshal(ctr.articles); err == nil {
		w.Write(byts)
	} else {
		w.Write([]byte(err.Error()))
	}
}

func (ctr *BaseController) UserCaptcha(w http.ResponseWriter, r *http.Request) {
	byts, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	pars := &struct {
		Identify string `json:"identify"`
		Kind     int    `json:"kind"`
	}{}
	err = json.Unmarshal(byts, pars)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(string(byts))
}

func (ctr *BaseController) UserSignin(w http.ResponseWriter, r *http.Request) {
	byts, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	pars := &struct {
		Identify string `json:"identify"`
		Captcha  string `json:"captcha"`
	}{}
	err = json.Unmarshal(byts, pars)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(byts))

	back := &struct {
		UserId    int    `json:"userId"`
		Name      string `json:"name"`
		Avatar    string `json:"avatar"`
		IP        string `json:"ip"`
		Bio       string `json:"bio"`
		Sex       int    `json:"sex"`
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expiredAt"`
	}{
		UserId:    1,
		Name:      "Name1",
		Avatar:    "assets/images/dog.png",
		IP:        "127.0.0.1",
		Bio:       "个性签名",
		Sex:       1,
		Token:     "tokenxxxxxtoken",
		ExpiredAt: time.Now().Add(365 * 24 * time.Hour).Unix(),
	}

	if byts, err := json.Marshal(back); err == nil {
		w.Write(byts)
	} else {
		w.Write([]byte(err.Error()))
	}
}
