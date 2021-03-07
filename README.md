# geolocation-api
This component is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to generate apis interface and using go mod for dependancies management.

# Use api/api.yml file in the swagger ui to see the description of the rest api design



## Setup of the service

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, build, run
you can run all of them 

```bash
make all
```

# Start the http server on port 8080:

```bash
make run
```

# Build and run docker image

```bash
docker build --tag geolocation-api .
```

```bash
docker run -p 8090:8080 geolocation-api
```
