package main

type Snake struct {
	points                      []*Coordinate
	columnVelocity, rowVelocity int
	symbol                      rune
}

func getInitialSnakeCoordinates() []*Coordinate {
	snakeInitialCoordinate1 := &Coordinate{8, 4}
	transformCoordinateInsideFrame(snakeInitialCoordinate1)

	snakeInitialCoordinate2 := &Coordinate{8, 5}
	transformCoordinateInsideFrame(snakeInitialCoordinate2)

	snakeInitialCoordinate3 := &Coordinate{8, 6}
	transformCoordinateInsideFrame(snakeInitialCoordinate3)

	snakeInitialCoordinate4 := &Coordinate{8, 7}
	transformCoordinateInsideFrame(snakeInitialCoordinate4)

	return []*Coordinate{
		{snakeInitialCoordinate1.x, snakeInitialCoordinate1.y},
		{snakeInitialCoordinate2.x, snakeInitialCoordinate2.y},
		{snakeInitialCoordinate3.x, snakeInitialCoordinate3.y},
		{snakeInitialCoordinate4.x, snakeInitialCoordinate4.y},
	}
}

func getSnakeHeadCoordinates() (int, int) {
	snakeHead := snake.points[len(snake.points)-1]
	return snakeHead.x, snakeHead.y
}

func setSnakeWithinFrame(snakeCoordinate *Coordinate) {
	leftX, topY, rightX, bottomY := getBoundaries()

	if snakeCoordinate.y <= topY {
		snakeCoordinate.y = bottomY - 1
	} else if snakeCoordinate.y >= bottomY {
		snakeCoordinate.y = topY + 1
	} else if snakeCoordinate.x >= rightX {
		snakeCoordinate.x = leftX + 1
	} else if snakeCoordinate.x <= leftX {
		snakeCoordinate.x = rightX - 1
	}
}

func isSnakeEatingItself() bool {
	headX, headY := getSnakeHeadCoordinates()
	for _, snakeCoordinate := range snake.points[:len(snake.points)-1] {
		if headX == snakeCoordinate.x && headY == snakeCoordinate.y {
			return true
		}
	}
	return false
}
