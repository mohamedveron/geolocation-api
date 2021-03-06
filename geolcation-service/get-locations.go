package geolcation_service

import (
	"context"
	"fmt"
	geolocation_service_interface "github.com/mohamedveron/geolocation-service/client"
	"net/http"
)

func (c *Client) GetLocationsByIP(ip_address string) (geolocation_service_interface.GeoLocationResponseData, error) {

	payload := geolocation_service_interface.GetLocationsJSONRequestBody{}
	payload.IpAddress = &ip_address

	ctx := context.Background()

	resp, err := c.GeolocationServiceInterface.GetLocations(ctx, payload)

	if err != nil {
		fmt.Print("failed to get locations", err)
	} else if resp.StatusCode != http.StatusOK {
		fmt.Print("failed to get locations", err)
	}

	parseRes, err := geolocation_service_interface.ParseGetLocationsResponse(resp)

	fmt.Print(parseRes.JSON200.Locations)

	return *parseRes.JSON200, nil
}