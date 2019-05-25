package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RigisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/usr", CreateUser)
	return router
}
func main() {
	r := RigisterHandler()
	http.ListenAndServe("8000", r)
}
