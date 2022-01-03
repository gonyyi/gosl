// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// sort.go (SortSlice, SortStrings, SortInts)
// This is to do what sort.SortInts, sort.SortStrings, sort.Slice does,
// however, without importing any libraries at all (including standard library).
// Note that this is just quick test and is slower than built in sort library.

// SortSlice will take an interface slice, and compare function;
// based on the compare function, it will sort the slice.
// If invalid param(s) given, it will return false.
func SortSlice(dst []interface{}, compare func(idx1, idx2 int) bool) (ok bool) {
	// if compare func is not exist, or invalid, return false
	if dst == nil || compare == nil || compare(1, 2) && compare(2, 1) {
		return false
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

// SortStrings will sort []string slice, if compare function is not given,
// it will default to alphabetical
func SortStrings(dst []string, compare func(idx1, idx2 int) bool) (ok bool) {
	// if compare func is not exist, or invalid, return false
	if dst == nil || (compare != nil && compare(1, 2) && compare(2, 1)) {
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
