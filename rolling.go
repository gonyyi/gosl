// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/1/2021

package gosl

// NewRollingIndex is used to create items such as recent 3 item. But, this only keeps the list.
// So, it should be paired with an array of same size. 
// After that, for each rollingIndex.Next(), it will create index from 0 to the size given.
// When List() is called, it will return the array of integer for next run.
//   ri := reqtest.NewRollingIndex(3)
// 	 ri = ri.Next() // r.Curr() => 0, ri.List() => [0]
// 	 ri = ri.Next() // r.Curr() => 1, ri.List() => [0, 1]
// 	 ri = ri.Next() // r.Curr() => 2, ri.List() => [0, 1, 2] // recent 3 oldest to newest
// 	 ri = ri.Next() // r.Curr() => 0, ri.List() => [1, 2, 0] 
// 	 ri = ri.Next() // r.Curr() => 1, ri.List() => [2, 0, 1] 
// 	 ri = ri.Next() // r.Curr() => 2, ri.List() => [0, 1, 2] 
func NewRollingIndex(size int) rollingIndex {
	return rollingIndex{
		size:    size - 1,
		curr:    -1,
		hasFull: false,
	}
}

type rollingIndex struct {
	hasFull bool
	size    int
	curr    int
}

// Reset the rollingIndex
func (r rollingIndex) Reset(size int) rollingIndex {
	r.hasFull = false
	if size > 0 {
		r.size = size - 1
	}
	r.curr = -1
	return r
}

// List current index, oldest to newest
func (r rollingIndex) List() (out []int) {
	if r.hasFull {
		for i:=r.curr+1; i <= r.size; i++ {
			out = append(out, i)
		}
	}
	for i:=0; i<r.curr+1; i++ {
		out = append(out, i)
	}
	return out
}

// Curr will return current index to use
func (r rollingIndex) Curr() int {
	return r.curr
}

// Next will be used to move to next index
func (r rollingIndex) Next() rollingIndex {
	if r.curr >= r.size {
		r.curr = 0
		r.hasFull = true
		return r
	}
	r.curr += 1
	return r
}
