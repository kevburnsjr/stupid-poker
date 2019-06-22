package controller

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

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
		c.gameCache.Set(hash, service.NewGame(5, 200))
		http.SetCookie(w, &http.Cookie{
			Name:    "gameHash",
			Value:   hash,
			Expires: time.Now().AddDate(0, 0, 30),
		})
	} else {
		hash = cookie.Value
	}

	game := c.gameCache.Get(hash)
	if game == nil {
		game = service.NewGame(5, 200)
		c.gameCache.Set(hash, game)
	}

	w.WriteHeader(200)
	w.Write([]byte(`<html><body>`))
	w.Write([]byte(`<style>span.hand{font-size: 8em;}</style>`))

	_, err = game.Deal()
	if err != nil {
		game = service.NewGame(5, 200)
		c.gameCache.Set(hash, game)
		game.Deal()
	}

	hand := strings.Join(game.GetHandUtf8(), " ")
	w.Write([]byte(fmt.Sprintf("<p>Hand: <br><span class='hand'>%v</span></p>", hand)))

	_, res, balance := game.Exchange([]int{0, 1, 2})

	hand = strings.Join(game.GetHandUtf8(), " ")
	w.Write([]byte(fmt.Sprintf("<p>Final Hand: <br><span class='hand'>%v</span><br>Result: %v<br>Balance: %v", hand, res, balance)))

	w.Write([]byte(`</body></html>`))

	return
}

func randStr(len int) string {
    buff := make([]byte, len)
    rand.Read(buff)
    str := base64.StdEncoding.EncodeToString(buff)
    return str[:len]
}
