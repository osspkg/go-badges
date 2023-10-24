# go-badges

[![Release](https://img.shields.io/github/release/osspkg/go-badges.svg?style=flat-square)](https://github.com/osspkg/go-badges/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/osspkg/go-badges)](https://goreportcard.com/report/github.com/osspkg/go-badges)
[![CI](https://github.com/osspkg/go-badges/actions/workflows/ci.yml/badge.svg)](https://github.com/osspkg/go-badges/actions/workflows/ci.yml)

## Install

```sh
go get -u go.osspkg.com/badges
```

## Colors

```go
ColorPrimary  
ColorSecondary 
ColorSuccess   
ColorDanger   
ColorWarning  
ColorInfo      
ColorLight
```

## Examples

```go
package main

import (
	"net/http"

	"go.osspkg.com/badges"
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

```

## License

BSD-3-Clause License. See the LICENSE file for details.