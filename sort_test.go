// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Sort_SortInts(t *testing.T) {
	a := []int{1, 5, 2, 4, 91, 3}
	gosl.SortInts(a)
	gosl.Test(t, "[1 2 3 4 5 91]", fmt.Sprint(a))

}

func Benchmark_Sort_SortInts(b *testing.B) {
	b.ReportAllocs()
	a := []int{1, 5, 2, 4, 91, 3}
	for i := 0; i < b.N; i++ {
		a[0] = 1
		a[1] = 5
		a[2] = 2
		a[3] = 4
		a[4] = 91
		a[5] = 3
		gosl.SortInts(a)
	}
	// fmt.Println(a)
}

func Test_Sort_SortStrings(t *testing.T) {
	t.Run("compare=nil", func(t *testing.T) {
		a := []string{"abc", "def", "b", "c"}
		gosl.SortStrings(a, nil)
		gosl.Test(t, "[abc b c def]", fmt.Sprint(a))
	})

	t.Run("compare=reverse", func(t *testing.T) {
		a := []string{"abc", "def", "b", "c"}
		gosl.SortStrings(a, func(idx1, idx2 int) bool {
			return a[idx1] > a[idx2]
		})
		gosl.Test(t, "[def c b abc]", fmt.Sprint(a))
	})
}

func Benchmark_Sort_SortStrings(b *testing.B) {
	b.ReportAllocs()
	a := []string{"abc", "def", "b", "c"}
	for i := 0; i < b.N; i++ {
		a[0] = "abc"
		a[1] = "def"
		a[2] = "b"
		a[3] = "c"
		gosl.SortStrings(a, func(idx1, idx2 int) bool {
			return a[idx1] > a[idx2]
		})
	}
	// fmt.Println(a)
}

func Test_Sort_SortSlice(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		a := []interface{}{1, 5, 2, 4, 91, 3}
		gosl.SortSlice(a, func(idx1, idx2 int) bool {
			return a[idx1].(int) < a[idx2].(int)
		})

		gosl.Test(t, "[1 2 3 4 5 91]", fmt.Sprintf("%v", a))
	})

	t.Run("int2", func(t *testing.T) {
		a := []interface{}{1, 5, 2, 4, 91, 3}
		gosl.SortSlice(a, func(idx1, idx2 int) bool {
			return a[idx1].(int) > a[idx2].(int)
		})

		gosl.Test(t, "[91 5 4 3 2 1]", fmt.Sprintf("%v", a))
	})

	t.Run("string1", func(t *testing.T) {
		a := []interface{}{"abc", "a", "b", "c", "def", "d"}
		gosl.SortSlice(a, func(idx1, idx2 int) bool {
			return a[idx1].(string) < a[idx2].(string)
		})

		gosl.Test(t, "[a abc b c d def]", fmt.Sprintf("%v", a))
	})
	t.Run("string2", func(t *testing.T) {
		type ID struct {
			name string
			age  int
		}

		a := []interface{}{
			ID{name: "AGON YI", age: 13},
			ID{name: "BJOHN YI", age: 13},
			ID{name: "AGON YI", age: 11},
			ID{name: "BZON YI", age: 11},
		}
		gosl.SortSlice(a, func(idx1, idx2 int) bool {
			// age low to high, then name high to low
			if a[idx1].(ID).age < a[idx2].(ID).age {
				return true
			}
			if a[idx1].(ID).age > a[idx2].(ID).age {
				return false
			}
			return a[idx1].(ID).name > a[idx2].(ID).name
		})
		gosl.Test(t, "[{BZON YI 11} {AGON YI 11} {BJOHN YI 13} {AGON YI 13}]", fmt.Sprintf("%v", a))
	})
}

