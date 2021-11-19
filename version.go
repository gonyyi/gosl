// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/18/2021

package gosl

// Version is a string type that holds version information:
// Format: "[Name] v[Major].[Minor].[Patch]-[Build]"

// NewVersion creates a Version string.
func NewVersion(name string, major, minor, patch, build int) Version {
	var buf = make([]byte, 0, 512)
	return Version(newVersion(buf, name, major, minor, patch, build))
}

// Version is a string type. That way it can be used as a part of const.
type Version string

// String will return version in string
func (v Version) String() string {
	return string(v)
}

// Name will return the name 
func (v Version) Name() string {
	for i := len(v) - 1; i > 0; i-- {
		if v[i] == ' ' {
			return string(v[0:i])
		}
	}
	return "" // otherwise none
}

// Version will parse only version information
func (v Version) Version() (major, minor, patch, build int) {
	_, major, minor, patch, build = v.Parse()
	return major, minor, patch, build
}

// IsNewer takes another Version and compares
// Note that this doesn't consider name part of the version.
func (v Version) IsNewer(old Version) bool {
	cMajor, cMinor, cPatch, cBuild := v.Version()
	oMajor, oMinor, oPatch, oBuild := old.Version()

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
func (v Version) Set(name string, major, minor, patch, build int) Version {
	var buf = make([]byte, 0, 512)
	buf = newVersion(buf, name, major, minor, patch, build)
	return Version(buf)
}


// Clean will parse and recreate version
func (v Version) Clean() Version {
	var buf = make([]byte, 0, 512)
	name, major, minor, patch, build := v.Parse()
	buf = newVersion(buf, name, major, minor, patch, build)
	if string(buf) == string(v) {
		return v
	}
	return Version(buf)
}

// Parse will parse current version
func (v Version) Parse() (name string, major, minor, patch, build int) {
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
	if end - cur > 0 && v[cur] == 'v' {
		cur += 1
	}
	num, ok := Atoi(string(v[cur:end]))
	if !ok && len(name) == 0 {
		name = string(v[cur:end])
	}
	vers[count] = num

	return name, vers[0], vers[1], vers[2], build
}

// newVersion will create new by appending it to dst
func newVersion(dst []byte, name string, major, minor, patch, build int) []byte {
	if len(name) > 0 {
		dst = append(append(dst, name...), ' ')
	}
	dst = AppendInt(append(dst, 'v'), major, false)
	dst = AppendInt(append(dst, '.'), minor, false)
	dst = AppendInt(append(dst, '.'), patch, false)
	dst = AppendInt(append(dst, '-'), build, false)
	return dst
}