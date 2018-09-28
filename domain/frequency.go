package domain

import (
	"fmt"
	"strconv"
)

type Frequency float64

const (
	Hz  Frequency = 1
	KHz           = 1000 * Hz
	MHz           = 1000 * KHz
	GHz           = 1000 * MHz
	THz           = 1000 * GHz
)

func (freq Frequency) Period() Period {
	return Second / Period(freq)
}

func (freq Frequency) String() string {
	var (
		units string
		div   Frequency
	)

	switch {
	case freq > THz/10:
		units = "THz"
		div = THz
	case freq > GHz/10:
		units = "GHz"
		div = GHz
	case freq > MHz/10:
		units = "MHz"
		div = MHz
	case freq > KHz/10:
		units = "kHz"
		div = KHz
	default:
		units = "Hz"
		div = Hz
	}

	return fmt.Sprintf("%s%s", strconv.FormatFloat(float64(freq/div), 'f', -1, 64), units)
}
