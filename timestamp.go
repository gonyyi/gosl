// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/19/2022

package gosl

// Timestamp is to be used to save and compare time.
// While time.Time{} is 24 bytes, this Timestamp only uses 8 bytes.
// This will NOT, and is NOT a suitable data type for computation of time.

const (
	// TimestampFormat will be used as a helper for time.Parse() of standard library
	TimestampFormat = "2006/01/02 15:04:05.000"

	tSec Timestamp = 1000
	tMin           = tSec * 100
	tHr            = tMin * 100
	tDay           = tHr * 100
	tMo            = tDay * 100
	tYr            = tMo * 100
)

// Timestamp is an int64 format of date, time with millisecond.
type Timestamp int64

// Date returns YYYYmmDD
func (t Timestamp) Date() int {
	return int(t / tDay)
}

// Year returns YYYY
func (t Timestamp) Year() int {
	return int(t/tYr) % 10000
}

// Month returns mm
func (t Timestamp) Month() int {
	return int(t/tMo) % 100
}

// Day returns DD
func (t Timestamp) Day() int {
	return int(t/tDay) % 100
}

// Time returns HHMMSS
func (t Timestamp) Time() int {
	return int(t%tDay) / 1000 // because of millisecond
}

// Hour returns HH
func (t Timestamp) Hour() int {
	return int(t/tHr) % 100
}

// Min returns MM
func (t Timestamp) Min() int {
	return int(t/tMin) % 100
}

// Sec returns SS
func (t Timestamp) Sec() int {
	return int(t/tSec) % 100
}

// SetDate will take year, month, date and return a new Timestamp
func (t Timestamp) SetDate(year, month, date int) Timestamp {
	tmp := int64(t)
	if 1900 < year && month < 3000 {
		tmp += -(int64(t.Year()) * int64(tYr)) + int64(year*int(tYr))
	}
	if 0 < month && month < 13 {
		tmp += -(int64(t.Month()) * int64(tMo)) + int64(month*int(tMo))
	}
	if 0 < date && date < 32 {
		tmp += -(int64(t.Day()) * int64(tDay)) + int64(date*int(tDay))
	}
	return Timestamp(tmp)
}

// SetTime will take hour/min/sec and set the Timestamp and return it
func (t Timestamp) SetTime(hour, min, sec int) Timestamp {
	tmp := int64(t)
	if -1 < hour && hour < 24 {
		tmp += -(int64(t.Hour()) * int64(tHr)) + int64(hour*int(tHr))
	}
	if -1 < min && min < 60 {
		tmp += -(int64(t.Min()) * int64(tMin)) + int64(min*int(tMin))
	}
	if -1 < sec && sec < 60 {
		tmp += -int64(t.Sec())*int64(tSec) + int64(sec*int(tSec))
	}
	return Timestamp(tmp)
}

// Parse will take string formatted time and converts to Timestamp
// Acceptable formats are as below:
// - 2006/01/02 15:04:05
// - 2006/01/02 15:04:05.000
func (t Timestamp) Parse(s string) (ts Timestamp, ok bool) {
	//      0123456789_123456789_12
	// Fmt1 2006/01/02 15:04:05
	// Fmt2 2006/01/02 15:04:05.000
	if len(s) != 19 && len(s) != 23 { // has to be either fmt1 or fmt2
		return 0, false
	}

	// basic validation
	for i, c := range s {
		if '0' <= c && c <= '9' {
			continue // this is number
		}
		// else
		switch c {
		case '/':
			if i != 4 && i != 7 {
				return 0, false
			}
		case ' ':
			if i != 10 {
				return 0, false
			}
		case ':':
			if i != 13 && i != 16 {
				return 0, false
			}
		case '.':
			if i != 19 {
				return 0, false
			}
		default:
			return 0, false
		}
	}

	var out Timestamp
	if n, ok := t.validate(s[0:4], 1000, 9999); ok {
		out += n * tYr
		if n, ok = t.validate(s[5:7], 1, 12); ok {
			out += n * tMo
			if n, ok = t.validate(s[8:10], 1, 31); ok {
				out += n * tDay
				if n, ok = t.validate(s[11:13], 0, 23); ok {
					out += n * tHr
					if n, ok = t.validate(s[14:16], 0, 59); ok {
						out += n * tMin
						if n, ok = t.validate(s[17:19], 0, 59); ok {
							out += n * tSec
							if len(s) == 23 {
								if n, ok = t.validate(s[20:23], 0, 999); ok {
									out += n
								}
							}
							return out, true
						}
					}
				}
			}
		}
	}
	return 0, false
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

// String will convert the Timestamp into string using TimestampFormat as its format.
// Note that this will cause an allocation as it writes to string.
func (t Timestamp) String() string {
	out := make([]byte, 0, 23)
	out = t.stringAdd(out, 4, t.Year(), '/')  // 5
	out = t.stringAdd(out, 2, t.Month(), '/') // + 3 = 8
	out = t.stringAdd(out, 2, t.Day(), ' ')   // + 3 = 11
	out = t.stringAdd(out, 2, t.Hour(), ':')  // + 3 = 14
	out = t.stringAdd(out, 2, t.Min(), ':')   // + 3 = 17
	out = t.stringAdd(out, 2, t.Sec(), '.')   // + 3 = 20
	out = t.stringAdd(out, 3, int(t)%1000, 0) // + 3 = 23
	return string(out)
}

func (t Timestamp) stringAdd(dst []byte, length, n int, suffix byte) []byte {
	div := 1
	tmp := 0
	for i := 0; i < length; i++ {
		div = div * 10
	}
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

func (t Timestamp) parseInt(p string) (out int, ok bool) {
	for i := 0; i < len(p); i++ {
		if '0' <= p[i] && p[i] <= '9' {
			out = out + int(p[i]-'0')*t.parseNPower(len(p)-i-1)
		} else {
			return 0, false
		}
	}
	return out, true
}

func (t Timestamp) validate(s string, min, max int) (Timestamp, bool) {
	var mul int = 1
	for i := 0; i < len(s)-1; i++ {
		mul = mul * 10
	}
	var out int
	for _, c := range s {
		if n := c - '0'; -1 < n && n < 10 {
			out = out + mul*int(n)
			mul = mul / 10
		} else {
			return -1, false
		}
	}
	return Timestamp(out), min <= out && out <= max
}
