package app

import (
	"net/http"

	controller "github.com/shakilbd009/go-microsrvcs/mvc/controllers"
)

//StartApp starts the application
func StartApp() {

	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}
