package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xdhuxc/xdhuxc-cache/cache"
	"github.com/xdhuxc/xdhuxc-cache/http"
	"os"
)

func init() {
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})
}

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
