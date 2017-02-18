package music

import (
	"math"

	"github.com/gordonklaus/portaudio"
)

var (
	Initialize = portaudio.Initialize
	Terminate  = portaudio.Terminate
)

type sounds []sound

func (s sounds) Len() int {
	return len(*s)
}

func (s sounds) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func (s sounds) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type sound struct {
	Note       Note
	Wave       func(float64, float64) float64
	Profile    func(float64, float64) float64
	Start, End int64
}

func (s sound) Val(rate, time float64) float64 {
	return s.Profile(s.Wave(math.Mod(time * float64(s.Note) / rate)))
}

type Player struct {
	*portaudio.Stream
	sampleRate float64
	time       uint64
	channels   []sounds
}

func New(sampleRate int64, channels int) (*Player, error) {
	p := &Player{sampleRate: sampleRate, channels: make([]sounds, channels)}
	var err error
	p.Stream, err = portaudio.OpenDefaultStream(0, channels, sampleRate, 0, p.process)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) process(data [][]float32) {
	for j, input := range data {
		channel = p.channels[j]
		for i := range input {
			var f, num float64
			for _, sound := range channel {
				if sound.Start <= p.time && sound.End > p.time {
					f += p.sounds[j].Val(p.sampleRate, p.time)
					num++
				}
			}
			if num > 0 {
				input[i] = float32(f / num)
			} else {
				input[i] = 0
			}
			p.time++
		}
	}
}
