package mock

import (
	"github.com/feditools/go-lib/metrics"
)

// GRPCRequest is a new database query metric measurer.
type GRPCRequest struct{}

// Done is called when the grpc request is complete.
func (GRPCRequest) Done(_ bool) {}

// HTTPRequest is a new database query metric measurer.
type HTTPRequest struct{}

// Done is called when the grpc request is complete.
func (HTTPRequest) Done(_ int) {}

// DBQuery is a new database query metric measurer.
type DBQuery struct{}

// Done is called when the db query is complete.
func (DBQuery) Done(_ bool) {}

// DBCacheQuery is a new database cache query metric measurer.
type DBCacheQuery struct{}

// Done is called when the db cache query is complete.
func (DBCacheQuery) Done(_, _ bool) {}

// MetricsCollector is a mock metrics collection.
type MetricsCollector struct{}

// Close does nothing.
func (MetricsCollector) Close() error {
	return nil
}

// NewGRPCRequest creates a new grpc metrics collector.
func (MetricsCollector) NewGRPCRequest(_ string) metrics.GRPCRequest {
	return &GRPCRequest{}
}

// NewHTTPRequest creates a new http metrics collector.
func (MetricsCollector) NewHTTPRequest(_, _ string) metrics.HTTPRequest {
	return &HTTPRequest{}
}

// NewDBQuery creates a new db query metrics collector.
func (MetricsCollector) NewDBQuery(_ string) metrics.DBQuery {
	return &DBQuery{}
}

// NewDBCacheQuery creates a new db cache query metrics collector.
func (MetricsCollector) NewDBCacheQuery(_ string) metrics.DBCacheQuery {
	return &DBCacheQuery{}
}

// NewMetricsCollector creates a new mock metrics collector.
func NewMetricsCollector() (metrics.Collector, error) {
	return &MetricsCollector{}, nil
}
