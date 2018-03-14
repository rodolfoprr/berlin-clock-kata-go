package main

import (
	"strings"
)

var redLamp = "R"
var yellowLamp = "Y"
var offLamp = "O"

func ConvertSeconds(seconds int) string {
	if seconds%2 == 0 {
		return yellowLamp
	}

	return offLamp
}

func ConvertSingleMinutes(minutes int) string {
	getNumberOfOnLamps := func() int {
		return minutes % 5
	}

	totalLamps := 4

	return getLamps(yellowLamp, getNumberOfOnLamps, totalLamps)
}

func ConvertSingleHours(hours int) string {
	getNumberOfOnLamps := func() int {
		return hours % 5
	}

	totalLamps := 4

	return getLamps(redLamp, getNumberOfOnLamps, totalLamps)
}

func ConvertFiveHours(hours int) string {
	getNumberOfOnLamps := func() int {
		return hours / 5
	}

	totalLamps := 4

	return getLamps(redLamp, getNumberOfOnLamps, totalLamps)
}

func ConvertFiveMinutes(minutes int) string {
	getNumberOfOnLamps := func() int {
		return minutes / 5
	}

	totalLamps := 11

	return strings.Replace(getLamps(yellowLamp, getNumberOfOnLamps, totalLamps), "YYY", "YYR", -1)
}

func getLamps(lampColor string, getNumberOfOnLamps func() int, totalLamps int) string {
	var onLamps strings.Builder
	var offLamps strings.Builder

	numberOfOnLamps := getNumberOfOnLamps()

	for i := 0; i < numberOfOnLamps; i++ {
		onLamps.WriteString(lampColor)
	}

	for i := numberOfOnLamps; i < totalLamps; i++ {
		offLamps.WriteString(offLamp)
	}

	return onLamps.String() + offLamps.String()
}

func main() {

}
