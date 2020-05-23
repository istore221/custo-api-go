package middleware

import (
	"fmt"
	"net/http"
)

type  RootMiddleware struct {
	Next http.Handler
}

func (rootMiddleware RootMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request){

	fmt.Println("is going through the middleware")
	if rootMiddleware.Next != nil {
		rootMiddleware.Next.ServeHTTP(w, r)
	}
}


