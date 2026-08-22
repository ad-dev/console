package aggregator

import (
	"math"
	"strconv"
)

func (fs *float64min) SetPrecision(p int) Aggregator {
	fs.p = p
	return fs
}

func (fs *float64min) Calc() error {
	var min float64 = math.MaxFloat64
	for _, v := range fs.d {
		fv, err := strconv.ParseFloat(v, bitSize)
		if err != nil {
			return err
		}
		if fv < min {
			min = fv
		}
		fs.v = min
	}
	return nil
}

func MinVertcally() Aggregator {
	return &float64min{float64aggr{t: Vertical}}
}
func MinHorizontally() Aggregator {
	return &float64min{float64aggr{t: Horizontal}}
}
