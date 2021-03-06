package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetLocations(w http.ResponseWriter, r *http.Request) {

	var filters GeoLocationRequestData
	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add filters
	ipAddress := "82.239.152.199"
	if filters.IpAddress != nil{
		ipAddress = *filters.IpAddress
	}

	responseList, err := s.geolocationServiceClient.GetLocationsByIP(ipAddress)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseList)
}