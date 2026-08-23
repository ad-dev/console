package aggregator

import (
	"fmt"
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

func (fs *float64aggr) Calc(p int) error {
	return nil
}

func (fs *float64aggr) String() string {
	vStr := strconv.FormatFloat(fs.v, 'f', fs.p, bitSize)
	if fs.df != "" {
		return fmt.Sprintf(fs.df, vStr)
	}
	return vStr
}
