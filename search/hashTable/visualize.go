package main

import (
	"fmt"
	"hashTable/hasher"
	"image/color"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Result struct {
	HashFunc          string
	StandardDeviation float64
	Mean              float64
	BucketSize        int
	// KeyLength         int
}

// TODO: y(standard deviation) values should correspond to x(bucket size) values
func WebPlot() {
	distribution := getDistributions()

	line := charts.NewLine()

	xData := []int{}

	for hashfunc, results := range distribution {
		data := []opts.LineData{}
		for _, v := range results {
			data = append(data, opts.LineData{Name: fmt.Sprintf("%d", v.BucketSize), Value: v.StandardDeviation, XAxisIndex: v.BucketSize})
			xData = append(xData, v.BucketSize)
		}
		line.AddSeries(hashfunc, data, charts.WithLineChartOpts(opts.LineChart{
			Smooth: false,
		}), charts.WithLabelOpts(opts.Label{
			Show: true,
		}))
	}

	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Standard Deviation vs. Bucket Size"}), charts.WithLegendOpts(opts.Legend{
		Left: "right",
		Show: true,
	}), charts.WithXAxisOpts(opts.XAxis{
		Name: "Bucket Size",
		Type: "category",
		Data: xData,
	}), charts.WithYAxisOpts(opts.YAxis{
		Name: "Standard Deviation",
		Type: "value",
	}))

	page := components.NewPage()
	page.AddCharts(line)

	f, err := os.Create("line.html")
	if err != nil {
		fmt.Println(err)
	}
	page.Render(io.MultiWriter(f))

}

func Plot() {
	distribution := getDistributions()

	p := plot.New()

	p.Title.Text = "Standard Deviation vs. Bucket Size"
	p.X.Label.Text = "Bucket Size"
	p.Y.Label.Text = "Standard Deviation"

	plotIdx := 0
	colors := []color.Color{
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{255, 0, 255, 255},
		color.RGBA{0, 255, 255, 255},
	}
	// Loop through the distribution data and create a plotter.Points
	for hashFunc, results := range distribution {
		// if hashFunc == "naiveHash" {
		// 	continue
		// }
		points := make(plotter.XYs, len(results))

		for i, result := range results {
			points[i].X = float64(result.BucketSize)
			points[i].Y = float64(result.StandardDeviation)
		}

		// Create a scatter plot
		s, err := plotter.NewScatter(points)
		if err != nil {
			log.Fatalf("Could not create scatter plot: %v", err)
		}
		s.GlyphStyle.Radius = vg.Points(3)
		s.GlyphStyle.Shape = draw.CircleGlyph{}
		s.GlyphStyle.Color = colors[plotIdx]

		// Add scatter plot to the plot
		p.Add(s)
		p.Legend.Add(hashFunc, s)
		p.Legend.Top = true
		p.Legend.Left = false

		// Update plot index and color
		plotIdx++
	}

	// Save plot to a file
	fileName := "stddev_vs_bucketsize.png"
	if err := p.Save(6*vg.Inch, 4*vg.Inch, fileName); err != nil {
		log.Fatalf("Could not save plot: %v", err)
	}
}

func getDistributions() map[string][]Result {
	hashFuncs := map[string]func(string) uint32{
		"djb2":        hasher.Djb2,
		"murmurHash3": hasher.MurmurHash3,
		"naiveHash":   hasher.NaiveHash,
		"fnv1a":       hasher.Fnv1a,
	}
	bucketSizes := []int{100, 500, 1000, 2500, 5000, 8000, 10000}
	keysByBucketSize := make([][]string, len(bucketSizes))

	for i, v := range bucketSizes {
		// keysByBucketSize[i] = generateFixedStrings(v, 10)
		keysByBucketSize[i] = generateEmails(v)
	}

	distribution := make(map[string][]Result)

	for name, function := range hashFuncs {
		for i, size := range bucketSizes {

			ht := NewHashTable(size, 80, function)

			bucket := make(map[int]int)

			for _, key := range keysByBucketSize[i] {
				index := ht._hash(key)
				if ht.Get(key) != nil {
					continue
				}
				bucket[index]++
				ht.Set(key, "Value")
			}
			std, mean := stdAndMean(bucket)
			distribution[name] = append(distribution[name], Result{HashFunc: name, StandardDeviation: std, Mean: mean, BucketSize: size})
		}
	}
	return distribution
}

func stdAndMean(buckets map[int]int) (float64, float64) {
	n := float64(len(buckets))
	sum := 0
	for _, val := range buckets {
		sum += val
	}
	mean := float64(sum) / n

	variance := float64(0)
	for _, val := range buckets {
		deviation := float64(val) - mean
		variance += deviation * deviation
	}
	std := math.Sqrt(variance / n)

	return std, mean

}

func generateEmails(size int) []string {
	results := []string{}

	for i := 0; i < size; i++ {
		results = append(results, randomdata.Email())
	}
	return results
}

func generateFixedStrings(count, size int) []string {
	rand.Seed(time.Now().UnixNano())

	var result []string
	charset := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < count; i++ {
		str := make([]byte, size)
		for j := 0; j < size; j++ {
			str[j] = charset[rand.Intn(len(charset))]
		}
		result = append(result, string(str))
	}

	return result
}
