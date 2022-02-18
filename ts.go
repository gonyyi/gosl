// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 02/18/2022

package gosl

// TS is a timestamp in int64 format to save and compare times.
// While time.Time{} is 24 bytes, this TS only uses 8 bytes.
// NOTE: This will NOT, and is NOT a suitable data type for any computation of time other than comparison.

const (
	// TSFormat will be used as a helper for time.Parse() of standard library
	TSFormat = "2006/01/02 15:04:05.000"

	TSFmtYearDate TSFmt = 1 << iota // TSFmtYearDate for 2006/01/02
	TSFmtDate                       // TSFmtDate for 01/02
	TSFmtTime                       // TSFmtTime for 15:04
	TSFmtTimeSec                    // TSFmtTimeSec for 15:04:05
	TSFmtTimeMS                     // TSFmtTimeMS for 15:04:05.000

	// TSFmtDefault will have 2006/01/02 15:04:05.000 format
	TSFmtDefault = TSFmtYearDate | TSFmtTimeMS

	tsSec TS = 1000
	tsMin    = tsSec * 100
	tsHr     = tsMin * 100
	tsDay    = tsHr * 100
	tsMo     = tsDay * 100
	tsYr     = tsMo * 100
)

type TSFmt uint8

// TS is an int64 format of date, time with millisecond.
type TS int64

// Date returns YYYYmmDD
func (t TS) Date() int64 {
	return int64(t / tsDay)
}

// Time returns HHMMSS
func (t TS) Time() int64 {
	return int64(t%tsDay) / 1000 // because of millisecond
}

// MS returns millisecond
func (t TS) MS() int64 {
	return int64(t) % 1000
}

// IsValid will check to make sure TS will be 17 bytes when converted,
// by simply checking for if it's greater than 10000000000000000.
// Maybe in the future, this can also check for month range, etc.
//     123456789_123456789
// Eg. 20060102150405000
func (t TS) IsValid() bool {
	return t > 10000000000000000
}

// SetDate will take year, month, date and return a new TS
func (t TS) SetDate(year, month, date int) TS {
	if (1000 < year && year < 9999) && (0 < month && month < 13) && (0 < date && date < 32) {
		d := TS(year*10000 + month*100 + date)
		return d*tsDay + TS(t.Time()*1000+t.MS())
	}
	return 0
}

// SetTime will take hour/min/sec and set the TS and return it
func (t TS) SetTime(hour, min, sec int) TS {
	if (-1 < hour && hour < 24) && (-1 < min && min < 60) && (-1 < sec && sec < 60) {
		return t - TS(t.Time()*1000) + TS(hour*10000+min*100+sec)*1000
	}
	return 0
}

// Parse will take string formatted time and converts to TS
// Acceptable formats are as below:
//        0123456789_123456789_12
// - 23c: 2006/01/02 15:04:05.000
// - 19c: 2006/01/02 15:04:05
// - 17c: 20060102150405000
// - 14c: 20060102150405
func (t TS) Parse(s string, fallback TS) (ts TS) {
	ls := len(s)
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

// Byte returns 8-byte value of TS
func (t TS) Byte() [8]byte {
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

// ParseByte takes 8-byte value to TS
func (t TS) ParseByte(b [8]byte) TS {
	return TS(
		int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
			int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56)
}

// String will convert the TS into string using TSFormat as its format.
// NOTE: that this will cause an allocation as it writes to string.
func (t TS) String() string {
	out := make([]byte, 0, 17)
	t.stringAdd(out, 17, int64(t), 0)
	out = t.stringAdd(out, 8, t.Date(), 0)
	out = t.stringAdd(out, 6, t.Time(), 0)
	out = t.stringAdd(out, 3, int64(t)%1000, 0) // + 3 = 23
	return string(out)
}

// Format will convert the TS into string using TSFormat as its format.
// Note that this will cause an allocation as it writes to string.
func (t TS) Format(flag TSFmt) string {
	if flag == 0 {
		flag = TSFmtDefault
	}

	out := make([]byte, 0, 23)
	if flag&(TSFmtDate|TSFmtYearDate) != 0 {
		if flag&TSFmtYearDate != 0 {
			out = t.stringAdd(out, 4, int64(t/tsYr), '/') // 5
		}
		out = t.stringAdd(out, 2, int64(t/tsMo)%100, '/') // + 3 = 8
		out = t.stringAdd(out, 2, int64(t/tsDay)%100, 0)  // + 3 = 11
		if flag&(TSFmtTime|TSFmtTimeMS|TSFmtTimeSec) != 0 {
			out = append(out, ' ')
		}
	}

	if flag&(TSFmtTime|TSFmtTimeSec|TSFmtTimeMS) != 0 {
		out = t.stringAdd(out, 2, int64(t/tsHr)%100, ':') // + 3 = 14
		out = t.stringAdd(out, 2, int64(t/tsMin)%100, 0)  // + 2 = 15
		if flag&(TSFmtTimeSec|TSFmtTimeMS) != 0 {
			out = append(out, ':')
			out = t.stringAdd(out, 2, int64(t/tsSec)%100, 0) // + 3 = 19
			if flag&TSFmtTimeMS != 0 {
				out = append(out, '.')
				out = t.stringAdd(out, 3, t.MS(), 0) // + 4 = 23
			}
		}
	}

	return string(out)
}

func (t TS) stringAdd(dst []byte, length int, n int64, suffix byte) []byte {
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

func (t TS) parseNPower(n int) (out int) {
	out = 1
	for i := 0; i < n; i++ {
		out = out * 10
	}
	return out
}

func (t TS) validate(prevOK bool, prevTS, outUnit TS, s []byte, min, max int) (TS, bool) {
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
	return (outUnit * TS(out)) + prevTS, min <= out && out <= max
}
