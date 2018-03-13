package main

import "testing"

func TestSecondsRow(t *testing.T) {
	testcases := []testcase{
		{0, "Y"},
		{1, "O"},
		{2, "Y"},
		{3, "O"},
		{59, "O"},
	}

	Assert(t, testcases, ConvertSeconds)
}

func TestSingleMinutesRow(t *testing.T) {
	testcases := []testcase{
		{0, "OOOO"},
		{9, "YYYY"},
		{59, "YYYY"},
		{31, "YOOO"},
		{32, "YYOO"},
		{33, "YYYO"},
		{34, "YYYY"},
		{35, "OOOO"},
		{5, "OOOO"},
	}

	Assert(t, testcases, ConvertSingleMinutes)
}

func TestSingleHoursRow(t *testing.T) {
	testcases := []testcase{
		{0, "OOOO"},
		{16, "ROOO"},
		{17, "RROO"},
		{18, "RRRO"},
		{19, "RRRR"},
	}

	Assert(t, testcases, ConvertSingleHours)
}

func TestFiveHoursRow(t *testing.T) {
	testcases := []testcase{
		{5, "ROOO"},
		{10, "RROO"},
		{15, "RRRO"},
		{20, "RRRR"},
		{0, "OOOO"},
	}

	Assert(t, testcases, ConvertFiveHours)
}

func Assert(t *testing.T, testcases []testcase, f func(int) string) {
	for _, test := range testcases {
		minutes := f(test.input)

		if minutes != test.expected {
			t.Errorf("Color was incorrect, %d got: %s, want: %s.", test.input, minutes, test.expected)
		}
	}
}

type testcase struct {
	input    int
	expected string
}
