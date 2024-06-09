package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func getUserInput(userInput chan string) string {
	var key string
	select {
	case key = <-userInput:
	default:
		key = ""
	}
	return key
}

func readUserInput() chan string {
	userInput := make(chan string)
	go func() {
		for {
			switch ev := Screen.PollEvent().(type) {
			case *tcell.EventKey:
				userInput <- ev.Name()
			}
		}
	}()
	return userInput
}

func handleUserInput(key string) {
	if key == "Rune[q]" {
		Screen.Fini()
		os.Exit(0)
	} else if key == "Rune[p]" {
		isGamePaused = !isGamePaused
		playSound(pauseSound) // Play pause sound
	} else if key == "Rune[ ]" { // Handle spacebar input
		if snakeSpeed > 25*time.Millisecond {
			snakeSpeed -= 10 * time.Millisecond
		}
	} else if !isGamePaused {
		if key == "Up" && snake.rowVelocity == 0 {
			snake.rowVelocity = -1
			snake.columnVelocity = 0
		} else if key == "Down" && snake.rowVelocity == 0 {
			snake.rowVelocity = 1
			snake.columnVelocity = 0
		} else if key == "Left" && snake.columnVelocity == 0 {
			snake.rowVelocity = 0
			snake.columnVelocity = -1
		} else if key == "Right" && snake.columnVelocity == 0 {
			snake.rowVelocity = 0
			snake.columnVelocity = 1
		}
	}
}
