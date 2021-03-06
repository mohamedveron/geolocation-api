package api

import geolcation_service "github.com/mohamedveron/geolocation-api/geolcation-service"

type Server struct {
	geolocationServiceClient *geolcation_service.Client
}

func NewServer(geolocationServiceClient *geolcation_service.Client) *Server {
	return &Server{
		geolocationServiceClient: geolocationServiceClient,
	}
}
