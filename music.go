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
	Wave       func(int64, int64) int64
	Profile    func(int64, int64) int64
	Start, End int64
}

func (s sound) Val(rate, time uint64) uint64 {
	return s.Profile(s.Wave(math.Mod(time * float64(s.Note) / rate)))
}

type Player struct {
	*portaudio.Stream
	sampleRate int64
	time       int64
	sounds
}

func New(sampleRate int64) (*Player, error) {
	p := &Player{sampleRate: sampleRate}
	var err error
	p.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, p.process)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) process(data [][]int32) {
	for i := range data[0] {
		var f, num int64
		for _, sound := range p.sounds {
			if sound.Start <= p.time && sound.End > p.time {
				f += p.sounds[j].Val(p.sampleRate, p.time)
				num++
			}
		}
		if num > 0 {
			data[0][i] = int32(f / num)
		} else {
			data[0][i] = 0
		}
		p.time++
	}
}
