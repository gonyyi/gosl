// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/16/2021

package gosl

// IsNumber will take a string and check if it's a number
func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if v < 48 || v > 57 {
			return false
		}
	}
	return true
}

// Splits writes to `dst` string slice,
// after process string `s` with a rune delimiter `delim`
func Splits(dst []string, s string, delim rune) []string {
	last := 0
	for idx, v := range s {
		if v == delim {
			if last != idx {
				dst = append(dst, s[last:idx])
			}
			last = idx + 1 // next line
		}
	}
	if last != len(s) {
		dst = append(dst, s[last:])
	}
	return dst
}

// Joins takes a `dst` byte slice,
// and write joined string to it using string slice `p` and byte `delim`
func Joins(dst []byte, p []string, delim byte) []byte {
	buf := make(Buf, 0, 4096)
	for i, v := range p {
		if i != 0 {
			buf = buf.WriteBytes(delim)
		}
		buf = buf.WriteString(v)
	}
	dst = append(dst, buf...)
	return dst
}

// Trims will trim the given string and return a byte slice
func Trims(s string, trimLeft, trimRight bool) string {
	// check first non space
	first := 0
	last := len(s) // this possibly can be an issue with unicode..

	if trimLeft {
		for ; first < len(s); first++ {
			if s[first] != ' ' {
				break
			}
		}
	}
	if trimRight {
		// now check from backward
		for ; last > first; last-- {
			if s[last-1] != ' ' {
				break
			}
		}
	}
	return s[first:last]
}

// HasPrefix will return true if has a prefix
func HasPrefix(s string, prefix string) bool {
	sLen := len(s)
	pfxLen := len(prefix)
	// Prefix can't be longer
	if pfxLen > sLen {
		return false
	}
	// If same length, it should be identical
	if pfxLen == sLen && s != prefix {
		return false
	}
	if s[:pfxLen] != prefix {
		return false
	}
	return true
}

// HasSuffix will return true if has a prefix
func HasSuffix(s string, suffix string) bool {
	sLen := len(s)
	sfxLen := len(suffix)
	// Prefix can't be longer
	if sfxLen > sLen {
		return false
	}
	// If same length, it should be identical
	if sfxLen == sLen && s != suffix {
		return false
	}
	if s[sLen-sfxLen:] != suffix {
		return false
	}
	return true
}

// TrimPrefix will trim prefix
func TrimPrefix(s string, prefix string) string {
	sLen := len(s)
	pfxLen := len(prefix)
	// Prefix can't be longer
	if pfxLen > sLen {
		return s
	}
	// If same length, it should be identical
	if pfxLen == sLen && s != prefix {
		return s
	}
	if s[:pfxLen] != prefix {
		return s
	}
	return s[pfxLen:]
}

// TrimSuffix will trim prefix
func TrimSuffix(s string, suffix string) string {
	sLen := len(s)
	sfxLen := len(suffix)
	// Prefix can't be longer
	if sfxLen > sLen {
		return s
	}
	// If same length, it should be identical
	if sfxLen == sLen && s != suffix {
		return s
	}
	if s[sLen-sfxLen:] != suffix {
		return s
	}
	return s[:sLen-sfxLen]
}
