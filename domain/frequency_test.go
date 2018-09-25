package domain

import (
	"testing"
)

func Test__Frequency_Period(t *testing.T) {
	testCases := []struct {
		freq     Frequency
		expected Period
	}{
		{
			0.001 * Hz,
			1000 * Second,
		}, {
			0.01 * Hz,
			100 * Second,
		}, {
			0.1 * Hz,
			10 * Second,
		}, {
			1 * Hz,
			1 * Second,
		}, {
			4 * Hz,
			250 * Millisecond,
		}, {
			1.5 * KHz,
			666.6666666666666 * Microsecond,
		}, {
			4 * KHz,
			250 * Microsecond,
		}, {
			4 * MHz,
			250 * Nanosecond,
		}, {
			1 * GHz,
			1 * Nanosecond,
		}, {
			4 * GHz,
			0.25 * Nanosecond,
		}, {
			4 * THz,
			0.00025 * Nanosecond,
		},
	}

	for _, tc := range testCases {
		res := tc.freq.Period()
		//t.Logf("for freq %v expected duration per cycle: %v, got %v", tc.freq, tc.expected, res)
		if res != tc.expected {
			t.Errorf("[FAILED] expected duration per cycle for freq %v: %v, got: %v", tc.freq, tc.expected, res)
		}
	}
}
