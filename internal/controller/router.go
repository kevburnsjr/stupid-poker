package controller

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/kevburnsjr/stupid-poker/internal/poker"
)

func NewRouter(logger *logrus.Logger, cache poker.GameCache) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", index{logger, cache})
	router.NotFoundHandler = &static{"static"}

	return router
}
