// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/19/2021

package gosl

// IntsJoin takes a `dst` byte slice,
// and write joined integer to it using string slice `p` and byte `delim`
func IntsJoin(dst []byte, p []int, delim byte) []byte {
	buf := make(Buf, 0, 4096)
	for i, v := range p {
		if i != 0 {
			buf = buf.WriteBytes(delim)
		}
		buf = buf.WriteInt(v)
	}
	dst = append(dst, buf...)
	return dst
}
