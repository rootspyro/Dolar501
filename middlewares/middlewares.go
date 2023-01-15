package middlewares

import (
	"log"
	"net/http"
)

type Middlewares struct {

}

func NewMiddlewares() *Middlewares{
	return &Middlewares{}
}

func( mid *Middlewares )RequestLogger(r *http.Request) {
		log.Printf("%s:%s - %s\n", r.Method, r.URL, r.RemoteAddr)
}
