package main

import (
    "net/http"
    "strings"
    "github.com/gorilla/mux"
)

//This method returns true if there are available docks and false if there aren't
func CheckIfBikesCanBeDocked(w http.ResponseWriter, r *http.Request) {
  s,err := getJSON()
  if(err != nil){
    handleError(w,err)
    return
  }
  stationID, bikesToReturn, err := GetParameters(r)
  if(err != nil){
    handleError(w,err)
    return
  }
  stationsMap := make(map[int64]int64)

  for _,station := range s.StationBeanList {
    stationsMap[station.Id] = station.AvailableDocks
  }
  var response DockBikesResponse
  _, ok := stationsMap[stationID]
   if(!ok)  {
    handleError(w, &ErrorResponse{Status: http.StatusUnprocessableEntity,Message:"StationId does not exist"})
    return
  }
  if(stationsMap[stationID] >= bikesToReturn){
    response = DockBikesResponse{Dockable: true, Message:"Suffecient docks available"}
  } else {
    response = DockBikesResponse{Dockable: false, Message:"Insuffecient docks available"}
  }
  SendResponse(w,response)
}

//This method returns any stations whose name or address contain the given search string
func GetStationsMatchingString(w http.ResponseWriter, r *http.Request) {
  s,err := getJSON()
  if(err != nil){
    handleError(w,err)
    return
  }
  vars := mux.Vars(r)
  searchString := vars["searchString"]
  var stationResponses = make([]StationResponse,0)
  for _,station := range s.StationBeanList {
    var search, name, address = strings.ToLower(searchString),strings.ToLower(station.StationName), strings.ToLower(station.StAddress1)
    if(strings.Contains(name,search) || strings.Contains(address,search)){
      var stationRes = StationResponse{Name : station.StationName, Address: station.address(), AvailableDocks : station.AvailableDocks, TotalDocks : station.TotalDocks}
      stationResponses = append(stationResponses, stationRes)
    }
  }
  SendResponse(w,stationResponses)
}

//This method returns all stations that are not in service; query by page supported
func GetStationsNotInService(w http.ResponseWriter, r *http.Request) {
  s,err := getJSON()
  if(err != nil){
    handleError(w,err)
    return
  }
  var stationResponses = make([]StationResponse,0)
  for _,station := range s.StationBeanList {
    if(station.StatusKey == 3){
      var stationRes = StationResponse{Name : station.StationName, Address: station.address(), AvailableDocks : station.AvailableDocks, TotalDocks : station.TotalDocks}
      stationResponses = append(stationResponses, stationRes)
    }
  }
  stationResponses,err = GetResultForPage(r,stationResponses)
  if(err != nil){
    handleError(w,err)
    return
  }
  SendResponse(w,stationResponses)
}

//This method returns all stations that are in service; query by page supported
func GetStationsInService(w http.ResponseWriter, r *http.Request) {
  s,err := getJSON()
  if(err != nil){
    handleError(w,err)
    return
  }
  var stationResponses = make([]StationResponse,0)
  for _,station := range s.StationBeanList {
    if(station.StatusKey == 1){
      var stationRes = StationResponse{Name : station.StationName, Address: station.address(), AvailableDocks : station.AvailableDocks, TotalDocks : station.TotalDocks}
      stationResponses = append(stationResponses, stationRes)
    }
  }
  stationResponses,err = GetResultForPage(r,stationResponses)
  if(err != nil){
    handleError(w,err)
    return
  }
  SendResponse(w,stationResponses)
}

//This method returns all the stations; query by paging supported
func GetStations(w http.ResponseWriter, r *http.Request) {
  s,err := getJSON()
  if(err != nil){
    handleError(w,err)
    return
  }
  var stationResponses = make([]StationResponse,0)
  for _,station := range s.StationBeanList {
    var stationRes = StationResponse{Name : station.StationName, Address : station.address(), AvailableDocks : station.AvailableDocks, TotalDocks : station.TotalDocks}
    stationResponses = append(stationResponses, stationRes)
  }
  stationResponses,err = GetResultForPage(r,stationResponses)
  if(err != nil){
    handleError(w,err)
    return
  }
  SendResponse(w,stationResponses)
}
