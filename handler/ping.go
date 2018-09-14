package handler

import (
	"github.com/jensoncs/bloge/logger"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("serving requet")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"pong\"}"))
}
