package main

import (
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s\n",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
        log.Println(
            "Query parameters ",r.URL.Query(),
        )
        log.Println(
            "Path parameters ",mux.Vars(r),
        )
    })
}
