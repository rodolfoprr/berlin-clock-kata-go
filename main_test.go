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

func TestFiveMinutesRow(t *testing.T) {
	testcases := []testcase{
		{0, "OOOOOOOOOOO"},
		{5, "YOOOOOOOOOO"},
		{10, "YYOOOOOOOOO"},
		{15, "YYROOOOOOOO"},
		{20, "YYRYOOOOOOO"},
		{25, "YYRYYOOOOOO"},
		{30, "YYRYYROOOOO"},
		{35, "YYRYYRYOOOO"},
		{40, "YYRYYRYYOOO"},
		{45, "YYRYYRYYROO"},
		{50, "YYRYYRYYRYO"},
		{55, "YYRYYRYYRYY"},
	}

	Assert(t, testcases, ConvertFiveMinutes)
}

func Assert(t *testing.T, testcases []testcase, convert func(int) string) {
	for _, test := range testcases {
		convertedTime := convert(test.input)

		if convertedTime != test.expected {
			t.Errorf("Color was incorrect, %d got: %s, want: %s.", test.input, convertedTime, test.expected)
		}
	}
}

type testcase struct {
	input    int
	expected string
}
