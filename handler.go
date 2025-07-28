package simple_bootstrap

import (
	"fmt"
	"net/http"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func HTTPHandlerFunc(handler HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
