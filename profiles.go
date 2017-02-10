package music

import "fmt"

type Profile interface {
	next() (float64, bool)
}

type Profiles []Profile

func (p *Profiles) next() (float64, bool) {
	if len(*p) == 0 {
		return 0, true
	}
	v, e := (*p)[0].next()
	if e {
		(*p) = (*p)[1:]
		return v, len(*p) == 0
	}
	return v, false
}

type Maintain struct {
	Duration uint64
}

func (m *Maintain) next() (float64, bool) {
	m.Duration--
	return 1, m.Duration == 0
}

type Adjust struct {
	Profile
	Adjustment float64
}

func (a Adjust) next() (float64, bool) {
	v, e := a.Profile.next()
	return v * a.Adjustment, e
}

type Debug struct {
	Profile
}

func (d Debug) next() (float64, bool) {
	v, e := d.Profile.next()
	fmt.Println(v, e)
	return v, e
}

type Ramp struct {
	Amount float64
	count  float64
}

func (r *Ramp) next() (float64, bool) {
	r.count += r.Amount
	if r.count >= 0 {
		if r.count >= 1 {
			return 1, true
		}
		return r.count, false
	}
	if r.count <= -1 {
		return 0, true
	}
	return 1 + r.count, false
}
