package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p, _ := plot.New()

	p.Title.Text = "Gonum plot示例"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	points := make(plotter.XYs, 4)

	points[0].X = 0.0
	points[0].Y = 0.0

	points[1].X = 1.0
	points[1].Y = 1.0

	points[2].X = 2.0
	points[2].Y = 4.0

	points[3].X = 3.0
	points[3].Y = 9.0

	plotutil.AddLinePoints(p, "y = x * x", points)

	p.Save(4*vg.Inch, 4*vg.Inch, "points.png")
}
