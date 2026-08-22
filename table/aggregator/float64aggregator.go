package aggregator

import (
	"strconv"
)

func (fs *float64aggr) Type() byte {
	return fs.t
}

func (fs *float64aggr) SetType(t byte) {
	fs.t = t
}

func (fs *float64aggr) SetData(d []string) {
	fs.d = d
}

func (fs *float64aggr) Data() []string {
	return fs.d
}

func (fs *float64aggr) SetPrecision(p int) *float64aggr {
	fs.p = p
	return fs
}

func (fs *float64aggr) Calc(p int) error {
	return nil
}

func (fs *float64aggr) String() string {
	return strconv.FormatFloat(fs.v, 'f', fs.p, bitSize)
}
