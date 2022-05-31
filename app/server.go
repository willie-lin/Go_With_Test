package app

import "net/http"

func PlayerServer(w http.ResponseWriter, r *http.Request) {

}

type HandlerFunc func(ResponseWriter, *Request)
