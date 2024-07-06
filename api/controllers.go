package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandlerFuncE func(http.ResponseWriter, *http.Request) error

func WriteJSON(writer http.ResponseWriter, statusCode int, v any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(v)
}

func HandleFunc(fn HandlerFuncE) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := fn(writer, request); err != nil {
			log.Println(err)
			WriteJSON(writer, http.StatusInternalServerError, err.Error())
		}
	}
}

func (server *API) CreateUser(writer http.ResponseWriter, request *http.Request) (err error) {
	user := new(User)
	err = json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		return err
	}
	return server.storage.CreateUser(*user)
}

func (server *API) ReadUsers(writer http.ResponseWriter, request *http.Request) (err error) {
	users, err := server.storage.ReadUsers()
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, users)
}

// func (server *API) UpdateUser(writer http.ResponseWriter, request *http.Request) (err error) {
// 	user := new(User)
// 	err = json.NewDecoder(request.Body).Decode(user)
// 	if err != nil {
// 		return err
// 	}
// 	return server.storage.UpdateUser(*user)
// }

func (server *API) DeleteUser(writer http.ResponseWriter, request *http.Request) (err error) {
	user := new(User)
	err = json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		return err
	}
	return server.storage.DeleteUser(*user)
}
