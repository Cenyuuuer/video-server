package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RigisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/usr", CreateUser)       //处理的时候会给handler新建一个协程来处理
	router.POST("/usr/:user_name", Login) //usr/后边的是传入的参数
	return router
}
func main() {
	r := RigisterHandler()
	http.ListenAndServe(":8000", r)
	// fmt.Println("aaaaaaaaa")
}

/**
* handler -> validation{1.request,2.user} ->business logic ->response
1.data model
2.error handling
*/
