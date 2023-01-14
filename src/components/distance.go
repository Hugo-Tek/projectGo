// Package components provides main function on components files
package components

import (
	"math"
)

// FindClosestOne func to find the closest one
func FindClosestOne(slice [][]int, x3 int, y3 int, wantedNumber int) (closestX int, closestY int, minDist int) {
	minDist = math.MaxInt32

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
