package main

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("[%s] %s - Remote Address: %v\n", req.Method, req.URL.Path, req.RemoteAddr)
		fmt.Printf("Host: %v\n", req.Host)
		fmt.Printf("User-Agent: %v\n", req.UserAgent())
		next.ServeHTTP(w, req)
	})
}

func HeaderLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("\n=== Request Headers ===")
		for name, values := range req.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}
		fmt.Println("====================")
		next.ServeHTTP(w, req)
	})
}
