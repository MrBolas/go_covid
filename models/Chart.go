package models

import (
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type CustomTSChart struct {
	chart.Chart
}

func (c *CustomTSChart) Initialize(timeseries []time.Time, valueseries []float64, graphLegend string) {
	c.Series = []chart.Series{
		chart.TimeSeries{
			Name:    graphLegend,
			XValues: timeseries,
			YValues: valueseries,
		},
	}

	c.Elements = []chart.Renderable{
		chart.Legend(&c.Chart),
	}
}

func (c *CustomTSChart) SetXAxis(XAxis chart.XAxis) {
	c.XAxis = XAxis
}

func (c *CustomTSChart) SetYAxis(YAxis chart.YAxis) {
	c.YAxis = YAxis
}
