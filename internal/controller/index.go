package controller

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"time"
	"html/template"
	"bytes"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/kevburnsjr/stupid-poker/internal/config"
	"github.com/kevburnsjr/stupid-poker/internal/service"
)

type index struct {
	cfg      *config.Api
	log      *logrus.Logger
	gameCache service.GameCache
}

func (c index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var hash string
	cookie, err := r.Cookie("gameHash")
	if err != nil {
		hash = randStr(32)
		http.SetCookie(w, &http.Cookie{
			Name:    "gameHash",
			Value:   hash,
			Expires: time.Now().AddDate(0, 0, 30),
		})
	} else {
		hash = cookie.Value
	}

	var page = "start"
	var hand []string
	var res string
	var balance int

	game := c.gameCache.Get(hash)

	deal := func() {
		page = "deal"
		hand, err = game.Deal()
		if err != nil {
			page = "broke"
		}
		balance = game.GetBalance()
	}

	if r.Method == "GET" {
		if game == nil {
			page = "start"
		} else {
			deal()
		}
	} else {
		switch r.FormValue("action"){
		case "start":
			if game == nil || game.GetBalance() < 1 {
				game = service.NewGame(5, 200)
				c.gameCache.Set(hash, game)
			}
			deal()
		case "exchange":
			if game == nil {
				page = "start"
			} else {
				page = "exchange"
				r.ParseForm()
				cards := r.Form["cards"]
				idx := []int{}
				for _, a := range cards {
					i, err := strconv.Atoi(a)
					if err == nil {
						idx = append(idx, i)
					}
				}
				c.log.Debug("Exchanging cards ", idx)
				hand, res, balance = game.Exchange(idx)
			}
		}
	}

	pageTpl, err := template.ParseFiles("template/"+page+".html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	b1 := &bytes.Buffer{}
	err = pageTpl.Execute(b1, struct{
		Hand []string
		Result string
		Balance int
	}{hand, res, balance})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	layoutTpl, err := template.ParseFiles("template/layout.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	b := &bytes.Buffer{}
	err = layoutTpl.Execute(b, struct{Page template.HTML}{template.HTML(string(b1.Bytes()))})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	w.Write(b.Bytes())
	return
}

func randStr(len int) string {
    buff := make([]byte, len)
    rand.Read(buff)
    str := base64.StdEncoding.EncodeToString(buff)
    return str[:len]
}

