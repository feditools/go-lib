package statsd

import (
	"sync"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/go-lib/metrics"
)

// Module represents a statsd metrics collector.
type Module struct {
	s statsd.Statter

	rate                 float32
	systemCollectionOnce sync.Once
	systemCollectionRate time.Duration

	done chan bool
}

// New creates a new Statsd metrics module.
func New(address, prefix string) (metrics.Collector, error) {
	statsConfig := &statsd.ClientConfig{
		Address: address,
		Prefix:  prefix,
	}
	client, err := statsd.NewClientWithConfig(statsConfig)
	if err != nil {
		return nil, err
	}

	m := &Module{
		s: client,

		rate:                 1.0,
		systemCollectionRate: 10 * time.Second,

		done: make(chan bool),
	}

	m.systemCollectionOnce.Do(m.systemCollector)

	return m, nil
}

// Close closes the statsd metrics collector.
func (m *Module) Close() error {
	return m.s.Close()
}
