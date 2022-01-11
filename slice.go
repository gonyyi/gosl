// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// slice.go (SortAny, SortStrings, SortInts)
// This is to do what sort.SortInts, sort.SortStrings, sort.Slice does,
// however, without importing any libraries at all (including standard library).
// Note that this is just quick test and is slower than built in sort library.

// SortAny is designed to sort any slice with no memory allocation.
// The usage is bit different than Go's `sort.Slice()` function.
// - pSize: size of slice that needs to be sorted
// - swap: a function that will swap the slice
// - less: a function that will return true when index i of slice is less than j's.
//
// Example:
// 	SortAny(
//		len(a), // size of the slice
//		func(i1, i2 int) { a[i1], a[i2] = a[i2], a[i1] },       // swap
//		func(i, j int) bool { return a[i].Score > a[j].Score }, // less
//	)
func SortAny(pSize int, swap func(i, j int), less func(i, j int) bool) {
	// This function requires both swap and less function.
	if pSize < 2 || swap == nil || less == nil {
		return
	}
	for {
		changed := false
		for i := 0; i < pSize-1; i++ {
			if !less(i, i+1) { // 2nd one is bigger
				if !less(i+1, i) {
					continue // equal
				}
				// swap
				swap(i, i+1)
				changed = true
			}
		}
		// Nothing changed, no more need to run
		if changed == false {
			break
		}
	}
}

// SortStrings will sort []string slice, if compare function is not given,
// it will default to alphabetical
func SortStrings(dst []string, compare func(idx1, idx2 int) bool) (ok bool) {
	// if compare func is not exist, or invalid, return false
	if dst == nil {
		return false
	}
	if compare == nil {
		compare = func(idx1, idx2 int) bool {
			return dst[idx1] < dst[idx2]
		}
	}

	maxAny := len(dst) - 1
	for {
		changed := false
		for i := 0; i < maxAny; i++ {
			if !compare(i, i+1) { // 2nd one is bigger
				if !compare(i+1, i) {
					continue // equal
				}
				// swap
				dst[i], dst[i+1] = dst[i+1], dst[i]
				changed = true
			}
		}
		if changed == false {
			break
		}
	}
	return true
}

// SortInts will sort []int slice low to high
func SortInts(dst []int) {
	// if compare func is not exist, or invalid, return false
	if dst == nil {
		return
	}

	maxAny := len(dst) - 1

	for {
		changed := false
		for i := 0; i < maxAny; i++ {
			// if i+1 < maxAny {
			if !(dst[i] < dst[i+1]) { // 2nd one is bigger
				if !(dst[i+1] < dst[i]) {
					continue // equal
				}
				// swap
				dst[i], dst[i+1] = dst[i+1], dst[i]
				changed = true
			}
			// }
		}
		if changed == false {
			break
		}
	}
	return
}

// DedupStrings will deduplicate string slice
// First, it will sort, then exam from the left to make sure every values are different from previous.
// NOTE: during the dedup process, this will alter original slice -- this will sort it.
func DedupStrings(p []string) []string {
	// when there's one or no element, no need to dedup
	if len(p) < 2 {
		return p
	}
	// sort the string
	SortStrings(p, nil)

	// starting with the 2nd element as 1st will be the baseline.
	// whenever new one's found, set its index to cur; and return up to cur
	cur := 0
	for i := 1; i < len(p); i++ {
		// if element is same as previous, skip
		if p[i-1] == p[i] {
			continue
		}
		cur += 1
		p[cur] = p[i]
	}
	return p[:cur+1]
}

// DedupInts will deduplicate int slice
// First, it will sort, then exam from the left to make sure every values are different from previous.
// NOTE: during the dedup process, this will alter original slice -- this will sort it.
func DedupInts(p []int) []int {
	// when there's one or no element, no need to dedup
	if len(p) < 2 {
		return p
	}
	// sort the string
	SortInts(p)

	// starting with the 2nd element as 1st will be the baseline.
	// whenever new one's found, set its index to cur; and return up to cur
	cur := 0
	for i := 1; i < len(p); i++ {
		// if element is same as previous, skip
		if p[i-1] == p[i] {
			continue
		}
		cur += 1
		p[cur] = p[i]
	}
	return p[:cur+1]
}
