package main

import geolocation_service_interface "github.com/mohamedveron/geolocation-service/client"

func main() {

	//ip := "82.239.152.199"
geolocation_service_interface.NewClient("http://localhost:8080")

}
