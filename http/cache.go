package http

import (
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/xdhuxc/xdhuxc-cache/cache"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())

	http.ListenAndServe(":8080", nil)
}

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.EscapedPath(), "/")[2]

	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	method := r.Method
	if method == http.MethodPut {
		b, _ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			err := h.Set(key, b)
			if err != nil {
				log.Errorln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		return
	}

	if method == http.MethodGet {
		bytes, err := h.Get(key)
		if err != nil {
			log.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(bytes) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = w.Write(bytes)
		if err != nil {
			log.Errorln(err)
		}
		return
	}

	if method == http.MethodDelete {
		err := h.Del(key)
		if err != nil {
			log.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}
