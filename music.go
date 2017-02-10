package music

import "github.com/gordonklaus/portaudio"

var (
	Initialize = portaudio.Initialize
	Terminate  = portaudio.Terminate
)

type Wave interface {
	next() float64
}

type note struct {
	Wave
	Profile
}

func (n *note) next() (float64, bool) {
	p, d := n.Profile.next()
	return n.Wave.next() * p, d
}

type Player struct {
	*portaudio.Stream
	sampleRate float64
	notes      []note
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

func (p *Player) AddSine(n Note, pr Profile) {
	p.notes = append(p.notes, note{
		NewSine(n, p.sampleRate),
		pr,
	})
}

func (p *Player) AddSaw(n Note, pr Profile) {
	p.notes = append(p.notes, note{
		NewSaw(n, p.sampleRate),
		pr,
	})
}

func (p *Player) AddTriangle(n Note, pr Profile) {
	p.notes = append(p.notes, note{
		NewTriangle(n, p.sampleRate),
		pr,
	})
}

func (p *Player) AddSquare(n Note, pr Profile) {
	p.notes = append(p.notes, note{
		NewSquare(n, p.sampleRate),
		pr,
	})
}

func (p *Player) process(data [][]float32) {
	for i := range data[0] {
		data[0][i] = 0
		l := float64(len(p.notes))
		for j := 0; j < len(p.notes); j++ {
			s, d := p.notes[j].next()
			data[0][i] += float32(s / l)
			if d {
				p.notes[j] = p.notes[len(p.notes)-1]
				p.notes = p.notes[:len(p.notes)-1]
				j--
			}
		}
	}
}
