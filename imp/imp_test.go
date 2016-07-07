package priorityQ

import (
	"fmt"
	S2 "github.com/Sreevani871/priorityQ/imp"
	"testing"
)

type Test struct {
	values      []int
	remele      int
	empty, full bool
}

var test = []Test{
	{values: []int{1, 2, 1, 7, 23}, remele: 1, empty: false, full: false},
	{values: []int{6, 0, 3, 1, -90, 67}, remele: -90, empty: false, full: false},
	{values: []int{}, remele: -1, empty: true, full: false},
}

func TestEnqueue(t *testing.T) {
	var e int

	P := S2.New()
	for c, _ := range test[1].values {
		e = P.Enqueue(c)
		if e != c {
			t.Error(
				"Inserting", c,
				"Expected", c,
				"Inserted", e,
			)

		}

	}
}

func TestDequeue(t *testing.T) {
	for _, v := range test {
		P := S2.New()
		fmt.Println("check", v)
		for _, c := range v.values {

			fmt.Println(c)
			P.Enqueue(c)
			fmt.Println("priority", P.PrintPQ())
		}
		d := P.Dequeue()
		if d != v.remele {
			t.Error(
				"For", v.values,
				"Expected", v.remele,
				"Got", d,
			)

		}

	}
}

func TestIsEmpty(t *testing.T) {
	for _, v := range test {
		P := S2.New()
		fmt.Println("check", v)
		for _, c := range v.values {

			fmt.Println(c)
			P.Enqueue(c)
			fmt.Println("priority", P.PrintPQ())
		}
		d := P.IsEmpty()
		if d != v.empty {
			t.Error(
				"For", v.values,
				"Expected", v.empty,
				"Got", d,
			)

		}

	}

}

func TestIsFull(t *testing.T) {
	for _, v := range test {
		P := S2.New()
		fmt.Println("check", v)
		for _, c := range v.values {

			fmt.Println(c)
			P.Enqueue(c)
			fmt.Println("priority", P.PrintPQ())
		}
		d := P.IsFull()
		fmt.Println(d)
		if d != v.full {
			t.Error(
				"For", v.values,
				"Expected", v.full,
				"Got", d,
			)

		}

	}

}

func TestMerge(t *testing.T) {
	mergeres := []int{67, 23, 7, 6, 3, 2, 1, 1, 1, 0, -90}
	P1 := S2.New()
	for _, c := range test[0].values {
		P1.Enqueue(c)
	}
	fmt.Println("P", P1.IntPq)
	N1 := S2.New()
	for _, c := range test[1].values {
		N1.Enqueue(c)
	}
	fmt.Println("N", N1.IntPq)
	P1.Merge(N1)
	res := P1.PrintPQ()
	fmt.Println(res)
	for v, i := range res {
		if res[v] != i {
			t.Error(
				"For", test[0].values, test[1].values,
				"expected", mergeres,
				"Got", res,
			)
		}
	}

}
