// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/19/2021

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

// MustAtoi converts string to integer.
// It takes a string and a fallback integer, and if conversion failed, it will return a fallback value.
func MustAtoi(s string, fallback int) int {
	out, ok := Atoi(s)
	if !ok {
		out = fallback
	}
	return out
}

// ToUpper will take a string and returns uppercase string
// Due to creation of new string, this will have an allocation.
// To avoid, use BytesToUpper instead
func ToUpper(s string) string {
	buf := make([]byte, 0, 512)
	buf = BytesToUpper(append(buf, s...))
	return string(buf)
}

// ToLower will take a string and returns lowercase string
// Due to creation of new string, this will have an allocation.
// To avoid, use BytesToLower instead
func ToLower(s string) string {
	buf := make([]byte, 0, 512)
	buf = BytesToLower(append(buf, s...))
	return string(buf)
}
