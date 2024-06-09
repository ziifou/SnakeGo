package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func initScreen() {
	var err error
	Screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = Screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	Screen.SetStyle(defStyle)
	screenWidth, screenHeight = Screen.Size()

	if screenWidth < FRAME_WIDTH || screenHeight < FRAME_HEIGHT {
		fmt.Printf("The game frame is defined with %d width and %d height. Increase terminal size and try again ", FRAME_WIDTH, FRAME_HEIGHT)
		os.Exit(1)
	}
}

func displayGameObjects() {
	displaySnake()
	displayVitamin()
	Screen.Show()
}

func displaySnake() {
	style := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen.TrueColor())
	for _, snakeCoordinate := range snake.points {
		print(snakeCoordinate.x, snakeCoordinate.y, 1, 1, style, snake.symbol)
	}
}

func displayVitamin() {
	style := tcell.StyleDefault.Foreground(tcell.ColorDarkRed.TrueColor())
	print(vitamin.point.x, vitamin.point.y, 1, 1, style, vitamin.symbol)
}

func displayGamePausedInfo() {
	_, frameY := getFrameOrigin()
	printAtCenter(frameY-2, "you deserve a break !!", true)
	printAtCenter(frameY-1, "Press p to continue the adventure", true)
}

func displayGameOverInfo() {
	centerY := (screenHeight - FRAME_HEIGHT) / 2
	printAtCenter(centerY-1, "Game Over!", false)
	printAtCenter(centerY, fmt.Sprintf("Score Final : %d", score), false)
}

func displayGameScore() {
	_, frameY := getFrameOrigin()
	printAtCenter(frameY+FRAME_HEIGHT+2, fmt.Sprintf("Score : %d", score), false)
}
