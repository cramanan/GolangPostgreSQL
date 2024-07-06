package api

import (
	"net/http"
)

type API struct {
	http.Server
	storage Storage
}

func New(addr string) (server *API, err error) {
	server = new(API)
	server.Addr = addr
	server.storage, err = NewPostgreSQLStore()
	if err != nil {
		return nil, err
	}
	return server, nil
}
