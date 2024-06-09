package main

func displayFrame() {
	frameOriginX, frameOriginY := getFrameOrigin()
	printUnfilledRectangle(frameOriginX, frameOriginY, FRAME_WIDTH, FRAME_HEIGHT, FRAME_BORDER_THICKNESS, FRAME_BORDER_HORIZONTAL, FRAME_BORDER_VERTICAL, FRAME_BORDER_TOP_LEFT, FRAME_BORDER_TOP_RIGHT, FRAME_BORDER_BOTTOM_RIGHT, FRAME_BORDER_BOTTOM_LEFT)
	Screen.Show()
}

func getFrameOrigin() (int, int) {
	return (screenWidth-FRAME_WIDTH)/2 - FRAME_BORDER_THICKNESS, (screenHeight-FRAME_HEIGHT)/2 - FRAME_BORDER_THICKNESS
}

func getBoundaries() (int, int, int, int) {
	originX, originY := getFrameOrigin()
	topY := originY
	bottomY := originY + FRAME_HEIGHT - FRAME_BORDER_THICKNESS
	leftX := originX
	rightX := originX + FRAME_WIDTH - FRAME_BORDER_THICKNESS
	return leftX, topY, rightX, bottomY
}
