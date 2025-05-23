package main

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"io"
	"os"
	"runtime"
)

//go:embed assets/*.wav
var soundAssets embed.FS

type sound struct {
	sampleRate   int
	cxt          *audio.Context
	soundEffects soundEffects
	enabled      bool
}

type soundEffects struct {
	click *audio.Player
	win   *audio.Player
	lose  *audio.Player
	flag  *audio.Player
}

func newSound() sound {
	enabled := true
	if runtime.GOOS == "linux" {
		// Check for WSL environment variables
		if _, present := os.LookupEnv("WSL_DISTRO_NAME"); present {
			enabled = false
			println("Sound disabled: Running under WSL")
		} else if _, present := os.LookupEnv("WSLENV"); present {
			enabled = false
			println("Sound disabled: Running under WSL")
		}
	}
	return sound{
		sampleRate: 44100,
		cxt:        nil,
		enabled:    enabled,
	}
}

func (s *sound) loadWav(path string) (*audio.Player, error) {
	if !s.enabled || s.cxt == nil {
		return nil, nil
	}
	f, err := soundAssets.Open("assets/" + path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewReader(data)

	d, err := wav.DecodeWithoutResampling(buffer)
	if err != nil {
		return nil, err
	}

	p, err := s.cxt.NewPlayer(d)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *sound) init() {
	if !s.enabled {
		return
	}

	if s.cxt == nil {
		s.cxt = audio.NewContext(s.sampleRate)
	}
	s.soundEffects.click, _ = s.loadWav("click2.wav")
	s.soundEffects.win, _ = s.loadWav("win.wav")
	s.soundEffects.lose, _ = s.loadWav("explosion.wav")
	s.soundEffects.flag, _ = s.loadWav("flag.wav")
}
