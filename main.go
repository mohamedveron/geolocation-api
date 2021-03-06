package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/geolocation-api/api"
	geolcation_service "github.com/mohamedveron/geolocation-api/geolcation-service"
	"net/http"
)

func main() {

	geolocationServiceClient := geolcation_service.NewClient("http://localhost:9090")

	server := api.NewServer(geolocationServiceClient)

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}


}
