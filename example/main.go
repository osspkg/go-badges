/*
 *  Copyright (c) 2022-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// Package main runs the badge HTTP example.
package main

import (
	"errors"
	"net/http"
	"time"

	"go.osspkg.com/badges"
)

const (
	exampleReadHeaderTimeout = 5 * time.Second
	exampleReadTimeout       = 10 * time.Second
	exampleWriteTimeout      = 10 * time.Second
	exampleIdleTimeout       = 60 * time.Second
	exampleMaxHeaderBytes    = 64 << 10
)

func main() {
	b, err := badges.New()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/image.svg", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		content := r.URL.Query().Get("data")

		if err := b.WriteResponse(w, badges.ColorInfo, title, content); err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, badges.ErrFieldTooLong) {
				status = http.StatusBadRequest
			}
			http.Error(w, http.StatusText(status), status)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("<html><body><img src=\"/image.svg?title=User ID&data=12 34 567890\"></body></html>")); err != nil {
			return
		}
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: exampleReadHeaderTimeout,
		ReadTimeout:       exampleReadTimeout,
		WriteTimeout:      exampleWriteTimeout,
		IdleTimeout:       exampleIdleTimeout,
		MaxHeaderBytes:    exampleMaxHeaderBytes,
	}
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
