// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build secure

package main

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	_ "github.com/Go-zh/tools/playground"
)

const runUrl = "http://golang.org/compile"

var (
	basePath = flag.String("base", "..", "base path of tour")
	httpAddr = flag.String("http", "127.0.0.1:8083", "HTTP service address (e.g., '127.0.0.1:8083')")
)

func main() {
	flag.Parse()

	http.HandleFunc("/lesson/", lessonHandler)
	http.HandleFunc("/", rootHandler)

	// Keep these static file handlers in sync with ../app.yaml.
	static := http.FileServer(http.Dir(*basePath))
	http.Handle("/content/img/", static)
	http.Handle("/static/", static)
	imgDir := filepath.Join(*basePath, "static", "img")
	http.Handle("/favicon.ico", http.FileServer(http.Dir(imgDir)))

	if err := initTour(*basePath, "HTTPTransport"); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderUI(w); err != nil {
		log.Printf("UI render: %v", err)
	}
}

func lessonHandler(w http.ResponseWriter, r *http.Request) {
	lesson := strings.TrimPrefix(r.URL.Path, "/lesson/")
	if err := writeLesson(lesson, w); err != nil {
		if err == lessonNotFound {
			http.NotFound(w, r)
		} else {
			log.Printf("tour render: %v", err)
		}
	}
}

// prepContent returns a Reader that produces the content from the given
// Reader, but strips the prefix "#appengine: " from each line. It also drops
// any non-blank like that follows a series of 1 or more lines with the prefix.
func prepContent(in io.Reader) io.Reader {
	var prefix = []byte("#appengine: ")
	out, w := io.Pipe()
	go func() {
		r := bufio.NewReader(in)
		drop := false
		for {
			b, err := r.ReadBytes('\n')
			if err != nil && err != io.EOF {
				w.CloseWithError(err)
				return
			}
			if bytes.HasPrefix(b, prefix) {
				b = b[len(prefix):]
				drop = true
			} else if drop {
				if len(b) > 1 {
					b = nil
				}
				drop = false
			}
			if len(b) > 0 {
				w.Write(b)
			}
			if err == io.EOF {
				w.Close()
				return
			}
		}
	}()
	return out
}

// socketAddr returns the WebSocket handler address.
// The App Engine version does not provide a WebSocket handler.
func socketAddr() string { return "" }
