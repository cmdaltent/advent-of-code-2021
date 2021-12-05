package main

type line struct {
	start point
	end   point
}

func (l line) points() []point {
	if l.start.x == l.end.x {
		return l.vertically()
	}
	if l.start.y == l.end.y {
		return l.horizontally()
	}
	return l.diagonally()
}

func (l line) horizontally() []point {
	points := make([]point, 0)
	if l.start.x < l.end.x {
		for i := l.start.x; i < l.end.x; i++ {
			points = append(points, point{
				x: i,
				y: l.start.y,
			})
		}
	} else {
		for i := l.start.x; i > l.end.x; i-- {
			points = append(points, point{
				x: i,
				y: l.start.y,
			})
		}
	}
	return append(points, l.end)
}

func (l line) vertically() []point {
	points := make([]point, 0)
	if l.start.y < l.end.y {
		for i := l.start.y; i < l.end.y; i++ {
			points = append(points, point{
				x: l.start.x,
				y: i,
			})
		}
	} else {
		for i := l.start.y; i > l.end.y; i-- {
			points = append(points, point{
				x: l.start.x,
				y: i,
			})
		}
	}
	return append(points, l.end)
}

func (l line) diagonally() []point {
	points := make([]point, 0)

	var diff int64 = 0
	var xSign int64 = 1
	var ySign int64 = 1

	// 5,5 -> 2,2
	if l.start.x > l.end.x && l.start.y > l.end.y {
		diff = l.start.x - l.end.x
		xSign = -1
		ySign = -1
	}
	// 5,5 -> 8,8
	if l.end.x > l.start.x && l.end.y > l.start.y {
		diff = l.end.x - l.start.x
	}
	// 5,5 -> 8,2
	if l.end.x > l.start.x && l.start.y > l.end.y {
		diff = l.end.x - l.start.x
		ySign = -1
	}
	// 5,5 -> 2,8
	if l.start.x > l.end.x && l.end.y > l.start.y {
		diff = l.start.x - l.end.x
		xSign = -1
	}

	var counter int64 = 0
	for counter < diff {
		points = append(points, point{
			x: l.start.x + (counter * xSign),
			y: l.start.y + (counter * ySign),
		})
		counter++
	}
	return append(points, l.end)
}

type point struct {
	x int64
	y int64
}
