package float64

import (
	"fmt"
	"math/rand"
	"time"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type Float64Array []float64

func (p Float64Array) Len() int           { return len(p) }
func (p Float64Array) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Array) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewFloat64Array() Float64Array {
	return make([]float64, 25)
}

func (p Float64Array) Fill(n int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < n; i++ {
		p[i] = 100 * (rand.Float64())
	}
}

func (p Float64Array) List() string {
	s := "{ "
	for i := 0; i < p.Len(); i++ {
		if p[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%3.1f ", p[i])
	}
	s += " }"
	return s
}

func (p Float64Array) String() string {
	return p.List()
}
