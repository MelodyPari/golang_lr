package main

import (
	"fmt"

	"io"
	"os"

	"github.com/davvo/mercator"
	"github.com/fogleman/gg"
	geojson "github.com/paulmach/go.geojson"
)

func main() {
	fmt.Print("Hello world")
	var coordX float64
	var coordY float64
	var z = 10

	fmt.Print("Введите координату x: ")
	fmt.Scan(&coordX)
	fmt.Print("Введите координату y: ")
	fmt.Scan(&coordY)
	var lat, lon = mercator.MetersToLatLon(coordX, coordY)
	var tx, ty = mercator.LatLonToTile(lat, lon, z)
	file, err := os.Open("lr.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)

	geomData := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		geomData = geomData + string(data[:n])
	}
	// Feature Collection
	rawFeatureJSON := []byte(geomData)

	fc1, err := geojson.UnmarshalFeatureCollection(rawFeatureJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dc := gg.NewContext(256, 256)
	dc.SetHexColor("fff")

	//dc.MoveTo(fc1.Features[0].Geometry.Polygon[0][0][0], fc1.Features[0].Geometry.Polygon[0][0][1])
	lat, lon = mercator.MetersToLatLon(fc1.Features[0].Geometry.Polygon[0][2][0], fc1.Features[0].Geometry.Polygon[0][2][1])
	tx, ty = mercator.LatLonToTile(lat, lon, z)
	fmt.Printf("Tile (zoom %d): %d, %d\n", z, tx, ty)
	/*dc.MoveTo(tx, ty)
	tx, ty = mercator.LatLonToTile(lat, lon, z)
	for i := 0; i < 5; i++ {
		lat, lon = mercator.MetersToLatLon(fc1.Features[0].Geometry.Polygon[0][i][0], fc1.Features[0].Geometry.Polygon[0][i][1])
		tx, ty = mercator.LatLonToTile(lat, lon, z)
		dc.LineTo(tx,ty)

	}*/

	dc.SetRGB(0, 0, 1)
	dc.InvertY()
	dc.Fill()
	dc.SavePNG("out.png")

}
