package main

import "fmt"

func main() {
	var times = []int{123} //working as good for one value as good as for array
	var distances = []int{12345}
	var totalChances []int
	for i := 0; i < len(times); i++ {
		var chances int
		for y := 1; y < times[i]; y++ {
			if isFaster(y, times[i], distances[i]) {
				chances++
			}
		}
		totalChances = append(totalChances, chances)
	}
	fmt.Println(totalChances[0] * totalChances[1] * totalChances[2] * totalChances[3])
}

func isFaster(pressTime, totalTime, distance int) bool {
	raceDist := (totalTime - pressTime) * pressTime
	return raceDist > distance
}
