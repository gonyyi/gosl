package inttime

const tMin = 100
const tHr = tMin * 100
const tDay = tHr * 100
const tMo = tDay * 100
const tYr = tMo * 100

type IntTime int64

func (t IntTime) Date() int {
	return int(int64(t) / tDay)
}
func (t IntTime) Year() int {
	return int(int64(t) / tYr % 10000)
}
func (t IntTime) Month() int {
	return int(int64(t) / tMo % 100)
}
func (t IntTime) Day() int {
	return int(int64(t) / tDay % 100)
}
func (t IntTime) Time() int {
	return int(int64(t) % tDay / 1000)
}
func (t IntTime) Hour() int {
	return int(int64(t) / tHr % 100)
}
func (t IntTime) Min() int {
	return int(int64(t) / tMin % 100)
}

func (t IntTime) Sec() int {
	return int(int64(t) % 100)
}

func (t IntTime) SetDate(year, month, date int) IntTime {
	tmp := int64(t)
	if 1900 < year && month < 3000 {
		tmp += -(int64(t.Year()) * int64(tYr)) + int64(year*tYr)
	}
	if 0 < month && month < 13 {
		tmp += -(int64(t.Month()) * int64(tMo)) + int64(month*tMo)
	}
	if 0 < date && date < 32 {
		tmp += -(int64(t.Day()) * int64(tDay)) + int64(date*tDay)
	}
	return IntTime(tmp)
}

func (t IntTime) SetTime(hour, min, sec int) IntTime {
	tmp := int64(t)
	if -1 < hour && hour < 24 {
		tmp += -(int64(t.Hour()) * int64(tHr)) + int64(hour*tHr)
	}
	if -1 < min && min < 60 {
		tmp += -(int64(t.Min()) * int64(tMin)) + int64(min*tMin)
	}
	if -1 < sec && sec < 60 {
		tmp += -int64(t.Sec()) + int64(sec)
	}
	return IntTime(tmp)
}

func (t IntTime) stringAdd(dst []byte, length, i int, suffix byte) []byte {
	if length == 4 {
		dst = append(dst, byte('0'+(i%10000)/1000))
		dst = append(dst, byte('0'+(i%1000)/100))
	}
	dst = append(dst, byte('0'+(i%100)/10))
	dst = append(dst, byte('0'+(i%10)/1))
	if suffix != 0 {
		dst = append(dst, suffix)
	}
	return dst
}

func (t IntTime) parseNPower(n int) (out int) {
	out = 1
	for i := 0; i < n; i++ {
		out = out * 10
	}
	return out
}

func (t IntTime) parseInt(p string) (out int, ok bool) {
	for i := 0; i < len(p); i++ {
		if '0' <= p[i] && p[i] <= '9' {
			out = out + int(p[i]-'0')*t.parseNPower(len(p)-i-1)
		} else {
			return 0, false
		}
	}
	return out, true
}

func (t IntTime) Parse(s string) (newTime IntTime, ok bool) {
	if len(s) != 19 {
		return 0, false
	}
	// 0123456789_12345678
	// 2006/01/02 15:04:05
	if s[4] != '/' || s[7] != '/' || s[10] != ' ' || s[13] != ':' || s[16] != ':' {
		return 0, false
	}

	ok, out, tmp := false, 0, 0
	if tmp, ok = t.parseInt(s[0:4]); ok {
		out += tmp * tYr
	} else {
		return 0, false
	}
	if tmp, ok = t.parseInt(s[5:7]); ok {
		out += tmp * tMo
	} else {
		return 0, false
	}
	if tmp, ok = t.parseInt(s[8:10]); ok {
		out += tmp * tDay
	} else {
		return 0, false
	}
	if tmp, ok = t.parseInt(s[11:13]); ok {
		out += tmp * tHr
	} else {
		return 0, false
	}
	if tmp, ok = t.parseInt(s[14:16]); ok {
		out += tmp * tMin
	} else {
		return 0, false
	}
	if tmp, ok = t.parseInt(s[17:19]); ok {
		out += tmp
	} else {
		return 0, false
	}
	return IntTime(out), true
}

func (t IntTime) String() string {
	out := make([]byte, 0, 19)
	out = t.stringAdd(out, 4, t.Year(), '/')
	out = t.stringAdd(out, 2, t.Month(), '/')
	out = t.stringAdd(out, 2, t.Day(), ' ')

	out = t.stringAdd(out, 2, t.Hour(), ':')
	out = t.stringAdd(out, 2, t.Min(), ':')
	out = t.stringAdd(out, 2, t.Sec(), 0)
	return string(out)
}
