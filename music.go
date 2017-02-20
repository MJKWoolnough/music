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
	channels   int
}

func New(sampleRate float64, channels int) (*Player, error) {
	p := &Player{sampleRate: sampleRate, channels: channels}
	var err error
	p.Stream, err = portaudio.OpenDefaultStream(0, channels, sampleRate, 0, p.process)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) Add(start, length uint64, note Note, wave func(float64) float64, profile func(float64) float64, channel int) {
	p.sounds = append(p.sounds, sound{
		Note:    note,
		Wave:    wave,
		Profile: profile,
		channel: channel,
		Start:   start,
		End:     start + length,
	})
}

func (p *Player) EndZeroNote(length uint64, note Note) uint64 {
	d := p.sampleRate / float64(note)
	i, f := math.Modf(float64(length) / d)
	if f == 0 {
		return length
	}
	return uint64((i + 1) * d)
}

func (p *Player) process(data []float32) {
	c := 0
	for i := range data {
		c++
		if c == p.channels {
			c = 0
			p.time++
		}
		var f, num float64
		for _, sound := range p.sounds {
			if sound.channel == c {
				if sound.Start <= p.time {
					if sound.End > p.time {
						f += sound.Val(p.sampleRate, float64(p.time-sound.Start))
						num++
					}
				} else {
					break
				}
			}
		}
		if num > 0 {
			data[i] = float32(f / num)
		} else {
			data[i] = 0
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
