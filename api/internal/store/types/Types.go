package types

type OsrmResponse struct {
	Route []Route `json:"routes"`
}

type Route struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
	Type        string      `json:"type"`
}

type Point struct {
	Longitude float64
	Latitude  float64
}

type Location struct {
    DisplayName string  `json:"display_name"`
    Lat         string  `json:"lat"`
    Lon         string  `json:"lon"`
}


type RouteResult struct {
    RouteID int  `json:"route_id"`
    OnRoute bool `json:"on_route"`
}

type AddressCheckRequest struct {
	From string `json:"from"`
	To string `json:"to"`
}

type AddressCheckResponse struct {
	OnRoute bool `json:"on_route"`
	RouteID int `json:"id"`
}

type RoutesResponse struct {
	Routes []Route `json:"routes"`
}
