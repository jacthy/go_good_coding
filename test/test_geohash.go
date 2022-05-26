package main

import (
	"fmt"
	"github.com/jonas-p/go-shp"
	"github.com/mmcloughlin/geohash"
	"github.com/twpayne/go-geom"
	"log"
	"strconv"
)

func testGeoHash() {
	lat := 30.549608
	lon := 114.376971
	hash_base32 := geohash.EncodeWithPrecision(lat, lon, 8)
	fmt.Println(hash_base32)

	neighbors := geohash.Neighbors(hash_base32)
	hashs := append(neighbors, hash_base32)

	geomMap := make(map[string]*geom.Polygon, 9)
	for _, hash := range hashs {
		box := geohash.BoundingBox(hash)
		polygon, _ := geom.NewPolygon(geom.XY).SetCoords([][]geom.Coord{
			{
				{box.MaxLng, box.MaxLat},
				{box.MaxLng, box.MinLat},
				{box.MinLng, box.MinLat},
				{box.MinLng, box.MaxLat},
				{box.MaxLng, box.MaxLat},
			}})
		geomMap[hash] = polygon
	}
	polygonMap := map[string]*shp.PolyLine{}
	for key, multiPlygon := range geomMap {
		coordsMultiPolygon := multiPlygon.Coords()
		points := make([][]shp.Point, len(coordsMultiPolygon), len(coordsMultiPolygon))
		for index, coordsPolygon := range coordsMultiPolygon {
			points2 := make([]shp.Point, len(coordsPolygon), len(coordsPolygon))
			for j, coord := range coordsPolygon {
				x := coord.X()
				y := coord.Y()
				point := shp.Point{x, y}
				points2[j] = point
			}
			points[index] = points2
		}
		polygonTemp := shp.NewPolyLine(points)
		polygonMap[key] = polygonTemp
	}

	// points to write

	fields := []shp.Field{
		// String attribute field with length 25
		shp.StringField("base_32", 25),
		shp.StringField("binary", 50),
	}
	// create and open a shapefile for writing points
	shape, err := shp.Create("./pop2pop/polygons.shp", shp.POLYGON)
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	// setup fields for attributes
	shape.SetFields(fields)

	// write points and attributes
	cursor := 0
	for key, polygon := range polygonMap {
		shape.Write(polygon)
		// write attribute for object n for field 0 (NAME)
		toInt, _ := geohash.ConvertStringToInt(key)
		binary := fmt.Sprintf("%b", toInt)
		shape.WriteAttribute(cursor, 0, key)
		shape.WriteAttribute(cursor, 1, binary)
		cursor++
	}

	points2 := []shp.Point{
		shp.Point{10.0, 10.0},
		shp.Point{10.0, 15.0},
		shp.Point{15.0, 15.0},
		shp.Point{15.0, 10.0},
	}

	// fields to write
	fields2 := []shp.Field{
		// String attribute field with length 25
		shp.StringField("NAME", 25),
	}

	// create and open a shapefile for writing points
	shape2, err := shp.Create("./pop2pop/points.shp", shp.POINT)
	if err != nil {
		log.Fatal(err)
	}
	defer shape2.Close()

	// setup fields for attributes
	shape2.SetFields(fields2)

	// write points and attributes
	for n, point := range points2 {
		shape2.Write(&point)
		// write attribute for object n for field 0 (NAME)
		shape2.WriteAttribute(n, 0, "Point "+strconv.Itoa(n+1))
	}
}
