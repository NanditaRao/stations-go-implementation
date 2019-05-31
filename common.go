package main

import (
    "log"
    "strconv"
    "time"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON() (*StationAPIResponse, *ErrorResponse) {
  res, err := http.Get("https://www.citibikenyc.com/stations/json")
  if err != nil {
      return nil,&ErrorResponse{Status: http.StatusInternalServerError,Message:"Internal Error"}
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
      return nil,&ErrorResponse{Status: http.StatusInternalServerError,Message:"Internal Error"}
  }

  s, err := getStations([]byte(body))
  if err != nil {
      return nil,&ErrorResponse{Status: http.StatusInternalServerError,Message:"Internal Error"}
  }
  return s,nil;
}

func getStations(body []byte) (*StationAPIResponse, error) {
    var s = new(StationAPIResponse)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        return nil,err
    }
    return s, nil
}

func getPage(r *http.Request) (int, *ErrorResponse) {
  pages, ok := r.URL.Query()["page"]
   if !ok || len(pages[0]) < 1 {
       return 0,nil;
   }
  page, err := strconv.Atoi(pages[0])
  if(err != nil){
      return 0,&ErrorResponse{Status: http.StatusInternalServerError,Message:"Internal Error"}
  }
  return page,nil
}

func SendResponse(w http.ResponseWriter, res interface{}){
  if err := json.NewEncoder(w).Encode(res); err != nil {
    panic(err)
  }
}

func GetResultForPage(r *http.Request, stationResponses []StationResponse) ([]StationResponse,*ErrorResponse){
  page,err := getPage(r)
  if(err != nil){
      return nil,&ErrorResponse{Status: http.StatusBadRequest,Message:"Invalid page value"}
  }
  if (page != 1 ){
      if((page % 20 == 0 && page > len(stationResponses) / 20) || (page % 20 != 0 && page > (len(stationResponses) / 20)+1)){
        return nil,&ErrorResponse{Status: http.StatusUnprocessableEntity,Message:"Page value exceeds limit"}
    }
  }
  if(page > 0){
      minIndex := (page - 1)*20
      stationResponses = stationResponses[minIndex:minIndex+20]
  }
  return stationResponses,nil
}

func handleError (w http.ResponseWriter, err *ErrorResponse){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(err.Status)
  log.Println (err.Status," ",err.Message)
  SendResponse(w,err.Message)
}

//This is method to read parameters for the dockable call
func GetParameters (r *http.Request)(int64, int64, *ErrorResponse){
  vars := mux.Vars(r)
  stationID, err := strconv.ParseInt(vars["stationId"],10,64)
  if(err != nil){
      return 0,0,&ErrorResponse{Status: http.StatusInternalServerError,Message:"Internal Error"}
  }
  bikesToReturn,err := strconv.ParseInt(vars["bikesToReturn"],10,64)
  if(err != nil){
      return 0,0, &ErrorResponse{Status: http.StatusInternalServerError,Message:"Internal Error"}
  }
  return stationID, bikesToReturn, nil
}
