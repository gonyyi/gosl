// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 02/18/2022

package gosl

// Timestamp is a timestamp in int64 format to save and compare times.
// While time.Time{} is 24 bytes, this Timestamp only uses 8 bytes.
// NOTE: This will NOT, and is NOT a suitable data type for any computation of time other than comparison.

// Timestamp is an int64 format of date, time with millisecond.
type Timestamp int64
type tFormat uint8

const (
	// TDefaultFormat will be used as a helper for time.Parse() of standard library
	// Default output format will be TDefaultFormat
	TDefaultFormat = "2006/01/02 15:04:05.000"

	TYrMoDay    tFormat = 1 << iota // TYrMoDay for 2006/01/02
	TMoDay                          // TMoDay for 01/02
	THrMin                          // THrMin for 15:04
	THrMinSec                       // THrMinSec for 15:04:05
	THrMinSecMs                     // THrMinSecMs for 15:04:05.000

	// TDefault will have 2006/01/02 15:04:05.000 format
	TDefault = TYrMoDay | THrMinSecMs

	tsSec Timestamp = 1000
	tsMin           = tsSec * 100
	tsHr            = tsMin * 100
	tsDay           = tsHr * 100
	tsMo            = tsDay * 100
	tsYr            = tsMo * 100
)

// Date returns YYYYmmDD
func (t Timestamp) Date() int64 {
	return int64(t / tsDay)
}

// Time returns HHMMSS
func (t Timestamp) Time() int64 {
	return int64(t%tsDay) / 1000 // because of millisecond
}

// MS returns millisecond
func (t Timestamp) MS() int64 {
	return int64(t) % 1000
}

// IsValid will check to make sure Timestamp will be 17 bytes when converted,
// by simply checking for if it's greater than 10000000000000000.
// Maybe in the future, this can also check for month range, etc.
//     123456789_123456789
// Eg. 20060102150405000
func (t Timestamp) IsValid() bool {
	return t > 10000000000000000
}

// SetDate will take year, month, date and return a new Timestamp
func (t Timestamp) SetDate(year, month, date int) Timestamp {
	if (1000 < year && year < 9999) && (0 < month && month < 13) && (0 < date && date < 32) {
		d := Timestamp(year*10000 + month*100 + date)
		return d*tsDay + Timestamp(t.Time()*1000+t.MS())
	}
	return 0
}

// SetTime will take hour/min/sec and set the Timestamp and return it
func (t Timestamp) SetTime(hour, min, sec int) Timestamp {
	if (-1 < hour && hour < 24) && (-1 < min && min < 60) && (-1 < sec && sec < 60) {
		return t - Timestamp(t.Time()*1000) + Timestamp(hour*10000+min*100+sec)*1000
	}
	return 0
}

// Parse will take string formatted time and converts to Timestamp
// Acceptable formats are as below:
//        0123456789_123456789_12
// - 23c: 2006/01/02 15:04:05.000
// - 19c: 2006/01/02 15:04:05
// - 17c: 20060102150405000
// - 14c: 20060102150405
func (t Timestamp) Parse(s string, fallback Timestamp) (ts Timestamp) {
	ls := len(s)
	if ls > 23 { // time.Now().String() ==> 2022-03-21 13:05:21.630601 -0500 CDT m=+0.000628238
		ls = 23
		s = s[:23]
	}
	if ls != 14 && ls != 17 && ls != 19 && ls != 23 { // has to be either fmt1 or fmt2
		return fallback
	}

	// Filter only integers into byte array. Total should be either 14 or 17.
	var tmp [17]byte
	var tmpCur = 0
	for _, c := range s {
		if '0' <= c && c <= '9' {
			tmp[tmpCur] = byte(c)
			tmpCur += 1
		}
	}
	if tmpCur != 14 && tmpCur != 17 {
		return fallback
	}

	n, ok := t.validate(true, 0, tsYr, tmp[0:4], 1000, 9999)
	n, ok = t.validate(ok, n, tsMo, tmp[4:6], 1, 12)
	n, ok = t.validate(ok, n, tsDay, tmp[6:8], 1, 31)
	n, ok = t.validate(ok, n, tsHr, tmp[8:10], 0, 23)
	n, ok = t.validate(ok, n, tsMin, tmp[10:12], 0, 59)
	n, ok = t.validate(ok, n, tsSec, tmp[12:14], 0, 59)
	if tmpCur == 17 {
		n, ok = t.validate(ok, n, 1, tmp[14:17], 0, 999)
	}

	if ok {
		return n
	}
	return fallback
}

// Byte returns 8-byte value of Timestamp
func (t Timestamp) Byte() [8]byte {
	return [8]byte{
		byte(0xff & t),
		byte(0xff & (t >> 8)),
		byte(0xff & (t >> 16)),
		byte(0xff & (t >> 24)),
		byte(0xff & (t >> 32)),
		byte(0xff & (t >> 40)),
		byte(0xff & (t >> 48)),
		byte(0xff & (t >> 56)),
	}
}

// ParseByte takes 8-byte value to Timestamp
func (t Timestamp) ParseByte(b [8]byte) Timestamp {
	return Timestamp(
		int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
			int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56)
}

// Append will append the timestamp to []byte dst.
// eg. 20220321130521630 (YYYYMMDDHHMMSS000)
func (t Timestamp) Append(dst []byte) []byte {
	return t.stringAdd(dst, 17, int64(t), 0)
}

// String will convert the Timestamp into string
// eg. 20220321130521630 (YYYYMMDDHHMMSS000)
func (t Timestamp) String() string {
	out := make([]byte, 0, 17)
	out = t.Append(out)
	return string(out)
}

// Format formats the string and append to []byte dst.
func (t Timestamp) Format(dst []byte, flag tFormat) []byte {
	if flag == 0 {
		flag = TDefault
	}
	if flag&(TMoDay|TYrMoDay) != 0 {
		if flag&TYrMoDay != 0 {
			dst = t.stringAdd(dst, 4, int64(t/tsYr), '/') // 5
		}
		dst = t.stringAdd(dst, 2, int64(t/tsMo)%100, '/') // + 3 = 8
		dst = t.stringAdd(dst, 2, int64(t/tsDay)%100, 0)  // + 3 = 11
		if flag&(THrMin|THrMinSecMs|THrMinSec) != 0 {
			dst = append(dst, ' ')
		}
	}

	if flag&(THrMin|THrMinSec|THrMinSecMs) != 0 {
		dst = t.stringAdd(dst, 2, int64(t/tsHr)%100, ':') // + 3 = 14
		dst = t.stringAdd(dst, 2, int64(t/tsMin)%100, 0)  // + 2 = 15
		if flag&(THrMinSec|THrMinSecMs) != 0 {
			dst = append(dst, ':')
			dst = t.stringAdd(dst, 2, int64(t/tsSec)%100, 0) // + 3 = 19
			if flag&THrMinSecMs != 0 {
				dst = append(dst, '.')
				dst = t.stringAdd(dst, 3, t.MS(), 0) // + 4 = 23
			}
		}
	}
	return dst
}

// Formats will convert the Timestamp into string
func (t Timestamp) Formats(flag tFormat) string {
	out := make([]byte, 0, 23)
	out = t.Format(out, flag)
	return string(out)
}

func (t Timestamp) stringAdd(dst []byte, length int, n int64, suffix byte) []byte {
	var tmp int64 = 0
	div := int64(t.parseNPower(length))

	for i := 0; i < length; i++ {
		tmp = div / 10
		dst = append(dst, byte('0'+(n%div)/tmp))
		div = tmp
	}
	if suffix != 0 {
		dst = append(dst, suffix)
	}
	return dst
}

func (t Timestamp) parseNPower(n int) (out int) {
	out = 1
	for i := 0; i < n; i++ {
		out = out * 10
	}
	return out
}

func (t Timestamp) validate(prevOK bool, prevTS, outUnit Timestamp, s []byte, min, max int) (Timestamp, bool) {
	if prevOK == false {
		return t, false
	}
	mul := t.parseNPower(len(s) - 1)

	var out int
	for _, c := range s {
		if c >= '0' && c-'0' < 10 {
			out = out + mul*int(c-'0')
			mul = mul / 10
		} else {
			return -1, false
		}
	}
	return (outUnit * Timestamp(out)) + prevTS, min <= out && out <= max
}
