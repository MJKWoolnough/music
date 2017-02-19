package music

import (
	"math"
	"sort"

	"github.com/gordonklaus/portaudio"
)

var (
	Initialize = portaudio.Initialize
	Terminate  = portaudio.Terminate
)

type sounds []sound

func (s sounds) Len() int {
	return len(s)
}

func (s sounds) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func (s sounds) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type sound struct {
	Note       Note
	Wave       func(phase float64) float64
	Profile    func(position float64) float64
	channel    int
	Start, End uint64
}

func (s sound) Val(rate, time float64) float64 {
	_, frac := math.Modf(time * float64(s.Note) / rate)
	return s.Profile(time/float64(s.End-s.Start)) * s.Wave(frac)
}

type Player struct {
	*portaudio.Stream
	sampleRate float64
	time       uint64
	sounds     sounds
}

func New(sampleRate float64, channels int) (*Player, error) {
	p := &Player{sampleRate: sampleRate}
	var err error
	p.Stream, err = portaudio.OpenDefaultStream(0, channels, sampleRate, 0, p.process)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) process(data [][]float32) {
	for j, input := range data {
		for i := range input {
			var f, num float64
			for _, sound := range p.sounds {
				if sound.channel == j && sound.Start <= p.time {
					if sound.End > p.time {
						f += sound.Val(p.sampleRate, float64(p.time-sound.Start))
						num++
					}
				} else {
					break
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
	changed := false
	for i := 0; i < len(p.sounds); i++ {
		if p.sounds[i].End <= p.time {
			p.sounds[i] = p.sounds[len(p.sounds)-1]
			p.sounds = p.sounds[:len(p.sounds)-1]
			i--
			changed = true
		}
	}
	if changed {
		sort.Sort(p.sounds)
	}
}
