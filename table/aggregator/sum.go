package aggregator

import (
	"strconv"
	"strings"
)

func (fs *float64sum) SetPrecision(p int) Aggregator {
	fs.p = p
	return fs
}

func (fs *float64sum) Calc() error {
	for _, v := range fs.d {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		parts := strings.Split(v, "\n")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			fv, err := strconv.ParseFloat(part, bitSize)
			if err != nil {
				return err
			}
			fs.v += fv

		}
	}
	return nil
}

func (fs *float64sum) SetDisplayFormat(f string) Aggregator {
	fs.df = f
	return fs
}

func SumVertcally() Aggregator {
	return &float64sum{float64aggr{t: Vertical}}
}
func SumHorizontally() Aggregator {
	return &float64sum{float64aggr{t: Horizontal}}
}
