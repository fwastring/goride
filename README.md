# Carpool App

Mock implementation of our carpool app - but in Golang with HTML and PostGIS

## Usage

```bash
docker compose up -d --build # Starts PostGIS DB

go run main.go handlers.go method.go # Starts API

firefox index.html & # Starts frontend
```
