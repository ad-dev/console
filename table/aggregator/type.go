package aggregator

import "fmt"

const (
	Vertical = iota
	Horizontal
)

type Aggregator interface {
	fmt.Stringer
	Type() byte
	SetType(byte)
	SetPrecision(int) Aggregator
	SetData([]string)
	Data() []string
	Calc() error
}
type float64aggr struct {
	t byte
	v float64
	p int
	d []string
}

type float64sum struct {
	float64aggr
}

type float64avg struct {
	float64aggr
}

type float64min struct {
	float64aggr
}

type float64max struct {
	float64aggr
}
