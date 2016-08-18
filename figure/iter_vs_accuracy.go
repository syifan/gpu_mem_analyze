package figure

import (
	"fmt"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

type Series struct {
	Name string
	XYs  plotter.XYs
}

type IterVsAccuracy struct {
	rawPlot *plot.Plot
	Data    []Series
}

func (f *IterVsAccuracy) Initialize() {
	p, err := plot.New()
	p.Title.Text = "Test"
	f.rawPlot = p
	if err != nil {
		panic(err)
	}
}

func (f *IterVsAccuracy) AddSeries(series Series) {
	f.Data = append(f.Data, series)

	fmt.Println(f.rawPlot)
	err := plotutil.AddLinePoints(f.rawPlot,
		series.Name, series.XYs)
	if err != nil {
		panic(err)
	}
	fmt.Println(f.rawPlot)
}

func (f *IterVsAccuracy) Save() {
	if err := f.rawPlot.Save(10*vg.Inch, 8*vg.Inch, "figure.eps"); err != nil {
		panic(err)
	}
}
