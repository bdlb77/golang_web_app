package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// struct with reference to HANDLER
type GzipMiddleware struct {
	Next http.Handler
}

// serve HTTP method for middleware struct
func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if next is nil, set next to be default mux
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}
	// get encoding
	encodings := r.Header.Get("Accept-Encoding")
	// check encoding to see if it supports gzp
	if !strings.Contains(encodings, "gzip") {
		// if it doesn't, return
		gm.Next.ServeHTTP(w, r)
		return
	}
	// if it does
	// add gzip encoding header
	w.Header().Add("Content-Encoding", "gzip")
	gzipWriter := gzip.NewWriter(w)
	// close buffer, send any response back down to client
	defer gzipWriter.Close()
	// grw which takes writer and Response writer
	grw := GzipResponseWriter{
		ResponseWriter: w,
		Writer:         gzipWriter,
	}
	// Next with new writer and pass Request
	gm.Next.ServeHTTP(grw, r)

}

type GzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

// override WRite method to modify for grw
func (grw GzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}
