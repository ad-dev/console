package aggregator

import (
	"strconv"
)

const (
	bitSize = 64
)

type float64sum struct {
	t byte
	v float64
	p int
	d []string
}

func (fs *float64sum) Type() byte {
	return fs.t
}

func (fs *float64sum) SetType(t byte) {
	fs.t = t
}

func (fs *float64sum) SetData(d []string) {
	fs.d = d
}

func (fs *float64sum) SetPrecision(p int) Aggregator {
	fs.p = p
	return fs
}

func (fs *float64sum) Data() []string {
	return fs.d
}

func (fs *float64sum) Calc() error {
	for _, v := range fs.d {
		fv, err := strconv.ParseFloat(v, bitSize)
		if err != nil {
			return err
		}
		fs.v += fv
	}
	return nil
}

func (fs *float64sum) String() string {
	return strconv.FormatFloat(fs.v, 'f', fs.p, bitSize)
}

func SumVertcally() Aggregator {
	return &float64sum{t: Vertical}
}
func SumHorizontally() Aggregator {
	return &float64sum{t: Horizontal}
}
