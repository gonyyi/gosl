// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/14/2021

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

// Split writes to `dst` string slice,
// after process string `s` with a rune delimiter `delim`
func Split(dst []string, s string, delim rune) []string {
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

// LastSplit will split the string, and return the last item
// eg. LastSplit("/123/456/abc") => "abc"
func LastSplit(s string, delim rune) string {
	idx := -1
	for i, v := range s {
		if v == delim {
			idx = i
			continue
		}
	}

	// delim not found => return all
	if idx == -1 {
		return s
	}

	// normal
	if len(s) > idx+1 {
		return s[idx+1:]
	}

	// ending with the delim
	return ""
}

// Join takes a `dst` byte slice,
// and write joined string to it using string slice `p` and byte `delim`
func Join(dst []byte, p []string, delim ...byte) []byte {
	buf := make(Buf, 0, 4096)
	for i, v := range p {
		if i != 0 {
			if delim == nil {
				buf = buf.WriteBytes(',')
			} else {
				buf = buf.WriteBytes(delim...)
			}
		}
		buf = buf.WriteString(v)
	}
	dst = append(dst, buf...)
	return dst
}

// Trim will trim left and right
func Trim(s string) string {
	return trim(s, true, true)
}

// TrimLeft will trim left and right
func TrimLeft(s string) string {
	return trim(s, true, false)
}

// TrimRight will trim left and right
func TrimRight(s string) string {
	return trim(s, false, true)
}

// trim will trim the given string and return a byte slice
func trim(s string, trimLeft, trimRight bool) string {
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

// FirstN will return first n character of string.
// if n is larger than the length of string, it will return whatever available.
func FirstN(s string, n int) string {
	if n < 0 {
		n = 0
	}
	if len(s) <= n {
		return s
	}
	return s[0:n]
}

// LastN will return last n character of string.
// if n is larger than the length of string, it will return whatever available.
func LastN(s string, n int) string {
	if n < 0 {
		n = 0
	}
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}

// Mask will take a string and mask except for first and last n bytes.
// This will be used to mask credentials.
func Mask(s string, firstN, lastN int) string {
	buf := GetBuffer()
	defer buf.Free()
	buf.Buffer = AppendStringMask(buf.Buffer, s, firstN, lastN)
	return buf.String()
}

// AddLinePrefix will take src and line prefix, and return a string.
// For each new line, it will add prefix
func AddLinePrefix(s string, linePrefix string) string {
	buf := make([]byte, 0, len(s)+len(s)/3) // guess some size..
	buf = AppendStringLinePrefix(buf, s, linePrefix)
	return string(buf)
}
