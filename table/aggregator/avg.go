package aggregator

import (
	"strconv"
)

func (fs *float64avg) Calc() error {
	var vs float64
	c := 0
	for _, v := range fs.d {
		fv, err := strconv.ParseFloat(v, bitSize)
		if err != nil {
			return err
		}
		vs += fv
		c++
	}
	if c != 0 {
		fs.v = vs / float64(c)
	}
	return nil
}

func (fs *float64avg) SetPrecision(p int) Aggregator {
	fs.p = p
	return fs
}

func AvgVertcally() Aggregator {
	return &float64avg{float64aggr{t: Vertical}}
}
func AvgHorizontally() Aggregator {
	return &float64avg{float64aggr{t: Horizontal}}
}
