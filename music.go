package music

import (
	"math"
	"sort"
	"sync"

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
	Wave       Wave
	Profile    Profile
	channel    int
	Start, End uint64
}

func (s sound) Val(rate, time float64) float64 {
	return s.Profile(time/float64(s.End-s.Start)) * s.Wave(time/rate)
}

type Player struct {
	*portaudio.Stream
	sampleRate float64

	mu     sync.Mutex
	time   uint64
	sounds sounds
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

func (p *Player) Add(start, length uint64, wave func(float64) float64, profile func(float64) float64, channel int) {
	p.mu.Lock()
	p.sounds = append(p.sounds, sound{
		Wave:    wave,
		Profile: profile,
		channel: channel,
		Start:   start,
		End:     start + length,
	})
	sort.Sort(p.sounds)
	p.mu.Unlock()
}

func (p *Player) EndZeroNote(length uint64, note Note) uint64 {
	d := p.sampleRate / float64(note)
	i, f := math.Modf(float64(length) / d)
	if f == 0 {
		return length
	}
	return uint64((i + 1) * d)
}

func (p *Player) Time() uint64 {
	p.mu.Lock()
	t := p.time
	p.mu.Unlock()
	return t
}

func (p *Player) Rate() uint64 {
	return uint64(p.sampleRate)
}

func (p *Player) process(data [][]float32) {
	p.mu.Lock()
	p.time = p.processMusic(data)
	go p.update()
}

func (p *Player) processMusic(data [][]float32) uint64 {
	var time uint64
	for channel, dc := range data {
		time = p.time
		for i := range dc {
			var f float64
			for _, sound := range p.sounds {
				if sound.channel == channel {
					if sound.Start <= time {
						if sound.End > time {
							f += sound.Val(p.sampleRate, float64(time-sound.Start))
						}
					} else {
						break
					}
				}
			}
			time++
			dc[i] = float32(f)
		}
	}
	return time
}

func (p *Player) update() {
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
	p.mu.Unlock()
}
