package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bytes, err := json.Marshal(h.GetStatus())
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Errorln(err)
	}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
