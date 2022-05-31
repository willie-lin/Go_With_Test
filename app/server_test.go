package app

import "net/http"

func PlayerServer(w http.ResponseWriter, r *http.Request) {

}

func ListenAndServe(addr string, handler Handler) error
