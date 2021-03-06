package main

import geolcation_service "github.com/mohamedveron/geolocation-api/geolcation-service"

func main() {

	ip := "82.239.152.199"
	geolocationServiceClient := geolcation_service.NewClient("http://localhost:8080")
	geolocationServiceClient.GetLocationsByIP(ip)


}
