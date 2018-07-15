package datetime

import (
	"time"
)

// Interval is the aggregation of each chart bar.
type Interval string

const (
	// OneMin interval of 1 minute.
	OneMin Interval = "1m"
	// TwoMins interval of 2 minutes.
	TwoMins Interval = "2m"
	// FiveMins interval of 5 minutes.
	FiveMins Interval = "5m"
	// FifteenMins interval of 15 minutes.
	FifteenMins Interval = "15m"
	// ThirtyMins interval of thirty minutes.
	ThirtyMins Interval = "30m"
	// SixtyMins interval of 60 minutes.
	SixtyMins Interval = "60m"
	// NinetyMins interval of 90 minutes.
	NinetyMins Interval = "90m"
	// OneHour interval of 1 hour.
	OneHour Interval = "1h"
	// OneDay interval of 1 day.
	OneDay Interval = "1d"
	// FiveDay interval of 5 days.
	FiveDay Interval = "5d"
	// OneMonth interval of 1 month.
	OneMonth Interval = "1mo"
	// ThreeMonth interval of 3 months.
	ThreeMonth Interval = "3mo"
	// SixMonth interval of 6 months.
	SixMonth Interval = "6mo"
	// OneYear interval of 1 year.
	OneYear Interval = "1y"
	// TwoYear interval of 2 years.
	TwoYear Interval = "2y"
	// FiveYear interval of 5 years.
	FiveYear Interval = "5y"
	// TenYear interval of 10 years.
	TenYear Interval = "10y"
	// YTD interval of year-to-date.
	YTD Interval = "ytd"
	// Max interval.
	Max Interval = "max"
)

// Datetime is a simple time construct,
// that is either the start point or the end point
// for a chart time-series.
type Datetime struct {
	Day   int
	Month int
	Year  int
	t     *time.Time
}

// New creates a new instance of Datetime from a go time struct.
func New(t *time.Time) *Datetime {
	year, month, day := t.Date()
	return &Datetime{
		Month: int(month),
		Day:   day,
		Year:  year,
		t:     t,
	}
}

// FromUnix returns a new instance of Datetime from a unix timestamp.
func FromUnix(timestamp int) *Datetime {
	t := time.Unix(int64(timestamp), 0)
	return New(&t)
}

// Time returns a go time struct from a datetime.
func (d *Datetime) Time() *time.Time {
	if d.t != nil {
		return d.t
	}
	d.calculateTime()
	return d.Time()
}

// Unix returns a valid unix timestamp from Datetime fields.
func (d *Datetime) Unix() int {
	if d.t != nil {
		return int(d.t.Unix())
	}

	d.calculateTime()
	return d.Unix()
}

func (d *Datetime) calculateTime() {
	t := time.Date(d.Year, time.Month(d.Month), d.Day, 9, 30, 0, 0, time.Local)
	d.t = &t
}
