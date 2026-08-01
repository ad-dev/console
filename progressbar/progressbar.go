package progressbar

import (
	"fmt"
	"math"
	"strings"
)

const (
	defaultBarLength = 5

	minPercent = 0
	maxPercent = 100

	threshold_almost = 0.5

	empty           = "░"
	below_threshold = "▒"
	above_trheshold = "▓"
	full            = "█"
)

type Bar interface {
	fmt.Stringer
	IsComplete() bool
	Set(float64) bool
	Inc(float64) bool
	Dec(float64) bool
	Percentage() float64
	Reset()
}

type bar struct {
	percent float64
	len     int
}

func (b *bar) String() string {
	l := b.len
	p := b.percent

	if l < 1 {
		l = defaultBarLength
	}
	if p > maxPercent {
		p = maxPercent
	}
	step := maxPercent / l

	fpl := b.percent / float64(step)
	pl := int(math.Floor(fpl))
	el := l - pl

	tb := ""
	if int(math.Floor(fpl)) != int(math.Ceil(fpl)) {
		tb = below_threshold
		el--
		if fpl-math.Floor(fpl) > threshold_almost {
			tb = above_trheshold
		}
	}

	if el < 0 {
		pl += el
		el = 0
	}

	return fmt.Sprintf("%s%s%s %.2f", strings.Repeat(full, pl), tb, strings.Repeat(empty, el), b.percent) + "%"
}

func (b *bar) IsComplete() bool {
	return b.percent == maxPercent
}

func (b *bar) Set(p float64) bool {
	if p >= minPercent && p <= maxPercent {
		b.percent = p
		return true
	}
	return false
}

func (b *bar) Inc(step float64) bool {
	if b.percent+step <= maxPercent {
		b.percent += step
		return true
	}
	return false
}

func (b *bar) Dec(step float64) bool {
	if b.percent-step >= minPercent {
		b.percent -= step
		return true
	}
	return false
}

func (b *bar) Percentage() float64 {
	return b.percent
}

func (b *bar) Reset() {
	b.percent = 0
}

func New(p float64, l int) Bar {
	return &bar{percent: p, len: l}
}
