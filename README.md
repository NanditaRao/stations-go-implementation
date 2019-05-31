# Go implementation of apis using CitiBike's Stations api
A RESTful API example for simple todo application with Go


## Installation & Run
```bash
# Build and Run
cd stations-go-implementation
go build
./stations-go-implementation

# API Endpoint : http://127.0.0.1:4000
```

## API

#### /stations
* `GET` : Gets all Stations, can be queried by page number,


#### /stations/in-service
* `GET` : Gets all stations that are in service. Can be queried by page number.

#### /stations/not-in-service
* `GET` : Gets all stations that are not in service. Can be queried by page number.

#### /stations/:searchString
* `GET` : Gets all the stations that have either the name or street address that contain the search string. The search is case-insensitive

#### /dockable/:stationId/:bikesToReturn
* `GET` : Returns if there are sufficient available docks at a station with an appropriate message.


## Logging and Error Handling

*Used "log" library to implement logging.
*Everytime a request is made, the endpoint, request parameters, path parameters and the time taken to complete the request is logged.
*When an error occurs , the error is logged and error message is returned to user with appropriate http status.

Todo:
*Implement Error Handling in a clean manner. Found [this](https://blog.golang.org/error-handling-and-go)  interesting, but because of lack of familiarity with Go could not implement it.

##Cache implementation

Used "github.com/victorspringer/http-cache/adapter/memory" to implement cache.
*The data is cached by method.
*The cache eviction policy is LRU.

Todo:
*Understand how the implementation can be made more advanced, to be able to plug into Redis.
*Cache by stationId instead of by method name.


##Testing

*The tests that have been added are not unit tests because they are making a call to the actual endpoint.
*I have tested all workflows in the 5 api endpoints.
*Error messages and codes have been tested as well.

Todo:
*Understand how to stub api outputs so that true unit tests can be written
