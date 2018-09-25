package domain

import (
	"bytes"
	"fmt"
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

type Frequency float64

const (
	Hz  Frequency = 1
	KHz           = 1000 * Hz
	MHz           = 1000 * KHz
	GHz           = 1000 * MHz
	THz           = 1000 * GHz
)

func (freq Frequency) Period() Period {
	return Period(float64(Second) / float64(freq))
}

func (freq Frequency) String() string {
	var (
		tunit string
		fmod  Frequency
	)

	switch {
	case freq > THz/10:
		tunit = "THz"
		fmod = THz
	case freq > GHz/10:
		tunit = "GHz"
		fmod = GHz
	case freq > MHz/10:
		tunit = "MHz"
		fmod = MHz
	case freq > KHz/10:
		tunit = "kHz"
		fmod = KHz
	default:
		tunit = "Hz"
		fmod = Hz
	}

	return fmt.Sprintf("%s%s", strconv.FormatFloat(float64(freq)/float64(fmod), 'f', -1, 64), tunit)
}
