# goride

Mock implementation of our carpool app - but in Golang with HTML and PostGIS

## Usage

Placename API - to get a queryable list of placenames we use [pelias](https://github.com/pelias/pelias) hosted on our own machine.
The enpoint is at:
`http://<host-ip>:4000/v1/autocomplete`

Prepare the OSRM backend:

```bash
wget http://download.geofabrik.de/europe/sweden-latest.osm.pbf

docker run -t -v "$(pwd):/data" ghcr.io/project-osrm/osrm-backend osrm-extract -p /opt/car.lua /data/sweden-latest.osm.pbf || echo "osrm-extract failed"

docker run -t -v "$(pwd):/data" ghcr.io/project-osrm/osrm-backend osrm-partition /data/sweden-latest.osrm || echo "osrm-partition failed"
docker run -t -v "$(pwd):/data" ghcr.io/project-osrm/osrm-backend osrm-customize /data/sweden-latest.osrm || echo "osrm-customize failed"
```

Deploy the app:

```bash
docker compose up -d --build # Starts DB, OSRM backend, go API and frontend
```
