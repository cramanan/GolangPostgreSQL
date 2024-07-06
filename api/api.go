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

	router := http.NewServeMux()
	router.HandleFunc("/create", HandleFunc(server.CreateUser))
	router.HandleFunc("/read", HandleFunc(server.ReadUsers))
	// router.HandleFunc("/create", HandleFunc(server.UpdateUser))
	router.HandleFunc("/delete", HandleFunc(server.DeleteUser))
	server.Handler = router

	return server, nil
}
