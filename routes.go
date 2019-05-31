package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetStations",
        "GET",
        "/stations",
        GetStations,
    },
    Route{
        "GetStationsInService",
        "GET",
        "/stations/in-service",
        GetStationsInService,
    },
    Route{
        "GetStationsNotInService",
        "GET",
        "/stations/not-in-service",
        GetStationsNotInService,
    },
    Route{
        "GetStationsMatchingString",
        "GET",
        "/stations/{searchString}",
        GetStationsMatchingString,
    },
    Route{
        "CheckIfBikesCanBeDocked",
        "GET",
        "/dockable/{stationId}/{bikesToReturn}",
        CheckIfBikesCanBeDocked,
    },
}
