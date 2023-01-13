package src

func IsAdjacent(x1, y1, x2, y2 int) bool {
	if x1 == x2 && y1 == y2 {
		return false
	}
	if (x1 == x2 && (y1 == y2+1 || y1 == y2-1)) || (y1 == y2 && (x1 == x2+1 || x1 == x2-1)) {
		return true
	}
	return false
}
