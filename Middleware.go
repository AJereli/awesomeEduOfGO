package main

import (
	"awesomeProject/Auth"
	"net/http"
	_"time"
)
func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}
func AuthToken (inner http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jsonToken Auth.JSONToken
		if err := UnmarshalRequest(r, &jsonToken); err != nil{
			unprocessableEntityApiErr.send(w)
			return
		}

		token := Auth.ParseToken(jsonToken.AccessToken)

		if !token.CheckExpTime() {
			tokenTimeOutApiErr.send(w)
		}
		inner.ServeHTTP(w, r)
	})
}

//func WraperLogger (inner http.HandlerFunc) http.HandlerFunc{
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//
//		inner.ServeHTTP(w, r)
//
//		log.Info(
//			"%s\t%s\t%s\t%s",
//			r.Method,
//			r.RequestURI,
//
//			time.Since(start),
//		)
//	})
//}