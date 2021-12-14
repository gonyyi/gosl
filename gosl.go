// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package gosl

// DoNothing will be used to prevent nil error
func DoNothing() {}

// Filesize
const (
	VERSION Ver = "Gosl v0.4.3-0"

	KB int64 = 1024
	MB       = KB * 1024
	GB       = MB * 1024
	TB       = GB * 1024
	PB       = TB * 1024
	EB       = PB * 1024

	IntType = 32 << (^uint(0) >> 63) // 64 or 32
)

var (
	DefaultBufferSize = 2 << 10 // init 1k, max 2k
)
