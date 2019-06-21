package controller

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/kevburnsjr/stupid-poker/internal/config"
)

func NewRouter(cfg *config.Api, logger *logrus.Logger) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/", index{cfg, logger})
	router.NotFoundHandler = &static{"static"}

	return router
}
