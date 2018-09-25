package domain

import (
	"bytes"
	"strconv"
)

type Period float64

const (
	Nanosecond  Period = 1
	Microsecond        = 1000 * Nanosecond
	Millisecond        = 1000 * Microsecond
	Second             = 1000 * Millisecond
	Minute             = 60 * Second
	Hour               = 60 * Minute
	Day                = 24 * Hour
	Week               = 7 * Day
)

func (p Period) Frequency() Frequency {
	return Frequency(Second / p)
}

func (p Period) String() string {
	if p >= Second {
		return convertMacroTime(p)
	}

	// Default is nanosecond precision
	unit := "ns"

	switch {
	case p >= Millisecond:
		unit = "ms"
		p /= Millisecond
	case p >= Microsecond:
		unit = "µs"
		p /= Microsecond
	}

	return strconv.FormatFloat(float64(p), 'f', -1, 64) + unit
}

func convertMacroTime(p Period) string {
	var buf bytes.Buffer
	if p >= Week {
		weeks := int(p / Week)
		buf.WriteString(strconv.Itoa(weeks) + "w")
		p -= Period(weeks) * Week
	}
	if p >= Day {
		days := int(p / Day)
		buf.WriteString(strconv.Itoa(days) + "d")
		p -= Period(days) * Day
	}
	if p >= Hour {
		hours := int(p / Hour)
		buf.WriteString(strconv.Itoa(hours) + "h")
		p -= Period(hours) * Hour
	}
	if p >= Minute {
		minutes := int(p / Minute)
		buf.WriteString(strconv.Itoa(minutes) + "m")
		p -= Period(minutes) * Minute
	}

	seconds := p / Second
	buf.WriteString(strconv.FormatFloat(float64(seconds), 'f', -1, 64))

	return buf.String() + "s"
}
