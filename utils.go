package main

import "github.com/gdamore/tcell/v2"

type Coordinate struct {
	x, y int
}

var coordinatesToClear []*Coordinate

const (
	FRAME_BORDER_THICKNESS    = 1
	FRAME_BORDER_VERTICAL     = '*'
	FRAME_BORDER_HORIZONTAL   = '*'
	FRAME_BORDER_TOP_LEFT     = '*'
	FRAME_BORDER_TOP_RIGHT    = '*'
	FRAME_BORDER_BOTTOM_RIGHT = '*'
	FRAME_BORDER_BOTTOM_LEFT  = '*'
	SNAKE_SYMBOL              = '█'
	VITAMIN_SYMBOL            = 'O'
)

func print(x, y, w, h int, style tcell.Style, char rune) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			Screen.SetContent(x+i, y+j, char, nil, style)
		}
	}
}

func printAtCenter(startY int, content string, trackClear bool) {
	startX := (screenWidth - len(content)) / 2
	for i := 0; i < len(content); i++ {
		print(startX+i, startY, 1, 1, tcell.StyleDefault, rune(content[i]))
		if trackClear {
			coordinatesToClear = append(coordinatesToClear, &Coordinate{startX + i, startY})
		}
	}
	Screen.Show()
}

func clearScreen() {
	for _, coordinate := range coordinatesToClear {
		print(coordinate.x, coordinate.y, 1, 1, tcell.StyleDefault, ' ')
	}
}

func printUnfilledRectangle(xOrigin, yOrigin, width, height, borderThickness int, horizontalOutline, verticalOutline, topLeftOutline, topRightOutline, bottomRightOutline, bottomLeftOutline rune) {
	var upperBorder, lowerBorder rune
	verticalBorder := verticalOutline
	for i := 0; i < width; i++ {
		if i == 0 {
			upperBorder = topLeftOutline
			lowerBorder = bottomLeftOutline
		} else if i == width-1 {
			upperBorder = topRightOutline
			lowerBorder = bottomRightOutline
		} else {
			upperBorder = horizontalOutline
			lowerBorder = horizontalOutline
		}
		// upper boundry
		print(xOrigin+i, yOrigin, borderThickness, borderThickness, tcell.StyleDefault, upperBorder)
		// lower boundry
		print(xOrigin+i, yOrigin+height-1, borderThickness, borderThickness, tcell.StyleDefault, lowerBorder)
	}

	// side boundry
	for i := 1; i < height-1; i++ {
		print(xOrigin, yOrigin+i, borderThickness, borderThickness, tcell.StyleDefault, verticalBorder)
		print(xOrigin+width-1, yOrigin+i, borderThickness, borderThickness, tcell.StyleDefault, verticalBorder)
	}
}

func transformCoordinateInsideFrame(coordinate *Coordinate) {
	leftX, topY, rightX, bottomY := getBoundaries()
	coordinate.x += leftX + FRAME_BORDER_THICKNESS
	coordinate.y += topY + FRAME_BORDER_THICKNESS
	for coordinate.x >= rightX {
		coordinate.x--
	}
	for coordinate.y >= bottomY {
		coordinate.y--
	}
}
