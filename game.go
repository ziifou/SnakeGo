package main

func initializeGameObjects() {
	snake = &Snake{
		points:         getInitialSnakeCoordinates(),
		columnVelocity: 1,
		rowVelocity:    0,
		symbol:         SNAKE_SYMBOL,
	}

	vitamin = &Vitamin{
		point:  getInitialVitaminCoordinates(),
		symbol: VITAMIN_SYMBOL,
	}
}

func updateGameState() {
	if isGamePaused {
		return
	}
	clearScreen()
	updateSnake()
	updateVitamin()
}

func updateSnake() {
	snakeHeadX, snakeHeadY := getSnakeHeadCoordinates()
	newSnakeHead := &Coordinate{
		snakeHeadX + snake.columnVelocity,
		snakeHeadY + snake.rowVelocity,
	}
	setSnakeWithinFrame(newSnakeHead)
	snake.points = append(snake.points, newSnakeHead)
	if !isVitaminInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, snake.points[0])
		snake.points = snake.points[1:]
	} else {
		score++
		displayGameScore()
		playSound(eatSound)
	}
	if isSnakeEatingItself() {
		isGameOver = true
	}
}

func updateVitamin() {
	for isVitaminInsideSnake() {
		coordinatesToClear = append(coordinatesToClear, vitamin.point)
		vitamin.point.x, vitamin.point.y = getNewVitaminCoordinate()
	}
}
