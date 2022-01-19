// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// BytesAppendBool will append bool b to bytes dst
func BytesAppendBool(dst []byte, b bool) []byte {
	if b {
		return append(dst, "true"...)
	}
	return append(dst, "false"...)
}

// BytesAppendInt will append int i to bytes dst
func BytesAppendInt(dst []byte, i int) []byte {
	var tmp [32]byte
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
		tmp[cur] = byte('0' + r) // 10/18/2021, ascii of dec(48) = 0
		cur += 1
		nCur += 1
		if i == 0 {
			break
		}
	}
	if neg {
		tmp[cur] = '-'
		cur += 1
	}

	BytesReverse(tmp[:cur])
	return append(dst, tmp[:cur]...)
}

// BytesAppendPrefix will append bytes prefix to bytes dst IF it doesn't have one already
func BytesAppendPrefix(dst []byte, prefix ...byte) []byte {
	if BytesHasPrefix(dst, prefix...) {
		return dst
	}
	return BytesInsert(dst, 0, prefix...)
}

// BytesAppendPrefixString will append string prefix to bytes dst IF it doesn't have one already
func BytesAppendPrefixString(dst []byte, prefix string) []byte {
	if BytesHasPrefixString(dst, prefix) {
		return dst
	}
	return BytesInsertString(dst, 0, prefix)
}

// BytesAppendSize will take a size in int64, convert it to
// string and append to byte slice.
// BytesAppendSize(nil, 4096) => 4KB
func BytesAppendSize(dst []byte, size int64, dec uint8) []byte {
	if size >= GB { // Equal to or greater than GB
		switch {
		case size >= EB:
			return append(BytesAppendFloat(dst, float64(size)/float64(EB), dec), 'E', 'B')
		case size >= PB:
			return append(BytesAppendFloat(dst, float64(size)/float64(PB), dec), 'P', 'B')
		case size >= TB:
			return append(BytesAppendFloat(dst, float64(size)/float64(TB), dec), 'T', 'B')
		default:
			return append(BytesAppendFloat(dst, float64(size)/float64(GB), dec), 'G', 'B')
		}
	} else { // less than GB
		switch {
		case size >= MB:
			return append(BytesAppendFloat(dst, float64(size)/float64(MB), dec), 'M', 'B')
		case size >= KB:
			return append(BytesAppendFloat(dst, float64(size)/float64(KB), dec), 'K', 'B')
		default: // Byte
			return append(BytesAppendInt(dst, int(size)), 'B')
		}
	}
}

// BytesAppendSizeIn will append the filesize in given format.
func BytesAppendSizeIn(dst []byte, size int64, unit int64, dec uint8) []byte {
	switch unit {
	case KB:
		return append(BytesAppendFloat(dst, float64(size)/float64(KB), dec), 'K', 'B')
	case MB:
		return append(BytesAppendFloat(dst, float64(size)/float64(MB), dec), 'M', 'B')
	case GB:
		return append(BytesAppendFloat(dst, float64(size)/float64(GB), dec), 'G', 'B')
	case TB:
		return append(BytesAppendFloat(dst, float64(size)/float64(TB), dec), 'T', 'B')
	case PB:
		return append(BytesAppendFloat(dst, float64(size)/float64(PB), dec), 'P', 'B')
	case EB:
		return append(BytesAppendFloat(dst, float64(size)/float64(EB), dec), 'E', 'B')
	default: // Byte
		return append(BytesAppendInt(dst, int(size)), 'B')
	}
}

// BytesAppendStrings appends string slice to bytes with the delim. If delim is nil,
// it won't add any delim.
func BytesAppendStrings(dst []byte, s []string, delim ...byte) []byte {
	lastIdx := len(s) - 1
	for idx, v := range s {
		dst = append(dst, v...)
		if delim != nil && idx != lastIdx {
			dst = append(dst, delim...)
		}
	}
	return dst
}

// BytesAppendSuffix will append bytes suffix to bytes dst IF it doesn't have one already
func BytesAppendSuffix(dst []byte, suffix ...byte) []byte {
	if BytesHasSuffix(dst, suffix...) {
		return dst
	}
	return append(dst, suffix...)
}

// BytesAppendSuffixString will append string suffix to bytes dst IF it doesn't have one already
func BytesAppendSuffixString(dst []byte, suffix string) []byte {
	if BytesHasSuffixString(dst, suffix) {
		return dst
	}
	return append(dst, suffix...)
}

// BytesCopy will deep copy the source
func BytesCopy(source []byte) []byte {
	out := make([]byte, len(source))
	for i, c := range source {
		out[i] = c
	}
	// println("BytesCopy ==>", string(source), "-->", string(out))
	return out
}

// BytesCount counts byte `c` value from bytes `p`, returns total
func BytesCount(p []byte, c byte) int {
	count := 0
	for _, v := range p {
		if v == c {
			count += 1
		}
	}
	return count
}

// BytesElem will split the bytes and find an item with the given index
// Example: BytesElem( []bytes("/abc/def/ghi"), 2 ) ==> def
func BytesElem(dst []byte, delim byte, index int) []byte {
	// Correct index for negative number
	// index = -1, get last item.
	// index = -2, get 2nd last item.
	if index < 0 {
		// Count how many delim's out there.
		delimCount := 0
		for i := 0; i < len(dst); i++ {
			if dst[i] == delim {
				delimCount += 1
			}
		}
		// println("delimCount=", delimCount, "index", index)
		if delimCount >= 0 && delimCount+1 >= -index {
			index = delimCount + 1 + index // index is negative and starts with -1
		} else {
			// negative too far, doesn't exist
			return dst[:0]
		}
	}

	curr, start, end := 0, 0, 0
	var i int

	for i = 0; i < len(dst); i++ {
		if dst[i] == delim {
			start = end
			end = i // note that new value of `end` points to where `delim` is.
			if index == curr {
				break
			}
			curr += 1
		}
	}

	// when delimiter is found, `curr` is for next item;
	// but, if that's smaller than wanted index, return nil
	if curr < index {
		return nil
	}

	// i is last processed byte index; if there was more after where last delim found,
	// then it's the case such as `ghi` from `/abc/def/ghi`
	if end < i {
		start = end
		end = len(dst)
	}

	// except for the first item, should skip the delim, therefore adding 1 to start.
	if curr > 0 {
		start += 1
	}

	return dst[start:end]
}

// BytesEqual will deep compare two byte slices
func BytesEqual(dst []byte, cmp []byte) bool {
	// if diff len, return false
	if len(dst) != len(cmp) {
		return false
	}
	// if one is nil, return false
	if (dst == nil && cmp != nil) || (dst != nil && cmp == nil) {
		return false
	}
	// if both are nil, return true
	if dst == nil && cmp == nil {
		return true
	}
	// At this point, length of both Buf are identical
	for i := 0; i < len(cmp); i++ {
		if dst[i] != cmp[i] {
			return false
		}
	}
	return true
}

// BytesFilterAny will take anyChar string, and filter the bytes
// If keep is true, it will only keep chars in the anyChar.
// If keep is false, it will only keep chars NOT in the anyChar.
func BytesFilterAny(dst []byte, any string, keep bool) []byte {
	// Instead of looping any, save it to speed up.
	item := make([]byte, 256)
	for i := 0; i < len(any); i++ {
		item[any[i]] = 'X'
	}

	cur := 0

	if keep {
		// include mode: only from the any to be added
		for i := 0; i < len(dst); i++ {
			if item[dst[i]] == 'X' {
				dst[cur] = dst[i]
				cur += 1
			}
		}
		return dst[:cur]
	}

	for i := 0; i < len(dst); i++ {
		if item[dst[i]] != 'X' {
			dst[cur] = dst[i]
			cur += 1
		}
	}
	return dst[:cur]
}

// BytesAppendFloat takes decimalPlace (0-4)
// Since float32's accuracy is so low, this will use float64 exclusively.
func BytesAppendFloat(dst []byte, value float64, decimal uint8) (out []byte) {
	// If panic, return dst
	defer IfPanic(func(i interface{}) { out = dst })
	
	// Handling NaN issue
	if value != value {
		return append(dst, "NaN"...)
	}
	
	// If decimal is not required, then treat it as an integer
	if decimal == 0 {
		return BytesAppendInt(dst, int(value))
	}

	// dec will be used for calculating decimal points (eg. 2 --> 100)
	var dec int = 1
	for i := 0; i < int(decimal); i++ {
		dec = dec * 10
	}

	// for simple calculation, pre-calculate signs.
	if value < 0 {
		dst = append(dst, '-')
		value = -value // flip the sign
	}

	// d for real number
	// r for below decimal point (reminder)
	d := int(value) // 123.456 --> 123
	r := int((value - float64(d)) * float64(dec))

	dst = BytesAppendInt(dst, d) // add real number
	dst = append(dst, '.')       // and decimal

	// Calculate how many leading zeros are needed below decimal place
	if dec > r { // this should be always true
		tmp := -1 // since last line will add at least 1 digit to it
		if r == 0 {
			tmp = int(decimal) - 1
		} else {
			for i := dec; r/i == 0; i = i / 10 {
				tmp += 1
			}
		}

		for i := 0; i < tmp; i++ {
			dst = append(dst, '0')
		}
	}

	dst = BytesAppendInt(dst, r) // add a dot and reminders
	return dst
}

// BytesHasPrefix will check if bytes prefix is in bytes dst
func BytesHasPrefix(dst []byte, prefix ...byte) bool {
	if len(dst) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if dst[i] != prefix[i] {
			return false
		}
	}
	return true
}

// BytesHasPrefixString will check if string prefix is in bytes dst
func BytesHasPrefixString(dst []byte, prefix string) bool {
	if len(dst) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if dst[i] != prefix[i] {
			return false
		}
	}
	return true
}

// BytesHasSuffix will check if bytes suffix is in bytes dst
func BytesHasSuffix(dst []byte, suffix ...byte) bool {
	if len(dst) < len(suffix) {
		return false
	}
	start := len(dst) - len(suffix)
	// Let say dst = "abcdef" and sfx = "def"
	//   len of dst is 6, sfx = 3
	//   dst - sfx = 3
	// therefore, dst[3:] will be compared with sfx[0:]
	for i := 0; i < len(suffix); i++ {
		if dst[start+i] != suffix[i] {
			return false
		}
	}
	return true
}

// BytesHasSuffixString will check if string suffix is in bytes dst
func BytesHasSuffixString(dst []byte, suffix string) bool {
	if len(dst) < len(suffix) {
		return false
	}
	start := len(dst) - len(suffix)
	for i := 0; i < len(suffix); i++ {
		if dst[start+i] != suffix[i] {
			return false
		}
	}
	return true
}

// BytesIndex will search first index where given byte (or bytes) c is in bytes dst.
// If not found, it will return -1.
func BytesIndex(dst []byte, c ...byte) int {
	slen := len(c)
	dstlen := len(dst)
	if dstlen < slen {
		return -1
	}
	// dstlen-slen ==> if dst = "abc", and c = "b", 3-1 = 2;
	// i:=0; i<2; i++ will do 0 and 1. but not 2. therefore adding 1 making it dstlen-slen(+1)
	for i := 0; i < dstlen-slen+1; i++ {
		// println("BytesIndex:", i, "=>", string(dst[i:i+slen]), dstlen, slen)
		if BytesEqual(dst[i:i+slen], c) {
			return i
		}
	}
	return -1
}

// BytesIndexString will search first index where given string s is in bytes dst.
func BytesIndexString(dst []byte, s string) int {
	slen := len(s)
	dstlen := len(dst)
	if dstlen < slen {
		return -1
	}
	for i := 0; i < dstlen-slen+1; i++ {
		if slen != 0 && string(dst[i:i+slen]) == s {
			return i
		}
	}
	return -1
}

// BytesInsert will append byte(bytes) `p`, into bytes `dst` at `index`
func BytesInsert(dst []byte, index int, p ...byte) []byte {
	if index < 0 {
		index = 0
	}
	if len(dst) < index {
		return append(dst, p...)
	}
	BytesReverse(dst[index:])
	BytesReverse(p)
	dst = append(dst, p...)
	BytesReverse(dst[index:])
	return dst
}

// BytesInsertString will append string `s`, into bytes `dst` at `index`
func BytesInsertString(dst []byte, index int, s string) []byte {
	if index < 0 {
		index = 0
	}
	if len(dst) < index {
		return append(dst, s...)
	}
	BytesReverse(dst[index:])
	dl := len(dst)
	dst = append(dst, s...)
	BytesReverse(dst[dl:])
	BytesReverse(dst[index:])
	return dst
}

// BytesLastByte will return the last byte of bytes `dst`
// If dst is nil or size is 0, it will return 0
func BytesLastByte(dst []byte) byte {
	if i := len(dst); i > 0 {
		return dst[i-1]
	}
	return 0
}

// BytesReplace will replace all `old` byte of bytes `p` into `new` byte
// Note: This modifies given slice.
func BytesReplace(p []byte, old, new byte) {
	for idx, c := range p {
		if c == old {
			p[idx] = new
		}
	}
}

// BytesReverse will reverse the byte slice
// Although this returns []byte, there's no need for that as original dst would been modified.
// Note: This modifies given slice.
func BytesReverse(p []byte) {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

// BytesShift will take index and length of the target, and shift it `shift`.
// If `shift` is a negative integer, it will shift to left, if positive, then to right.
// Note: This modifies given slice.
// Example. BytesShift([]byte("abc123"), 3, 3, -3) => "123abc"
func BytesShift(p []byte, index, length, shift int) bool {
	// if no shift, or length is 0, no need idxTo process,
	// also if idx
	if shift == 0 || length < 1 || index < 0 || index+length > len(p) {
		return false
	}

	// When shift direction is left, new index can't be negative.
	if shift < 0 && (index+shift) < 0 {
		return false
	}

	// When shift is right, can't move further than the len of p
	if shift > 0 && index+length+shift > len(p) {
		return false
	}

	//       ___
	//    0123456789
	// eg.ABC123DEFG, switching idx index=3, idxTo=5 ("12"), shift +2 (right 2)
	//    ABC213DEFG, Step 1. reverse index to idxTo (12 -> 21)
	//    ABC21D3EDG, Step 2. reverse index+shift to idxTo+shift (3D -> D3)
	//    ABC3D12EDG, Step 3. reverse index to idxTo+shift (21D3 -> 3D12)

	// Shift left (negative)
	if shift < 0 {
		BytesReverse(p[index : index+length])
		BytesReverse(p[index+shift : index])
		BytesReverse(p[index+shift : index+length])
		return true
	}

	// Shift right (positive)
	BytesReverse(p[index : index+length])
	BytesReverse(p[index+length : index+length+shift])
	BytesReverse(p[index : index+length+shift])
	return true
}

// BytesToLower will replace bytes dst into lowercase
// Note: This modifies given slice.
func BytesToLower(p []byte) {
	adj := byte('a') - byte('A')
	for idx, c := range p {
		if 'A' <= c && 'Z' >= c {
			p[idx] = c + adj
		}
	}
}

// BytesToUpper will replace bytes dst into uppercase
// Note: This modifies given slice.
func BytesToUpper(p []byte) {
	adj := byte('a') - byte('A')
	for idx, c := range p {
		if 'a' <= c && 'z' >= c {
			p[idx] = c - adj
		}
	}
}

// BytesTrimPrefix will check if bytes `dst` has a byte(s) `prefix`,
// if it has this func will trim the prefix.
func BytesTrimPrefix(dst []byte, prefix ...byte) []byte {
	if BytesHasPrefix(dst, prefix...) {
		return dst[len(prefix):]
	}
	return dst
}

// BytesTrimPrefixString will check if bytes `dst` has a string `prefix`,
// if it has this func will trim the prefix.
func BytesTrimPrefixString(dst []byte, prefix string) []byte {
	if BytesHasPrefixString(dst, prefix) {
		return dst[len(prefix):]
	}
	return dst
}

// BytesTrimSuffix will check if bytes `dst` has a suffix byte(s) `suffix`,
// if it has this func will trim the prefix.
func BytesTrimSuffix(dst []byte, suffix ...byte) []byte {
	if BytesHasSuffix(dst, suffix...) {
		return dst[:len(dst)-len(suffix)]
	}
	return dst
}

// BytesTrimSuffixString will check if bytes `dst` has a suffix string `suffix`,
// if it has this func will trim the prefix.
func BytesTrimSuffixString(dst []byte, suffix string) []byte {
	if BytesHasSuffixString(dst, suffix) {
		return dst[:len(dst)-len(suffix)]
	}
	return dst
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

// HexToBytes will convert hex string and append it to byte
// Note that this is code is completely identical to HexStringToBytes.
// Eg. BytesAppendHex(nil, "48656c6c6f20476f6e") // => Hello Gon
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
