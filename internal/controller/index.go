package controller

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/kevburnsjr/stupid-poker/internal/config"
	"github.com/kevburnsjr/stupid-poker/internal/service"
)

type index struct {
	cfg *config.Api
	log *logrus.Logger
}

func (c index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<html><body>`))

	game := service.NewGame(5, 200)
	hand := game.Deal()

	w.Write([]byte(fmt.Sprintf("<p>Hand: %v</p>", hand)))

	hand, res, balance := game.Exchange([]int{0, 1, 2})

	w.Write([]byte(fmt.Sprintf("<p>Final Hand: %v<br>Result: %v<br>Balance: %v", hand, res, balance)))

	w.Write([]byte(`</body></html>`))

	return
}
