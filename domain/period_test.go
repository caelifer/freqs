package domain

import (
	"testing"
)

func Test__Period_String(t *testing.T) {
	testCases := []struct {
		period   Period
		expected string
	}{
		{
			1 * Minute,
			"1m0s",
		}, {
			100 * Second,
			"1m40s",
		}, {
			100 * Minute,
			"1h40m0s",
		}, {
			10*Hour + 23*Minute + 50*Second + 20*Millisecond,
			"10h23m50.02s",
		}, {
			2*Day + 10*Hour + 23*Minute + 50*Second + 20*Millisecond,
			"2d10h23m50.02s",
		}, {
			7*Week + 10*Hour + 23*Minute + 50*Second + 20*Millisecond,
			"7w10h23m50.02s",
		}, {
			5000 * Nanosecond,
			"5µs",
		}, {
			1*Week + 5000*Microsecond,
			"1w0.005s",
		}, {
			1 * Millisecond,
			"1ms",
		}, {
			1 * Microsecond,
			"1µs",
		}, {
			1 * Nanosecond,
			"1ns",
		},
	}

	for _, tc := range testCases {
		if res := tc.period.String(); res != tc.expected {
			t.Errorf("for Period(%v) expecting: %v but getting: %v", tc.period, tc.expected, res)
		}
	}
}
