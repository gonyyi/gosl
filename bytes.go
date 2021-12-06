// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/06/2021

package gosl

// NewBytesFilter will return a function that takes a byte slice,
// and returns filtered results.
// IF `allow` param is set to true, newly created function
// will only take what's in the `list` param.
// IF `allow` param is false, it will return everything except what's
// in the `list`.
// Usage:
//     clean := NewBytesFilter(true, []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"])
//     ...
//     userInput := "this has some system; echo ${password}"
//     newInput := string( clean([]byte(userInput) )
func NewBytesFilter(allow bool, list []byte) func([]byte) []byte {
	// byte is uint8, that has 0-255
	var allowed [256]bool

	// IF allow == false, then listed list are the only
	// one function should accept.
	// IF allow == true, since default bool will have false,
	// it should be fine.
	if allow == false {
		for i := 0; i < 256; i++ {
			allowed[i] = true
		}
	}

	for _, c := range list {
		allowed[int(c)] = allow
	}

	return func(s []byte) []byte {
		cur := 0
		for i := 0; i < len(s); i++ {
			if c := s[i]; allowed[c] == true {
				s[cur] = c
				cur++
			}
		}
		return s[:cur]
	}
}

// BytesInsert will take a byte slice, and append a byte slice at the given position.
func BytesInsert(dst []byte, index int, p []byte) []byte {
	if len(dst) <= index {
		return append(dst, p...)
	}
	return append(dst[:index], append(p, dst[index:]...)...)
}

// BytesReverse will reverse the byte slice
func BytesReverse(dst []byte) []byte {
	for i, j := 0, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}


// BytesToUpper converts the slice into uppercase
func BytesToUpper(dst []byte) []byte {
	// A:65 Z:90 a:97 z:122
	for i := 0; i < len(dst); i++ {
		if c := dst[i]; 'a' <= c && c <= 'z' {
			dst[i] = c - 32
		}
	}
	return dst
}

// BytesToLower converts the slice into lowercase
func BytesToLower(dst []byte) []byte {
	// A:65 Z:90 a:97 z:122
	for i := 0; i < len(dst); i++ {
		if c := dst[i]; 'A' <= c && c <= 'Z' {
			dst[i] = c + 32
		}
	}
	return dst
}

// BytesEqual takes two byte slices and check if they are equal
func BytesEqual(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i:=0; i<len(b2); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
