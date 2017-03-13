package middlewares

import (
	"fmt"
	"net/http"
)

func Logger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("serving: " + r.URL.Path)
	next(w, r)
}
