package statsd

import (
	"fmt"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/go-lib/metrics"
)

// DBCacheQuery is a database cache query metric measurer.
type DBCacheQuery struct {
	s    statsd.Statter
	rate float32

	name  string
	start time.Time
}

// NewDBCacheQuery creates a new db cache query metrics collector.
func (m *Module) NewDBCacheQuery(name string) metrics.DBCacheQuery {
	return &DBCacheQuery{
		s:    m.s,
		rate: m.rate,

		name:  name,
		start: time.Now(),
	}
}

// Done is called when the db cache query is complete.
func (d *DBCacheQuery) Done(hit, isError bool) {
	l := logger.WithField("type", "DBCacheQuery").WithField("func", "Done")

	t := time.Since(d.start)

	err := d.s.TimingDuration(
		metrics.StatDBCacheQueryTiming,
		t,
		d.rate,
		statsd.Tag{"name", d.name},
		statsd.Tag{"hit", fmt.Sprintf("%v", hit)},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("kind", "timing").Warn(err.Error())
	}

	err = d.s.Inc(
		metrics.StatDBCacheQueryCount,
		1,
		d.rate,
		statsd.Tag{"name", d.name},
		statsd.Tag{"hit", fmt.Sprintf("%v", hit)},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("kind", "count").Warn(err.Error())
	}
}

// DBQuery is a database query metric measurer.
type DBQuery struct {
	s    statsd.Statter
	rate float32

	name  string
	start time.Time
}

// NewDBQuery creates a new db query metrics collector.
func (m *Module) NewDBQuery(name string) metrics.DBQuery {
	return &DBQuery{
		s:    m.s,
		rate: m.rate,

		name:  name,
		start: time.Now(),
	}
}

// Done is called when the db query is complete.
func (d *DBQuery) Done(isError bool) {
	l := logger.WithField("type", "DBQuery").WithField("func", "Done")

	t := time.Since(d.start)

	err := d.s.TimingDuration(
		metrics.StatDBQueryTiming,
		t,
		d.rate,
		statsd.Tag{"name", d.name},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("kind", "timing").Warn(err.Error())
	}

	err = d.s.Inc(
		metrics.StatDBQueryCount,
		1,
		d.rate,
		statsd.Tag{"name", d.name},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("kind", "count").Warn(err.Error())
	}
}
