 

package types

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/paulsmith/gogeos/geos"
	gogeom "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkt"
)


type Geometry4326 struct {
    *geos.Geometry
}

type Point4326 struct {
    *geos.Geometry
}

type Point struct {
	Latitude float64
	Longitude float64
}

func (p Point) ToWKT() string {
	return fmt.Sprintf("POINT(%f %f)", p.Longitude, p.Latitude)
}

func (p Point) ToPoint4326() Point4326 {
	geom, err := geos.FromWKT(p.ToWKT())
	if err != nil {

	}

	return Point4326{Geometry: geom}
}


func (g Geometry4326) MarshalJSON() ([]byte, error) {
     
    str, err := g.Geometry.ToWKT()  
    if err != nil {
        return nil, err
    }

     
    geom, err := wkt.Unmarshal(str)
    if err != nil {
        return nil, err
    }

     
    geoJSON := make(map[string]interface{})
    
    switch v := geom.(type) {
    case *gogeom.Point:
        geoJSON["type"] = "Point"
        geoJSON["coordinates"] = []float64{v.X(), v.Y()}
    case *gogeom.LineString:
        coords := make([][]float64, v.NumCoords())
        for i := 0; i < v.NumCoords(); i++ {
            coords[i] = []float64{v.Coord(i).X(), v.Coord(i).Y()}
        }
        geoJSON["type"] = "LineString"
        geoJSON["coordinates"] = coords
    case *gogeom.Polygon:
        coords := make([][][]float64, v.NumLinearRings())
        for i := 0; i < v.NumLinearRings(); i++ {
            ring := v.LinearRing(i)
            ringCoords := make([][]float64, ring.NumCoords())
            for j := 0; j < ring.NumCoords(); j++ {
                ringCoords[j] = []float64{ring.Coord(j).X(), ring.Coord(j).Y()}
            }
            coords[i] = ringCoords
        }
        geoJSON["type"] = "Polygon"
        geoJSON["coordinates"] = coords
    default:
        return nil, fmt.Errorf("unsupported geometry type: %T", v)
    }

     
    return json.Marshal(geoJSON)
}


func (g *Geometry4326) Value() (driver.Value) {
    geom := g.Geometry
	if geom == nil {
		return nil
	}
	str, err := geom.ToWKT()
    if err != nil {
        return nil
    }
    return "SRID=4326;" + str
}

func (g *Geometry4326) Scan(value interface{}) error {
	if value == nil {
        return nil  
    }

     
    str, ok := value.(string)
    if !ok {
        return fmt.Errorf("unsupported type for geometry: %T", value)
    }

     
    bytes, err := hex.DecodeString(str)
    if err != nil {
        return fmt.Errorf("error decoding string to bytes: %w", err)
    }

     
    geom, err := geos.FromWKB(bytes)
    if err != nil {
        return fmt.Errorf("error converting WKB to geometry: %w", err)
    }

	*g = Geometry4326{geom}
    return nil
}



 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 

