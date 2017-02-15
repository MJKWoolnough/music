package music

import "github.com/gordonklaus/portaudio"

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
	Note
	Profile
	Start, End uint32
}

type Player struct {
	*portaudio.Stream
	sampleRate float64
	time       uint32
	sounds
}

func New(sampleRate float64) (*Player, error) {
	p := &Player{sampleRate: sampleRate}
	var err error
	p.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, p.process)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) process(data [][]float32) {
	for i := range data[0] {
		data[0][i] = 0
		p.time++
	}
}
