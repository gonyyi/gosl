// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/3/2021

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

// StringJoin takes a `dst` byte slice,
// and write joined string to it using string slice `p` and byte `delim`
func StringJoin(dst []byte, p []string, delim byte) []byte {
	buf := make(Buf, 0, 1024)
	for i, v := range p {
		if i != 0 {
			buf = buf.WriteByte(delim)
		}
		buf = buf.WriteString(v)
	}
	dst = append(dst, buf...)
	return dst
}


