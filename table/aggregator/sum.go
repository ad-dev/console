package aggregator

import (
	"strconv"
)

func (fs *float64sum) SetPrecision(p int) Aggregator {
	fs.p = p
	return fs
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

func SumVertcally() Aggregator {
	return &float64sum{float64aggr{t: Vertical}}
}
func SumHorizontally() Aggregator {
	return &float64sum{float64aggr{t: Horizontal}}
}
