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
