// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

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
	if len(dst) == index {
		return append(dst, p...)
	}
	return append(dst[:index], append(p, dst[index:]...)...)
}

// BytesToLower will take a byte slice, and convert to lowercase
func BytesToLower(p []byte) []byte {
	for idx, c := range p {
		if 64 < c && c < 91 { // 64=A, 90=Z
			p[idx] = c + 32
		}
	}
	return p
}

// BytesToUpper will take a byte slice, and convert to uppercase
func BytesToUpper(p []byte) []byte {
	for idx, c := range p {
		if 96 < c && c < 123 { // 97=a, 122=z
			p[idx] = c - 32
		}
	}
	return p
}

// BytesReverse will reverse the byte slice
func BytesReverse(dst []byte) []byte {
	for i, j := 0, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}

