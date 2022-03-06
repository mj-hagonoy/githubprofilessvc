package server

import (
	"log"
	"net/http"
	"time"
)

func (s *Server) withExecutionLog(handlerfunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("INFO: request received for uri=%s", r.RequestURI)
		defer log.Printf("INFO: request complete for uri=%s, execution time=%v", r.RequestURI, time.Since(start))

		handlerfunc(w, r)
	}
}
