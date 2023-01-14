// Package components provides main function on components files
package components

import (
	"math"
)

func FindClosestOne(slice [][]int, x3 int, y3 int, wantedNumber int) (int, int, int) {
	minDist := math.MaxInt32
	var closestX, closestY int

	for y2, row := range slice {
		for x2, val := range row {
			if int(math.Abs(float64(val))) == wantedNumber {
				dist := int(math.Abs(float64(x2-x3)) + math.Abs(float64(y2-y3)))
				if dist < minDist {
					minDist = dist
					closestX = x2
					closestY = y2
				}
			}
		}
	}
	return closestX, closestY, minDist
}
