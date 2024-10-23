package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"goride/store"
	"goride/store/dbstore"
	"goride/store/types"

	"github.com/gin-gonic/gin"
)

type JoinRouterHandler struct {
	logger         slog.Logger
	database 	   *sql.DB
	routeStore	   dbstore.RouteStore
}

type JoinRouterHandlerParams struct {
	Logger         slog.Logger
	Database 	   *sql.DB
	RouteStore	   dbstore.RouteStore
}


func NewJoinRouterHandler(params JoinRouterHandlerParams) *JoinRouterHandler {
	return &JoinRouterHandler{
		logger: params.Logger,	
		database: params.Database,
		routeStore: params.RouteStore,
	}
}

func (h *JoinRouterHandler) ServeHTTP(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	 
	var trip struct {
		TripID uint `json:"trip_id"`
		RiderID uint `json:"rider_id"`
		From string `json:"from"`
		To   string `json:"to"`
	}

	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	route, err := h.routeStore.GetRoute(trip.TripID)
	if err != nil {

	}

	points := []types.Point{h.getLocation(route.StartAddress), h.getLocation(trip.From), h.getLocation(trip.To), h.getLocation(route.EndAddress)}

	geometry, err := h.getRouteFromOSRM(points)

	if err != nil {
		h.logger.Error("Failed getting route", err)
	}

	err = h.routeStore.UpdateRoute(trip.TripID, geometry)
	if err != nil {
		h.logger.Error("Error error adding route to db", err)
	}

     
    c.JSON(http.StatusOK, "Joined route successfully")
}

func (h *JoinRouterHandler) getRouteFromOSRM(points []types.Point) (store.Geometry, error) {
	var pointString strings.Builder
	for i, point := range points {
		pointString.WriteString(strconv.FormatFloat(point.Longitude, 'f', -1, 64))
		pointString.WriteString(",")
		pointString.WriteString(strconv.FormatFloat(point.Latitude, 'f', -1, 64))
		if i != len(points)-1 {
			pointString.WriteString(";")
		}
	}
	fmt.Print(pointString.String())
	url := fmt.Sprintf("http: 
	blank := store.Geometry{}

	resp, err := http.Get(url)

	if err != nil {
		return blank, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return blank, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return blank, err
	}

	var osrmResp store.OsrmResponse
	if err := json.Unmarshal(body, &osrmResp); err != nil {
		return blank, err
	}

	for _, coordinate := range osrmResp.Route[0].Geometry.Coordinates {
		temp := coordinate[0]
		coordinate[0] = coordinate[1]
		coordinate[1] = temp
	}

	 
	if len(osrmResp.Route) > 0 {
		return osrmResp.Route[0].Geometry, nil
	}

	return blank, fmt.Errorf("no routes found")
}



func (h *JoinRouterHandler) getLocation(address string) types.Point {
    baseURL := "https: 

    params := url.Values{}
    params.Add("q", address)
    params.Add("format", "json")

    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    resp, err := http.Get(fullURL)
    if err != nil {
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
		h.logger.Error("Error reading body: %s \n", err)
    }

     
    var locations []store.Location
    if err := json.Unmarshal(body, &locations); err != nil {
		h.logger.Error("Error reading body: %s \n", err)
    }

     
	lon, err := strconv.ParseFloat(locations[0].Lon, 64)
	lat, err := strconv.ParseFloat(locations[0].Lat, 64)
	point := types.Point{Latitude: lat, Longitude: lon}
	return point
}

