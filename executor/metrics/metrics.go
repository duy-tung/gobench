package metrics

type MetricType string

const (
	// Counter is single additive value. New values are simply added to the current one.
	Counter MetricType = "counter"
	// Histogram is a set of numerical values that quantify a distribution of values. New values are added to the distribution.
	Histogram = "histogram"
	// Gauge is a single non-additive value. New value replaces the previous one.
	Gauge = "gauge"
)

type Metric struct {
	Title string
	Type  MetricType
}

type Graph struct {
	Title   string
	Unit    string
	Metrics []Metric
}

type Group struct {
	Name   string
	Graphs []Graph
}
