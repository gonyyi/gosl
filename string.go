// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/16/2021

package gosl

// Itoa converts int to string
func Itoa(i int) (s string) {
	return string(AppendInt(make([]byte, 0, 128), i, false))
}

// Itoaf takes option (comma)
func Itoaf(i int, comma bool) (s string) {
	return string(AppendInt(make([]byte, 0, 128), i, comma))
}

// Ftoa converts float64 to string
func Ftoa(f64 float64) (s string) {
	return string(AppendFloat64(make([]byte, 0, 128), f64, 2, false))
}

// Ftoaf converts float64 with an option (decimal, comma)
func Ftoaf(f64 float64, decimal uint8, comma bool) (s string) {
	return string(AppendFloat64(make([]byte, 0, 128), f64, decimal, comma))
}

// MustAtoi converts string to integer.
// It takes a string and a fallback integer, and if conversion failed, it will return a fallback value.
func MustAtoi(s string, fallback int) int {
	out, ok := Atoi(s)
	if !ok {
		out = fallback
	}
	return out
}

// Atoi takes a string and converts to integer.
// Atoi returns integer and a boolean -- if it returns true, it converted without an issue.
func Atoi(s string) (num int, ok bool) {
	// considering negative sign,
	// available length are
	// 32bit system: 10
	// 64bit system: 20
	if sl := len(s); (IntType == 64 && (0 < sl && sl < 21)) || (IntType == 32 && (0 < sl && sl < 11)) {
		neg := false
		start := 0

		switch s[0] {
		case '-':
			neg = true
			start++
		case '+':
			start++
		case '0': // check if there's more leading zeros
			// for idx, c := range s {
			for ; start < len(s); start++ {
				if s[start] != '0' {
					break
				}
			}
		}

		n := 0
		for _, c := range s[start:] {
			// byte(and rune) to int -- for those char, it's same
			// '+' -> 43, '0' -> 48, ' ' -> 32
			// '-' -> 45, '1' -> 49
			// '.' -> 46, '9' -> 57
			if c == ',' {
				continue
			}
			c -= '0' //
			if c > 9 {
				return 0, false
			}
			n = n*10 + int(c)
		}
		if neg {
			n = -n
		}
		return n, true
	}
	return 0, false
}

// StringSplit writes to `dst` string slice,
// after process string `s` with a rune delimiter `delim`
func StringSplit(dst []string, s string, delim rune) []string {
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

// IntsJoin takes a `dst` byte slice,
// and write joined integer to it using string slice `p` and byte `delim`
func IntsJoin(dst []byte, p []int, delim byte) []byte {
	buf := make(Buf, 0, 4096)
	for i, v := range p {
		if i != 0 {
			buf = buf.WriteByte(delim)
		}
		buf = buf.WriteInt(v)
	}
	dst = append(dst, buf...)
	return dst
}

// StringsJoin takes a `dst` byte slice,
// and write joined string to it using string slice `p` and byte `delim`
func StringsJoin(dst []byte, p []string, delim byte) []byte {
	buf := make(Buf, 0, 4096)
	for i, v := range p {
		if i != 0 {
			buf = buf.WriteByte(delim)
		}
		buf = buf.WriteString(v)
	}
	dst = append(dst, buf...)
	return dst
}

// StringTrim will trim the given string and return a byte slice
func StringTrim(s string, trimLeft, trimRight bool) string {
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

// StringsFollow will take string slice and a function to modify content
func StringsFollow(s []string, f func(s string) string) {
	if f == nil {
		return
	}
	for i, line := range s {
		s[i] = f(line)
	}
}

// IntsFollow will take string slice and a function to modify content
func IntsFollow(s []int, f func(s int) int) {
	if f == nil {
		return
	}
	for i, line := range s {
		s[i] = f(line)
	}
}

// AppendStringMiddle will append a string at the middle of given targetLength.
func AppendStringMiddle(dst []byte, s string, targetLength int, overflow bool) []byte {
	if targetLength <= 0 {
		return dst
	}

	lens := len(s)
	if lens >= targetLength {
		if overflow {
			return append(dst, s...)
		}
		return append(dst, s[:targetLength]...)
	}
	d, r := (targetLength-lens)/2, (targetLength-lens)%2
	dst = AppendRepeat(dst, []byte(" "), d)
	dst = append(dst, s...)
	dst = AppendRepeat(dst, []byte(" "), d+r)
	return dst
}

// AppendStringRight will append a string at the right aligned with a given targetLength.
func AppendStringRight(dst []byte, s string, targetLength int, overflow bool) []byte {
	if targetLength <= 0 {
		return dst
	}

	lens := len(s)
	if lens >= targetLength {
		if overflow {
			return append(dst, s...)
		}
		return append(dst, s[:targetLength]...)
	}
	dst = AppendRepeat(dst, []byte(" "), targetLength-lens)
	dst = append(dst, s...)
	return dst
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

// AppendFit will trim string if longer than wanted length,
// if shorter, it will fill with the filler byte.
// This as generates a new string, it will allocation to memory
func AppendFit(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte {
	if length <= 0 {
		return dst
	}
	slen := len(s)
	if slen > length {
		if overflowMarker && length > 2 {
			return append(append(dst, s[:length-2]...), '.', '.')
		}
		return append(dst, s[:length]...)
	}
	dst = append(dst, s...)
	for i := 0; i < length-slen; i++ {
		dst = append(dst, filler)
	}
	return dst
}

