package music

import (
	"image/draw"

	"vimagination.zapto.org/limage/lcolor"
)

var (
	topLine    = lcolor.RGB{R: 255}
	middleLine = lcolor.RGB{G: 255}
	bottomLine = lcolor.RGB{B: 255}
	point      = lcolor.RGB{}
	between    = lcolor.RGB{R: 127, G: 127, B: 127}
)

func (p *Player) Visualise(im draw.Image, channels int) {
	if channels < 1 {
		return
	}
	b := im.Bounds()
	n := b.Dx()
	sounds := make([][]float32, channels)
	for i := range sounds {
		sounds[i] = make([]float32, n)
	}
	p.mu.Lock()
	p.processMusic(sounds)
	p.mu.Unlock()
	channelHeight := b.Dy() / channels
	top := 0
	middle := channelHeight >> 1
	bottom := channelHeight
	for _, channel := range sounds {
		x := b.Min.X
		last := 0
		for _, sample := range channel {
			p := int(sample*float32(channelHeight>>1)) + middle
			im.Set(x, top, topLine)
			im.Set(x, middle, middleLine)
			im.Set(x, bottom, bottomLine)
			if last < p {
				for i := last; i < p; i++ {
					im.Set(x, i, between)
				}
			} else if last > p {
				for i := last; i > p; i-- {
					im.Set(x, i, between)
				}
			}
			last = p
			im.Set(x, p, point)
			x++
		}
		top += channelHeight
		middle += channelHeight
		bottom += channelHeight
	}
}
