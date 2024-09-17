package logger

import (
	"log"
	"net/http"
	"time"
)

func Logger(r *http.Request, statusCode int){
	log.Printf("\n [%s] %s %s %s %d", time.Now().Format(time.RFC822), r.Method, r.RequestURI, r.Host, statusCode)
}