package main

import (
	"github.com/gonum/plot/plotter"
	"github.com/sunyifan112358/gpu_mem_analyze/figure"
)

func main() {
	// LoadExperimentFromFile("raw_data/b128_g2.out")
	iterVsAccuracy := new(figure.IterVsAccuracy)
	iterVsAccuracy.Initialize()

	series := new(figure.Series)
	series.Name = "Test"
	series.XYs = make(plotter.XYs, 2)
	series.XYs[0].X = 0
	series.XYs[0].Y = 0
	series.XYs[1].X = 1
	series.XYs[1].Y = 1
	iterVsAccuracy.AddSeries(*series)

	iterVsAccuracy.Save()
}
