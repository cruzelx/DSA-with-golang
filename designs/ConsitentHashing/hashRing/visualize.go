package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"git.sr.ht/~sbinet/gg"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
	dc.DrawString(name, x, y)
	dc.DrawCircle(x, y, 10)
	dc.LineTo(x, y)
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

	nodeColors := make(map[string]color.Color)

	for _, nodeHash := range mainNodes {
		nodeName := hr.Hashmap[nodeHash]

		if _, ok := nodeColors[nodeName]; !ok {
			r, g, b := generateContrastingColor()
			nodeColors[nodeName] = color.RGBA{r, g, b, 255}
		}

		x, y := positionOnCircle(nodeHash, circleRadius, math.MaxUint32)
		x, y = imageSize/2+x, imageSize/2+y

		dc.SetColor(nodeColors[nodeName])
		drawNode(dc, x, y, nodeRadius, nodeName)
	}

	dc.SavePNG("./analysis/consistent_hash_ring.png")
}

func generateContrastingColor() (uint8, uint8, uint8) {
	for {
		// Generate random RGB values, ensuring at least one component is dark
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		if r < 128 || g < 128 || b < 128 {
			return uint8(r), uint8(g), uint8(b)
		}
	}
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

func visualizeStandardDeviation(m *Manager) {
	const imageSize = 800
	const maxReplicaNodes = 100 // Adjust as needed

	// Calculate standard deviations for different numbers of replica nodes
	stdDevs := make([]float64, maxReplicaNodes+1)
	for replicaNodes := 0; replicaNodes <= maxReplicaNodes; replicaNodes++ {
		// Create a new Manager with the specified number of replica nodes
		tempManager := NewManager()
		tempManager.HashRing.ReplicaNodes = replicaNodes

		// Add a reasonable number of keys and nodes for the calculation
		numNodes := 10
		numKeys := 1000
		for i := 0; i < numNodes; i++ {
			tempManager.AddNode(fmt.Sprintf("node%d", i))
		}
		for i := 0; i < numKeys; i++ {
			tempManager.PutKey(KeyVal{Key: fmt.Sprintf("key%d", i), Value: "value"})
		}

		// Calculate the standard deviation of keyvals across nodes
		keyvalsInNodes := make([]int, 0)
		tempManager.Nodes.Range(func(key, value any) bool {
			keyvalsInNodes = append(keyvalsInNodes, len(value.([]KeyVal)))
			return true
		})

		stdDevs[replicaNodes] = stddev(keyvalsInNodes)
	}

	// Create a plot using gonum
	p := plot.New()

	// Draw the plot
	x := make([]float64, maxReplicaNodes+1)
	for i := 0; i <= maxReplicaNodes; i++ {
		x[i] = float64(i)
	}
	y := stdDevs
	XY := []plotter.XY{}
	for i, scale := range x {
		XY = append(XY, plotter.XY{X: scale, Y: y[i]})
	}

	XYs := plotter.XYs(XY)

	// Create a line plot
	l, err := plotter.NewLine(XYs)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(l)

	// Customize plot
	p.Title.Text = "Standard Deviation vs Number of Replica Nodes"
	p.X.Label.Text = "Number of Replica Nodes"
	p.Y.Label.Text = "Standard Deviation"

	// Save the plot as a PNG image
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "./analysis/std_dev_vs_replica_nodes.png"); err != nil {
		log.Fatal(err)
	}
}

func stddev(data []int) float64 {
	// Calculate mean
	mean := 0.0
	for _, value := range data {
		mean += float64(value)
	}
	mean /= float64(len(data))

	// Calculate variance
	variance := 0.0
	for _, value := range data {
		variance += math.Pow(float64(value)-mean, 2)
	}
	variance /= float64(len(data))

	return math.Sqrt(variance)
}
