package prometheus

import "github.com/prometheus/client_golang/prometheus/collectors/gocollector"

// NewGoCollector is the obsolete version of gocollector.NewGoCollector.
// See there for documentation.
//
// Deprecated: Use gocollector.NewGoCollector instead.
func NewGoCollector() Collector {
	return gocollector.NewGoCollector() // CIRCULAR DEPENDENCY.
}
