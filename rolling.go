// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/13/2021

package gosl

// ********************************************************************************
// RollingIndex
// ********************************************************************************

// NewRollingIndex is used to create items such as recent 3 item. But, this only keeps the list.
// So, it should be paired with an array of same size.
// After that, for each RollingIndex.Next(), it will create index from 0 to the size given.
// When List() is called, it will return the array of integer for next run.
//   ri := gosl.NewRollingIndex(3)
// 	 ri = ri.Next() // r.Curr() => 0, ri.List() => [0]
// 	 ri = ri.Next() // r.Curr() => 1, ri.List() => [0, 1]
// 	 ri = ri.Next() // r.Curr() => 2, ri.List() => [0, 1, 2] // recent 3 oldest to newest
// 	 ri = ri.Next() // r.Curr() => 0, ri.List() => [1, 2, 0]
// 	 ri = ri.Next() // r.Curr() => 1, ri.List() => [2, 0, 1]
// 	 ri = ri.Next() // r.Curr() => 2, ri.List() => [0, 1, 2]
func NewRollingIndex(size int) RollingIndex {
	return RollingIndex{
		size:    size - 1,
		curr:    -1,
		hasFull: false,
	}
}

// RollingIndex is an index that can be used along with an array or a slice to get
// most recent n items while not saving entire records.
// Example of using this can be found from gosl.RollingBuffer.
type RollingIndex struct {
	hasFull bool
	size    int
	curr    int
}

// Reset the RollingIndex
func (r RollingIndex) Reset(size int) RollingIndex {
	r.hasFull = false
	if size > 0 {
		r.size = size - 1
	}
	r.curr = -1
	return r
}

// List current index, oldest to newest
func (r RollingIndex) List() (out []int) {
	if r.hasFull {
		for i := r.curr + 1; i <= r.size; i++ {
			out = append(out, i)
		}
	}
	for i := 0; i < r.curr+1; i++ {
		out = append(out, i)
	}
	return out
}

// Curr will return current index to use
func (r RollingIndex) Curr() int {
	return r.curr
}

// Next will be used to move to next index
func (r RollingIndex) Next() RollingIndex {
	if r.curr >= r.size {
		r.curr = 0
		r.hasFull = true
		return r
	}
	r.curr += 1
	return r
}

// ********************************************************************************
// RollingBuffer is a byte buffer for debug writer
// ********************************************************************************

// NewRollingBuffer will create RollingBuffer with n records to keep.
// This is compatible with io.Writer, and can be helpful for debugger.
func NewRollingBuffer(keep int) *RollingBuffer {
	if keep < 1 {
		keep = 10
	}
	return &RollingBuffer{
		cache: make([][]byte, keep),
		index: NewRollingIndex(keep),
	}
}

// RollingBuffer is a cache that holds n records, but also compatible as a writer.
type RollingBuffer struct {
	cache   [][]byte
	index   RollingIndex
	NewLine bool
}

// Reset will resize the buffer's keep (how many recs to keep) value
func (w *RollingBuffer) Reset(keep int) {
	if len(w.cache) < keep {
		w.cache = make([][]byte, keep)
	} else {
		w.cache = w.cache[:keep]
	}
	w.index = w.index.Reset(keep)
}

// Write is for io.Writer interface
func (w *RollingBuffer) Write(p []byte) (n int, err error) {
	w.index = w.index.Next()
	w.cache[w.index.Curr()] = append(w.cache[w.index.Curr()][:0], p...)
	return len(p), nil
}

// Last will append n records to dst. When n is less than 1, it will return all
func (w *RollingBuffer) Last(dst []byte, n int) []byte {
	idxList := w.index.List()

	for id, idx := range idxList {
		last := len(w.cache[idx])
		if id < n || n <= 0 {
			dst = append(dst, w.cache[idx]...)
			if w.NewLine && last > 0 && w.cache[idx][last-1] != '\n' {
				dst = append(dst, '\n')
			}
		}
	}
	return dst
}
