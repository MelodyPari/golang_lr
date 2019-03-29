package main

import (
	"fmt"

	"io"
	"os"

	"github.com/fogleman/gg"
	geojson "github.com/paulmach/go.geojson"
)

func main() {
	fmt.Print("Hello world")
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
	dc := gg.NewContext(1000, 1000)
	dc.SetHexColor("fff")

	///
	/*file2, err := os.Open("russia.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file2.Close()
	data2 := make([]byte, 64)

	geomData2 := ""
	for {
		n, err := file2.Read(data2)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		geomData2 = geomData2 + string(data2[:n])
	}
	rawFeatureJSON2 := []byte(geomData2)
	fc2, err := geojson.UnmarshalFeatureCollection(rawFeatureJSON2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Print(len(fc2.Features[0].Geometry.MultiPolygon[0][0])
	/*dc.MoveTo(fc2.Features[0].Geometry.MultiPolygon[0][0][0][0], fc2.Features[0].Geometry.MultiPolygon[0][0][0][1])
	for i := 0; i < len(fc2.Features[0].Geometry.MultiPolygon[0][0]); i++ {

		dc.LineTo(fc2.Features[0].Geometry.MultiPolygon[0][0][i][0], fc2.Features[0].Geometry.MultiPolygon[0][0][i][1])

	}*/
	///

	dc.MoveTo(fc1.Features[0].Geometry.Polygon[0][0][0], fc1.Features[0].Geometry.Polygon[0][0][1])
	for i := 0; i < 5; i++ {

		dc.LineTo(fc1.Features[0].Geometry.Polygon[0][i][0], fc1.Features[0].Geometry.Polygon[0][i][1])

	}

	dc.SetRGB(0, 0, 1)
	dc.InvertY()
	//dc.Stroke()
	dc.Fill()
	dc.SavePNG("out.png")

}
