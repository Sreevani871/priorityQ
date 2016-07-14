package priorityQ

import (
	S2 "github.com/Sreevani871/priorityQ/imp"
	"testing"
)

type Test struct {
	values []int
	remele int
	empty  bool
	full   bool
}

var test = []Test{
	{values: []int{1, 2, 1, 7, 23}, remele: 1, empty: false, full: false},
	{values: []int{6, 0, 3, 1, -90, 67}, remele: -90, empty: false, full: false},
	{values: []int{}, remele: -1, empty: true, full: false},
}

func BenchmarkEnqueue(b *testing.B) {
	b.SetBytes(2)
	P := S2.New()
	for i := 0; i < b.N; i++ {
		for c, _ := range test[1].values {
			P.Enqueue(c)

		}
	}
}

func BenchmarkDequeue(b *testing.B) {
	b.SetBytes(2)
	P := S2.New()
	for i := 0; i < b.N; i++ {
		P.Dequeue()
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	b.SetBytes(2)
	P := S2.New()
	for c, _ := range test[1].values {
		P.Enqueue(c)

	}
	for i := 0; i < b.N; i++ {
		P.IsEmpty()
	}

}

func BenchmarkIsFull(b *testing.B) {
	b.SetBytes(2)
	P := S2.New()
	for c, _ := range test[1].values {
		P.Enqueue(c)

	}
	for i := 0; i < b.N; i++ {
		P.IsFull()
	}

}
func BenchmarkMerge(b *testing.B) {
	b.SetBytes(2)
	P := S2.New()
	for c, _ := range test[1].values {
		P.Enqueue(c)

	}
	P1 := S2.New()
	for c, _ := range test[2].values {
		P1.Enqueue(c)

	}
	for i := 0; i < b.N; i++ {
		P.Merge(P1)
	}

}
