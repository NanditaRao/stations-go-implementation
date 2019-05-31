package main

import (
    "net/http"
    "os"
    "fmt"
    "time"

    "github.com/gorilla/mux"
    "github.com/victorspringer/http-cache"
    "github.com/victorspringer/http-cache/adapter/memory"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    //Implemented caching in memory, the cache will be cleared in 30 mins
    memcached, err := memory.NewAdapter(
        memory.AdapterWithAlgorithm(memory.LRU),
        memory.AdapterWithCapacity(10000000),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    cacheClient, err := cache.NewClient(
        cache.ClientWithAdapter(memcached),
        cache.ClientWithTTL(30 * time.Minute),
        cache.ClientWithRefreshKey("opn"),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for _, route := range routes {
        var handler http.Handler
        handler = cacheClient.Middleware(route.HandlerFunc)
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}
