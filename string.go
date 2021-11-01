// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

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

func StringSplit(dst []string, delim rune, s string) []string {
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

func StringJoin(in []string, delim byte, dst []byte) []byte {
	buf := make(Buf, 0, 1024)
	for i, v := range in {
		if i != 0 {
			buf = buf.WriteByte(delim)
		}
		buf = buf.WriteString(v)
	}
	dst = append(dst, buf...)
	return dst
}

