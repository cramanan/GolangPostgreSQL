package api

import "net/http"

type API struct {
	http.Server
}

func New(addr string) *API {
	server := new(API)
	server.Addr = addr
	return server
}
