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
	return GetOnOffLamps(yellowLamp, minutes)
}

func ConvertSingleHours(hours int) string {
	return GetOnOffLamps(redLamp, hours)
}

func GetOnOffLamps(lampColor string, time int) string {
	var onLamps strings.Builder
	var offLamps strings.Builder
	totalLamps := 4

	numberOfOnLamps := time % 5

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
