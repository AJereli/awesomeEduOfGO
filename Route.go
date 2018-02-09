package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct{
	Name, Method, Pattern string
	HandlerFun http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Welcom",
		"GET",
		"/",
		welcome,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},

	Route{
		"LogIn",
		"POST",
		"/login",
		Login,
	},
	Route{
		"QueryTest",
		"GET",
		"/querytest",
		QueryTest,
	},
	Route{
		"Registration",
		"GET",
		"/registration",
		Registration,
	},
}


func InitRouter () *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes{

		var handler http.Handler
		handler = route.HandlerFun
		handler = WraperLogger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
				Name(route.Name).
					Handler(handler)

	}
	return  router
}
