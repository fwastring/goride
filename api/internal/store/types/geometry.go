// db/types/geometry.go

package types

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/paulsmith/gogeos/geos"
	"github.com/twpayne/go-geom/encoding/wkt"
	gogeom "github.com/twpayne/go-geom"
)


type Geometry4326 struct {
    *geos.Geometry
}

func (g Geometry4326) MarshalJSON() ([]byte, error) {
    // Convert the geometry to WKT string
    str, err := g.Geometry.ToWKT() // Assuming this gets WKT from your Geometry type
    if err != nil {
        return nil, err
    }

    // Parse WKT into a go-geom Geometry
    geom, err := wkt.Unmarshal(str)
    if err != nil {
        return nil, err
    }

    // Define the GeoJSON structure
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

    // Marshal the GeoJSON map into a JSON object
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
        return nil // Handle nil values properly
    }

    // Expecting value as a string
    str, ok := value.(string)
    if !ok {
        return fmt.Errorf("unsupported type for geometry: %T", value)
    }

    // Convert the string representation of WKB to a byte slice
    bytes, err := hex.DecodeString(str)
    if err != nil {
        return fmt.Errorf("error decoding string to bytes: %w", err)
    }

    // Create geometry from WKB
    geom, err := geos.FromWKB(bytes)
    if err != nil {
        return fmt.Errorf("error converting WKB to geometry: %w", err)
    }

	*g = Geometry4326{geom}
    return nil
}



// func (g Geometry4326) Value() (driver.Value, error) {
//     wkb, err := g.ToWKB()  // Convert to Well-Known Binary (WKB)
//     if err != nil {
//         return nil, err
//     }
//     return wkb, nil  // Store as WKB
// }
//
// func (g *Geometry4326) Scan(value interface{}) error {
//     bytes, ok := value.([]byte)
//     if !ok {
//         return errors.New("cannot convert database value to geometry")
//     }
//
//     geom, err := geos.FromWKB(bytes)  // Use WKB to read from DB
//     if err != nil {
//         return err
//     }
//
//     *g = Geometry4326{geom}
//     return nil
// }

