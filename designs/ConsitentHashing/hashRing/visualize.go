package main

import (
	"math"

	"git.sr.ht/~sbinet/gg"
)

type Point struct {
	X, Y float64
}

func drawCircle(dc *gg.Context, x, y, radius float64) {
	dc.DrawArc(x, y, radius, 0, 2*math.Pi)
	dc.SetRGB(0, 0, 1)
	dc.Stroke()
}

func drawNode(dc *gg.Context, x, y, radius float64, name string) {
	dc.SetRGB(1, 0, 1)
	dc.DrawString(name, x, y)
	dc.DrawCircle(x, y, 10)
	// dc.DrawArc(x, y, radius, startAngle, endAngle)
	dc.LineTo(x, y)
	dc.SetRGB(0, 0, 1)
	dc.FillPreserve()
	dc.Stroke()
}

func visualizeHashRing(hr *HashRing) {
	const imageSize = 800
	const circleRadius = 225
	const nodeRadius = 5

	dc := gg.NewContext(imageSize, imageSize)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Draw the circle
	drawCircle(dc, imageSize/2, imageSize/2, circleRadius)

	mainNodes := extractUniqueValuesFromMap(hr.SortedHash)

	for _, nodeHash := range mainNodes {
		x, y := positionOnCircle(nodeHash, circleRadius, math.MaxUint32)
		x, y = imageSize/2+x, imageSize/2+y

		drawNode(dc, x, y, nodeRadius, hr.Hashmap[nodeHash])
	}

	dc.SavePNG("consistent_hash_ring.png")
}

func positionOnCircle(value uint32, radius int, maxVal uint32) (float64, float64) {
	normalizeValue := float64(value) / float64(maxVal)
	angleInDegrees := normalizeValue * 360
	angleInRadian := angleInDegrees * math.Pi / 180.0

	x, y := float64(radius)*math.Cos(angleInRadian), float64(radius)*math.Sin(angleInRadian)
	return x, y
}

func extractUniqueValuesFromMap(hashes []uint32) []uint32 {
	uniqueValues := make([]uint32, 0)
	seen := make(map[uint32]bool)

	for _, value := range hashes {
		if !seen[value] {
			uniqueValues = append(uniqueValues, value)
			seen[value] = true
		}
	}

	return uniqueValues
}
