package geolcation_service

import geolocation_service_interface "github.com/mohamedveron/geolocation-service/client"

type Client struct{

	GeolocationServiceInterface geolocation_service_interface.ClientInterface
}

func NewClient(geolocationServiceUrl string) *Client {
	client , _ := geolocation_service_interface.NewClient(geolocationServiceUrl)

	return &Client{
		GeolocationServiceInterface: client,
	}
}
