// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// Atoi takes a string and converts to integer.
// Atoi returns integer and a boolean -- if it returns true, it converted without an issue.
// Note:
//   -9223372036854775808 is lowest possible number.
//   However, this will be flagged as invalid number because during the calculation,
//   it calculates all in positive number, then flip the sign at the end to speed up.
//   (To be able to give correct output for -9223372036854775808, either code needs to be
//    longer (positive, and negative separate logic), or check the sign for each digit.)
func Atoi(s string) (num int, ok bool) {
	// considering negative sign,
	// available length are
	// 32bit system: 10
	// 64bit system: 19
	if len(s) == 0 {
		return 0, false
	}

	neg := false
	start := 0

	switch s[0] {
	case '-':
		neg = true
		start++
		// if input was `-`, it shouldn't be considered as a valid number
		if len(s) == 1 {
			return 0, false
		}
	case '+':
		start++
		// if input was `+`, it shouldn't be considered as a valid number
		if len(s) == 1 {
			return 0, false
		}
	case '0': // check if there's more leading zeros
		// for idx, c := range s {
		for ; start < len(s); start++ {
			if s[start] != '0' {
				break
			}
		}
	}

	sl := len(s) - start
	// for `000000`, it's len(s) will be same as start (index).
	// it's a valid number, therefore return 0.
	if sl == 0 {
		return 0, true
	}

	if (IntType == 64 && sl < 20) || (IntType == 32 && sl < 11) {
		n := 0
		for _, c := range s[start:] {
			if c == ',' { // ignore commas
				continue
			}
			// ASCII code goes like 30 for 0, 39 for 9. Therefore 9 (ascii 39) - 0 (ascii 30) = 9.
			// Therefore c -= '0' converts ASCII into actual integer.
			c -= '0'
			if c > 9 {
				return 0, false
			}
			n = n*10 + int(c)
		}
		// When the number became too big, it will flip the sign.
		// Therefore if the sign has changed to negative, it's larger than what it can handle.
		// For instance, (64bit)
		//   9223372036854775806 ->  9223372036854775806
		//   9223372036854775807 ->  9223372036854775807  // Largest positive
		//   9223372036854775808 -> -9223372036854775808
		//  -9223372036854775807 -> -9223372036854775807  // Largest possible for this code
		//  -9223372036854775808 -> -9223372036854775808  // Largest possible actual negative.
		//  -9223372036854775809 -> -9223372036854775807
		if n < 0 {
			return 0, false
		}
		if neg {
			n = -n
		}
		return n, true
	}
	return 0, false
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

// Itoa converts int to string
func Itoa(i int) string {
	return string(BytesAppendInt(make([]byte, 0, 32), i))
}

// IsNumber will take a string and check if it's a number
func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, v := range s {
		// Number should consist of -, +, and 0-9
		// And only first character can be - or +.
		// Otherwise, it should be between 0-9.
		if i == 0 && (v == '-' || v == '+') {
			continue
		}
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// Count counts string lookup value from string s, returns total
func Count(s string, lookup string) int {
	if len(lookup) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(s)-len(lookup)+1; i++ {
		if s[i:i+len(lookup)] == lookup {
			count += 1
		}
	}
	return count
}

// Split writes to `dst` string slice,
// after process string `s` with a rune delimiter `delim`
func Split(dst []string, s string, delim rune) []string {
	last := 0
	for idx, v := range s {
		if v == delim {
			// if last != idx {
			dst = append(dst, s[last:idx])
			// }
			last = idx + 1 // next line
		}
	}
	if last == len(s) {
		dst = append(dst, "")
	}
	if last != len(s) {
		dst = append(dst, s[last:])
	}
	return dst
}

// Elem will take string, delimiter, and index and return nth (index) item.
func Elem(s string, delim rune, index int) string {
	// Correct index for negative number
	// index = -1, get last item.
	// index = -2, get 2nd last item.

	delimFound := 0
	for _, c := range s {
		if c == delim {
			delimFound += 1
		}
	}
	// If no delim was found, but index is -1 or 0, return the input as is
	if delimFound == 0 && (index == -1 || index == 0) {
		return s
	}

	if index < 0 {
		// Count how many delim's out there.
		if delimFound >= 0 && delimFound+1 >= -index {
			index = delimFound + 1 + index // index is negative and starts with -1
		} else {
			// negative too far, doesn't exist
			return ""
		}
	}

	cur, start, end := 0, 0, 0 // to find
	var idxLastChar int = 0    // idxLastChar is last processed index -- for when string not end with the delim
	var idxLastDelim int = 0
	for i, c := range s {
		// println("** i=", i, "c=", string(c), c == delim)
		idxLastChar = i
		if c == delim {
			start = end
			end = i // note that new value of `end` points to where `delim` is.
			if index == cur {
				// println("** index == cur: ", cur)
				idxLastDelim = i
				break
			}
			cur += 1
			// println("** Match.Cur: ", cur)
		}
	}

	// Looking for outside expected one
	if index > cur {
		return ""
	}

	// Last delimiter found was at 0 (idxLastDelim) but end was pointing last char
	// which means last character was delimiter itself, then return empty string
	if idxLastDelim == 0 && end == len(s)-1 {
		return ""
	}

	// i is last processed byte index; if there was more after where last delim found,
	// then it's the case such as `ghi` from `/abc/def/ghi`, 2
	// its `end` will be 7, idxLastChar will be 10. As it was ended without any delim,
	// its `end` value is invalid
	if end < idxLastChar {
		start = end
		end = len(s)
	}

	// except for the first item (cur==0), we should skip the delim, therefore adding 1 to start.
	if cur > 0 {
		start += 1
	}
	return s[start:end]
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

// Index will return index of a string `sub` in string `s`
// If string `sub` isn't in string `s`, this function will return -1.
func Index(s, sub string) int {
	fLen := len(s)
	sLen := len(sub)

	// Substring can't be larger than s string
	if fLen < sLen {
		return -1
	}

	for i := 0; i < fLen-sLen+1; i++ {
		if sLen != 0 && s[i:i+sLen] == sub {
			return i
		}
	}
	return -1
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

// Left will return first n character of string.
// if n is larger than the length of string, it will return whatever available.
func Left(s string, n int) string {
	if n < 0 {
		n = 0
	}
	if len(s) <= n {
		return s
	}
	return s[0:n]
}

// Right will return last n character of string.
// if n is larger than the length of string, it will return whatever available.
func Right(s string, n int) string {
	if n < 0 {
		n = 0
	}
	if len(s) <= n {
		return s
	}
	return s[len(s)-n:]
}
