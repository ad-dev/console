package aggregator

import (
	"math"
	"strconv"
)

func (fs *float64max) SetPrecision(p int) Aggregator {
	fs.p = p
	return fs
}

func (fs *float64max) Calc() error {
	var min float64 = -math.MaxFloat64
	for _, v := range fs.d {
		fv, err := strconv.ParseFloat(v, bitSize)
		if err != nil {
			return err
		}
		if fv > min {
			min = fv
		}
		fs.v = min
	}
	return nil
}

func MaxVertcally() Aggregator {
	return &float64max{float64aggr{t: Vertical}}
}
func MaxHorizontally() Aggregator {
	return &float64max{float64aggr{t: Horizontal}}
}
