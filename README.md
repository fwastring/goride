# goride

Hobby implementation of a carpooling service

## Tech

__Backend__: Golang

__Frontend__: Vue.js

__DB__: PostGIS

__Routing__: OSRM

## Requirements

<!-- Placename API - to get a queryable list of placenames we use [pelias](https://github.com/pelias/pelias) hosted on our own machine. -->
<!-- The enpoint is at: -->
<!-- `http://<host-ip>:4000/v1/autocomplete` -->

Prepare the OSRM backend:

```bash
wget http://download.geofabrik.de/europe/sweden-latest.osm.pbf

docker run -t -v "$(pwd):/data" ghcr.io/project-osrm/osrm-backend osrm-extract -p /opt/car.lua /data/sweden-latest.osm.pbf || echo "osrm-extract failed"

docker run -t -v "$(pwd):/data" ghcr.io/project-osrm/osrm-backend osrm-partition /data/sweden-latest.osrm || echo "osrm-partition failed"
docker run -t -v "$(pwd):/data" ghcr.io/project-osrm/osrm-backend osrm-customize /data/sweden-latest.osrm || echo "osrm-customize failed"
```

## Development

Backend:
```bash
cd api
nix-shell
make dev
```

Frontend:
```bash
cd frontend-vue
npm run dev
```

DB and Routing:
```bash
cd <project-root>
docker compose up -d --build
```
