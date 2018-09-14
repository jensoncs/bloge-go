package server

import (
	"github.com/gorilla/mux"
	"github.com/jensoncs/bloge/config"
	"github.com/jensoncs/bloge/handler"
	"github.com/jensoncs/bloge/logger"
	"net/http"
	"strconv"
)

func Startserver() {
	logger.Log.Info("Start bloge")
	appPort := (":" + strconv.Itoa(config.AppPort()))
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	http.ListenAndServe(appPort, router)
}
