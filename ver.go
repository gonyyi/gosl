// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/06/2021

package gosl

// Ver is a string type that holds version information:
// Format: "[Name] v[Major].[Minor].[Patch]-[Build]"

// NewVer creates a Ver string.
func NewVer(name string, major, minor, patch, build int) Ver {
	var buf = make([]byte, 0, 512)
	return Ver(newVer(buf, name, major, minor, patch, build))
}

// Ver is a string type. That way it can be used as a part of const.
type Ver string

// String will return version in string
func (v Ver) String() string {
	return string(v)
}

// Name will return the name
func (v Ver) Name() string {
	for i := len(v) - 1; i > 0; i-- {
		if v[i] == ' ' {
			return string(v[0:i])
		}
	}
	return "" // otherwise none
}

// Version will show only version part
func (v Ver) Version() string {
	for i := len(v) - 1; i > 0; i-- {
		if v[i] == ' ' && i != len(v)-1 {
			return string(v[i+1:])
		}
	}
	return ""
}

// IsNewer takes another Ver and compares
// Note that this doesn't consider name part of the version.
func (v Ver) IsNewer(old Ver) bool {
	_, cMajor, cMinor, cPatch, cBuild := v.Parse()
	_, oMajor, oMinor, oPatch, oBuild := old.Parse()

	switch {
	case cMajor > oMajor:
		return true
	case cMajor == oMajor:
		switch {
		case cMinor > oMinor:
			return true
		case cMinor == oMinor:
			switch {
			case cPatch > oPatch:
				return true
			case cPatch == oPatch:
				if cBuild > oBuild {
					return true
				}
			}
		}
	}
	return false
}

// Set will set new version
func (v Ver) Set(name string, major, minor, patch, build int) Ver {
	var buf = make([]byte, 0, 512)
	buf = newVer(buf, name, major, minor, patch, build)
	return Ver(buf)
}

// Clean will parse and recreate version
func (v Ver) Clean() Ver {
	var buf = make([]byte, 0, 512)
	name, major, minor, patch, build := v.Parse()
	buf = newVer(buf, name, major, minor, patch, build)
	if string(buf) == string(v) {
		return v
	}
	return Ver(buf)
}

// Parse will parse current version
func (v Ver) Parse() (name string, major, minor, patch, build int) {
	count := 0
	end := len(v)
	beg := 0
	for i := end - 1; i > 0; i-- {
		if v[i] == '-' {
			build = MustAtoi(string(v[i+1:]), 0)
			end = i
			continue
		}
		if v[i] == ' ' {
			name = string(v[0:i])
			beg = i + 1
			break
		}
	}
	var vers [4]int // only need 3, but extra space...
	cur := beg
	for i := beg; i < end; i++ {
		if v[i] == '.' {
			if v[cur] == 'V' || v[cur] == 'v' {
				cur = cur + 1
			}
			vers[count] = MustAtoi(string(v[cur:i]), 0)
			count += 1
			cur = i + 1
		}
	}
	// Last number
	if end-cur > 0 && v[cur] == 'v' {
		cur += 1
	}
	num, ok := Atoi(string(v[cur:end]))
	if !ok && len(name) == 0 {
		name = string(v[cur:end])
	}
	vers[count] = num

	return name, vers[0], vers[1], vers[2], build
}

// newVer will create new by appending it to dst
func newVer(dst []byte, name string, major, minor, patch, build int) []byte {
	if len(name) > 0 {
		dst = append(append(dst, name...), ' ')
	}
	dst = AppendInt(append(dst, 'v'), major, false)
	dst = AppendInt(append(dst, '.'), minor, false)
	dst = AppendInt(append(dst, '.'), patch, false)
	dst = AppendInt(append(dst, '-'), build, false)
	return dst
}
