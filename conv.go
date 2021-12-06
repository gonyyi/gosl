// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/06/2021

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
	buf := make([]byte, 0, 128)
	buf = BytesToUpper(append(buf, s...))
	return string(buf)
}

// ToLower will take a string and returns lowercase string
// Due to creation of new string, this will have an allocation.
// To avoid, use BytesToLower instead
func ToLower(s string) string {
	buf := make([]byte, 0, 128)
	buf = BytesToLower(append(buf, s...))
	return string(buf)
}

// BytesToHex will take byte (or bytes) and append it to dst
// Eg. BytesToHex(nil, []byte("Hello Gon")) // => 48656c6c6f20476f6e
func BytesToHex(dst []byte, b []byte) []byte {
	for _, c := range b {
		d, r := int(c)/16, int(c)%16
		if d < 10 {
			d = '0' + d
		} else {
			d = 'a' + (d - 10)
		}
		if r < 10 {
			r = '0' + r
		} else {
			r = 'a' + (r - 10)
		}
		dst = append(dst, byte(d), byte(r))
	}
	return dst
}

// StringToHex will take string and append it to dst
// Eg. StringToHex(nil, "Hello Gon") // => 48656c6c6f20476f6e
func StringToHex(dst []byte, b string) []byte {
	for _, c := range b {
		d, r := int(c)/16, int(c)%16
		if d < 10 {
			d = '0' + d
		} else {
			d = 'a' + (d - 10)
		}
		if r < 10 {
			r = '0' + r
		} else {
			r = 'a' + (r - 10)
		}
		dst = append(dst, byte(d), byte(r))
	}
	return dst
}

// HexToBytes will convert hex string and append it to byte
// Note that this is code is completely identical to HexStringToBytes.
// Eg. HexToBytes(nil, "48656c6c6f20476f6e") // => Hello Gon
func HexToBytes(dst []byte, hex []byte) (out []byte, ok bool) {
	// Validate the length
	if len(hex)%2 != 0 {
		return dst, false
	}
	// Validate if only expected are there
	for i := 0; i < len(hex); i++ {
		if ('0' > hex[i] || '9' < hex[i]) && ('a' > hex[i] || 'z' < hex[i]) {
			return dst, false
		}
	}
	// For each pair (as 2 hex code will be 1 byte)
	for i := 0; i < len(hex); i += 2 {
		var cidx int // character index
		// First char of the hex code
		if '0' <= hex[i] && hex[i] <= '9' {
			// for the first char, it will need to multiply by 16
			cidx += int(hex[i]-'0') * 16
		} else if 'a' <= hex[i] && hex[i] <= 'z' {
			// since the order is 0-9 then a-f, a-f will need to be +10.
			cidx += (int(hex[i]-'a') + 10) * 16
		}
		// Second char of the hex code
		if '0' <= hex[i+1] && hex[i+1] <= '9' {
			cidx += int(hex[i+1] - '0')
		} else if 'a' <= hex[i+1] && hex[i+1] <= 'z' {
			cidx += int(hex[i+1]-'a') + 10
		}

		dst = append(dst, byte(cidx))
	}

	return dst, true
}

// HexStringToBytes will convert hex string and append it to byte
// Note that this is code is completely identical to HexToBytes.
// Eg. HexStringToBytes(nil, "48656c6c6f20476f6e") // => Hello Gon
func HexStringToBytes(dst []byte, hex string) (out []byte, ok bool) {
	if len(hex)%2 != 0 {
		return dst, false
	}

	for i := 0; i < len(hex); i++ {
		if ('0' > hex[i] || '9' < hex[i]) && ('a' > hex[i] || 'z' < hex[i]) {
			return dst, false
		}
	}

	for i := 0; i < len(hex); i += 2 {
		var cidx int // character index
		if '0' <= hex[i] && hex[i] <= '9' {
			cidx += int(hex[i]-'0') * 16
		} else if 'a' <= hex[i] && hex[i] <= 'z' {
			cidx += (int(hex[i]-'a') + 10) * 16
		}

		if '0' <= hex[i+1] && hex[i+1] <= '9' {
			cidx += int(hex[i+1] - '0')
		} else if 'a' <= hex[i+1] && hex[i+1] <= 'z' {
			cidx += int(hex[i+1]-'a') + 10
		}

		dst = append(dst, byte(cidx))
	}

	return dst, true
}
