package main

import (
	"math/rand"
	"time"
)

type Vitamin struct {
	point  *Coordinate
	symbol rune
}

func getInitialVitaminCoordinates() *Coordinate {
	vitaminInitialCoordinate := &Coordinate{FRAME_WIDTH / 2, FRAME_HEIGHT / 2}
	transformCoordinateInsideFrame(vitaminInitialCoordinate)

	return vitaminInitialCoordinate
}

func getNewVitaminCoordinate() (int, int) {
	rand.Seed(time.Now().UnixMicro())
	randomX := rand.Intn(FRAME_WIDTH - 2*FRAME_BORDER_THICKNESS)
	randomY := rand.Intn(FRAME_HEIGHT - 2*FRAME_BORDER_THICKNESS)

	newCoordinate := &Coordinate{
		randomX, randomY,
	}

	transformCoordinateInsideFrame(newCoordinate)

	return newCoordinate.x, newCoordinate.y
}

func isVitaminInsideSnake() bool {
	for _, snakeCoordinate := range snake.points {
		if snakeCoordinate.x == vitamin.point.x && snakeCoordinate.y == vitamin.point.y {
			return true
		}
	}
	return false
}
