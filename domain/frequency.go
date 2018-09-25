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
