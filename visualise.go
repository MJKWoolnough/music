package music

import (
	"image/draw"

	"github.com/MJKWoolnough/limage/lcolor"
)

var (
	topLine    = lcolor.RGB{R: 255}
	middleLine = lcolor.RGB{G: 255}
	bottomLine = lcolor.RGB{B: 255}
	point      = lcolor.RGB{}
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
		for _, sample := range channel {
			p := int(sample*float32(channelHeight>>1)) + middle
			im.Set(x, top, topLine)
			im.Set(x, middle, middleLine)
			im.Set(x, bottom, bottomLine)
			im.Set(x, p, point)
			x++
		}
		top += channelHeight
		middle += channelHeight
		bottom += channelHeight
	}
}
