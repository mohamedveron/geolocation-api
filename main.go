package main

import (
	"context"
	"fmt"
	geolocation_service_interface "github.com/mohamedveron/geolocation-service/client"
	"net/http"
)

func main() {

	ip := "82.239.152.199"
	client2 , _ := geolocation_service_interface.NewClient("http://localhost:8080")

	payload := geolocation_service_interface.GetLocationsJSONRequestBody{}
	payload.IpAddress = &ip

	ctx := context.Background()

	resp, err := client2.GetLocations(ctx, payload)

	if err != nil {
		fmt.Print("failed to get locations", err)
	} else if resp.StatusCode != http.StatusOK {
		fmt.Print("failed to get locations", err)
	}

	parseRes, err := geolocation_service_interface.ParseGetLocationsResponse(resp)

	fmt.Print(parseRes.JSON200.Locations)

}
