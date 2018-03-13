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
	return getOnOffLamps(yellowLamp, func() int {
		return minutes % 5
	})
}

func ConvertSingleHours(hours int) string {
	return getOnOffLamps(redLamp, func() int {
		return hours % 5
	})
}

func ConvertFiveHours(hours int) string {
	return getOnOffLamps(redLamp, func() int {
		return hours / 5
	})
}

func getOnOffLamps(lampColor string, getNumberOfOnLamps func() int) string {
	var onLamps strings.Builder
	var offLamps strings.Builder
	totalLamps := 4

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
