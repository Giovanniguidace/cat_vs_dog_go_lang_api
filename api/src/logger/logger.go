package logger

import (
	"log"
	"net/http"
	"time"
)

func Logger(r *http.Request, statusCode int){
	log.Printf("\n [%s] %s %s %s %d", time.Now(), r.Method, r.RequestURI, r.Host, statusCode)
}