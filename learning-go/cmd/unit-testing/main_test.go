package main

import "testing"

func Test_sum_oneplusone(t *testing.T) {
	want := 2
	got := sum(1, 1)

	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func Test_sum_oneplustwo(t *testing.T) {
	want := 3
	got := sum(1, 2)

	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func Test_RangeOfValue(t *testing.T) {

	testCases := []struct {
		testCase string
		x        int
		y        int
		want     int
	}{
		{"One and one", 1, 1, 2},
		{"Odd and even", 1, 2, 3},
		{"Even and even", 2, 2, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.testCase, func(t *testing.T) {
			got := sum(tc.x, tc.y)
			if got != tc.want {
				t.Errorf("for (%d+%d) want: %d, got: %d", tc.x, tc.y, tc.want, got)
			}

		})
	}
}
