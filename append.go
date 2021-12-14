// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/14/2021

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

// AppendFills will read `src` and fill dst for `n` bytes.
// Difference between AppendRepeats is that this stops when the target size reached.
// Eg.
//    dst = [], src = ['-', '_']
//    AppendFills(dst, src, 9) --> ['-', '_', '-', '_', '-', '_', '-', '_', '-']
func AppendFills(dst []byte, src []byte, n int) []byte {
	// if src is empty, can't do this.
	if src == nil {
		return dst
	}

	cur := 0
	lastIdx := len(src) - 1
	for i := 0; i < n; i++ {
		dst = append(dst, src[cur])
		cur++
		if cur > lastIdx {
			cur = 0
		}
	}
	return dst
}

// AppendRepeats will append to `dst` for `n` times of `rep`
func AppendRepeats(dst []byte, rep []byte, n int) []byte {
	for i := 0; i < n; i++ {
		dst = append(dst, rep...)
	}
	return dst
}

// AppendRepeat will append to `dst` for `n` times of `rep`
func AppendRepeat(dst []byte, rep byte, n int) []byte {
	for i := 0; i < n; i++ {
		dst = append(dst, rep)
	}
	return dst
}

// AppendStringFit will trim string if longer than wanted length,
// if shorter, it will fill with the filler byte.
// This as generates a new string, it will allocation to memory
func AppendStringFit(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte {
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

// AppendStringFitCenter will append a string at the middle of given targetLength.
func AppendStringFitCenter(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte {
	if length <= 0 {
		return dst
	}
	slen := len(s)
	if slen >= length {
		return AppendStringFit(dst, s, length, filler, overflowMarker)
	}

	d, r := (length-slen)/2, (length-slen)%2
	dst = AppendRepeat(dst, filler, d)
	dst = append(dst, s...)
	dst = AppendRepeat(dst, filler, d+r)
	return dst
}

// AppendStringFitRight will append a string at the right aligned with a given targetLength.
func AppendStringFitRight(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte {
	if length <= 0 {
		return dst
	}
	slen := len(s)
	if slen >= length {
		return AppendStringFit(dst, s, length, filler, overflowMarker)
	}

	dst = AppendRepeat(dst, filler, length-slen)
	dst = append(dst, s...)
	return dst
}

// AppendSize will take a size in int64, convert it to
// string and append to byte slice.
// AppendSize(nil, 4096) => 4KB
func AppendSize(dst []byte, size int64, dec uint8) []byte {
	if size >= GB { // Equal to or greater than GB
		switch {
		case size >= EB:
			return append(AppendFloat64(dst, float64(size)/float64(EB), dec, false), 'E', 'B')
		case size >= PB:
			return append(AppendFloat64(dst, float64(size)/float64(PB), dec, false), 'P', 'B')
		case size >= TB:
			return append(AppendFloat64(dst, float64(size)/float64(TB), dec, false), 'T', 'B')
		default:
			return append(AppendFloat64(dst, float64(size)/float64(GB), dec, false), 'G', 'B')
		}
	} else { // less than GB
		switch {
		case size >= MB:
			return append(AppendFloat64(dst, float64(size)/float64(MB), dec, false), 'M', 'B')
		case size >= KB:
			return append(AppendFloat64(dst, float64(size)/float64(KB), dec, false), 'K', 'B')
		default: // Byte
			return append(AppendInt(dst, int(size), false), 'B')
		}
	}
}

// AppendSizeIn will append the filesize in given format.
func AppendSizeIn(dst []byte, size int64, unit int64, dec uint8, comma bool) []byte {
	switch unit {
	case KB:
		return append(AppendFloat64(dst, float64(size)/float64(KB), dec, comma), 'K', 'B')
	case MB:
		return append(AppendFloat64(dst, float64(size)/float64(MB), dec, comma), 'M', 'B')
	case GB:
		return append(AppendFloat64(dst, float64(size)/float64(GB), dec, comma), 'G', 'B')
	case TB:
		return append(AppendFloat64(dst, float64(size)/float64(TB), dec, comma), 'T', 'B')
	case PB:
		return append(AppendFloat64(dst, float64(size)/float64(PB), dec, comma), 'P', 'B')
	case EB:
		return append(AppendFloat64(dst, float64(size)/float64(EB), dec, comma), 'E', 'B')
	default: // Byte
		return append(AppendInt(dst, int(size), false), 'B')
	}
}

// AppendStringMask will create a new masked string except for first N bytes and last N
// This will be used to mask credentials.
func AppendStringMask(dst []byte, s string, firstN, lastN int) []byte {
	if firstN < 0 {
		firstN = 0
	}
	if lastN < 0 {
		lastN = 0
	}
	// in case of dst was not given, it will create one with the right size
	// to reduce allocation as much as possible
	if dst == nil {
		dst = make([]byte, 0, len(s))
	}
	if len(s) <= firstN+lastN {
		return append(dst, s...)
	}
	dst = append(dst, FirstN(s, firstN)...)
	dst = AppendRepeat(dst, '*', len(s)-firstN-lastN)
	dst = append(dst, LastN(s, lastN)...)
	return dst
}
