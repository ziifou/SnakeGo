package main

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var eatSound, gameOverSound, pauseSound beep.StreamSeekCloser
var format beep.Format

func initSounds() {
	eatSound = loadSound("eat.mp3")
	gameOverSound = loadSound("gameover.mp3")
	pauseSound = loadSound("pause.mp3")
}

func loadSound(filename string) beep.StreamSeekCloser {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	streamer, retournedFormat, err := mp3.Decode(f)

	format = retournedFormat
	if err != nil {
		panic(err)
	}

	return streamer
}

func playSound(sound beep.StreamSeekCloser) {
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(sound, beep.Callback(func() {
		sound.Seek(0)
	})))
}
