package controller

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/kevburnsjr/stupid-poker/internal/config"
	"github.com/kevburnsjr/stupid-poker/internal/service"
)

func NewRouter(cfg *config.Api, logger *logrus.Logger, cache service.GameCache) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", index{cfg, logger, cache})
	router.NotFoundHandler = &static{"static"}

	return router
}
