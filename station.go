package main


  type StationResponse struct {
    Name string `json:"name"`
    Address string `json:"address"`
    AvailableDocks int64 `json:"availableDocks"`
    TotalDocks int64 `json:"totalDocks"`
  }
  type Station struct {
      Id int64 `json:"id"`
      StationName string `json:"stationName"`
      AvailableDocks int64 `json:"availableDocks"`
      TotalDocks int64 `json:"totalDocks"`
      Latitude float64 `json:"latitude"`
      Longitude float64 `json:"longitude"`
      StatusValue string `json:"statusValue"`
      StatusKey int64 `json:"statusKey"`
      AvailableBikes int64 `json:"availableBikes"`
      StAddress1 string `json:"stAddress1"`
      StAddress2 string `json:"stAddress2"`
      City string `json:"city"`
      PostalCode string `json:"postalCode"`
      Location string `json:"location"`
      Altitude string `json:"altitude"`
      TestStation bool `json:"testStation"`
      LastCommunicationTime string `json:"lastCommunicationTime"`
      LandMark string `json:"landMark"`
   }

    type StationAPIResponse struct {
        ExecutionTime string `json:"executionTime"`
        StationBeanList []Station `json:"stationBeanList"`
   }

   type DockBikesResponse struct {
     Dockable bool `json:"dockable"`
     Message string `json:"message"`
   }

   type ErrorResponse struct {
     Status int
     Message string
   }

   func (s *Station) address() string {
    return s.StAddress1 +" "+ s.StAddress2+" "+s.City+" "+s.PostalCode
  }
