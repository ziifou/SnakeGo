package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

var Screen tcell.Screen
var screenWidth, screenHeight int
var isGamePaused, isGameOver bool
var score int
var snakeSpeed time.Duration

const FRAME_WIDTH = 80
const FRAME_HEIGHT = 15

var snake *Snake
var vitamin *Vitamin

func main() {
	initScreen()
	initSounds() // Initialize sounds here
	initializeGameObjects()
	displayFrame()
	displayGameScore()
	userInput := readUserInput()
	snakeSpeed = 75 * time.Millisecond // Initialize snake speed
	var key string
	for !isGameOver {
		if isGamePaused {
			displayGamePausedInfo()
		}
		key = getUserInput(userInput)
		handleUserInput(key)
		updateGameState()
		displayGameObjects()
		time.Sleep(snakeSpeed) // Use the snakeSpeed variable here
	}

	displayGameOverInfo()
	playSound(gameOverSound) // Play game over sound
	time.Sleep(3 * time.Second)
	Screen.Fini()
}
