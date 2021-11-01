// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// AppendBool will append the `b` to a byte slice.
func AppendBool(dst []byte, b bool) []byte {
	if b {
		return append(dst, `true`...)
	}
	return append(dst, `false`...)
}

// AppendInt will append int to byte slice.
// If param useComma is true, it will add comma for every thousand.
func AppendInt(dst []byte, i int, comma bool) (b []byte) {
	var tmp [64]byte
	var cur uint8
	var nCur int // cursor only for number (when adding a comma)

	var neg bool
	if i < 0 {
		i = -i
		neg = true
	}

	var r int // for reminder
	for {
		i, r = i/10, i%10
		tmp[cur] = byte(48 + r) // 10/18/2021, ascii of dec(48) = 0
		cur += 1
		nCur += 1
		if i == 0 {
			break
		}
		// Anything passing here means there's next number
		// if comma, and number cursor is 3rd digit, then add comma
		if comma && nCur%3 == 0 {
			tmp[cur] = ','
			cur++
		}
	}
	if neg {
		tmp[cur] = '-'
		cur += 1
	}

	// As the logic keep checks remainder of division 10, it was recorded in reversed order.
	// For a float `-123.123`, it would been `321.321-`. Therefore reverse it.
	BytesReverse(tmp[:cur])
	return append(dst, tmp[:cur]...)
}

// AppendFloat64 takes decimalPlace (0-4)
// Since float32's accuracy is so low, this will use float64 exclusively.
func AppendFloat64(dst []byte, value float64, decimal uint8, comma bool) []byte {
	if decimal > 4 {
		decimal = 4
	}
	dec := 100
	switch decimal {
	case 4:
		dec = 10000
	case 3:
		dec = 1000
	case 2:
		dec = 100
	case 1:
		dec = 10
	case 0: // if decimal given is 0, run AppendInt instead.
		return AppendInt(dst, int(value), comma)
	}

	f := int(value * float64(dec))
	if f < 0 {
		dst = append(dst, '-')
		f = -f
	}

	dst = append(AppendInt(dst, f/dec, comma), '.')
	cur := len(dst)
	dst = AppendInt(dst, f%dec, false)
	curNew := len(dst)
	// if added was below target decimal, add more 0..
	for i := curNew - cur; i < int(decimal); i++ {
		dst = append(dst, '0')
	}
	return dst
}

// AppendString will clean the string -- trim string, and
// if there's more than 1 spaces, it will trim it down to 1.
// eg. "   HA     HA   " --> "HA HA"
func AppendString(dst []byte, s string, trim bool) []byte {
	if !trim {
		return append(dst, s...)
	}

	started, lcSpace := false, true
	for i := 0; i < len(s); i++ {
		// if space, ignore for now, but note that it was a space
		if s[i] == ' ' {
			lcSpace = true
			continue
		}
		// if not space, but
		if started && lcSpace {
			dst = append(dst, ' ', s[i])
			lcSpace = false
			continue
		}
		if !started {
			started = true
			lcSpace = false
		}
		dst = append(dst, s[i])
	}
	return dst
}

func AppendPath(dst []byte, path ...string) []byte {
	const delim = '/'
	for _, v := range path {
		// if previous one didn't end with /, then add /
		if len(v) > 0 {
			// get prev
			prev := byte(0)
			if len(dst) > 0 {
				prev = dst[len(dst)-1]
			}

			// if previous = '/' and current = '/', ignore current '/'
			// if previous != '/' and current != '/', add
			if prev == delim && v[0] == delim {
				dst = append(dst, v[1:]...)
				continue
			}
			if prev != delim && v[0] != delim {
				dst = append(dst, delim)
			}
			dst = append(dst, v...)
		}
	}
	return dst
}


