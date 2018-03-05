package main

import (
	_"awesomeProject/MiddelewareChain"
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct{
	Name, Method, Pattern string
	HandlerFunc http.Handler

}

type middleware func(http.Handler) http.Handler

//func buildChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
//	// if our chain is done, use the original handlerfunc
//	if len(m) == 0 {
//		return f
//	}
//	// otherwise nest the handlerfuncs
//	return m[0](buildChain(f, m[1:cap(m)]...))
//}


type Routes []Route

type HandlFunc func(w http.ResponseWriter, r *http.Request)

func getH (hf HandlFunc) http.Handler{
	return http.HandlerFunc(hf)
}

var routes = Routes{
	Route{
		"Welcom",
		"GET",
		"/",
		getH(welcome),

	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		getH(TodoCreate),

	},

	Route{
		"LogIn",
		"GET",
		"/login",
		getH(Login),

	},
	Route{
		"QueryTest",
		"GET",
		"/querytest",
		getH(QueryTest),

	},
	Route{
		"Registration",
		"GET",
		"/registration",
		getH(Registration),
	},
	Route{
		"SendMessage",
		"POST",
		"/massage/send",
		getH(SendMassage),
	},
	Route{
		"GetMassagesToUser",
		"GET",
		"/massage/getmassagestouser",
		getH(GetMassagesToUser),
	},
	Route{
		"GetMassagesFromUser",
		"GET",
		"/massage/getmassagesfromuser",
		getH(GetMassagesFromUser),
	},
}


func InitRouter () *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		var handler http.Handler
		handler = route.HandlerFunc
		handler = WraperLogger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
				Name(route.Name).
					Handler(handler)

	}
	return  router
}
