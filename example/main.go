package main

import (
	"net/http"

	"github.com/deweppro/go-badges"
)

func main() {
	b, err := badges.New()
	if err != nil {
		panic(err)
	}

	http.Handle("/image.svg", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		content := r.URL.Query().Get("data")

		err = b.WriteResponse(w, badges.ColorInfo, title, content)
		if err != nil {
			panic(err)
		}
	}))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write([]byte("<html><body><img src=\"/image.svg?title=User ID&data=12 34 567890\"></body></html>"))
		if err != nil {
			panic(err)
		}
	}))

	if err = http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
