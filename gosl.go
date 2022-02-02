// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 02/02/2022

package gosl

type Ver = string 

const (
	VERSION Ver = "Gosl v0.7.8"

	KB int64 = 1024
	MB       = KB * 1024
	GB       = MB * 1024
	TB       = GB * 1024
	PB       = TB * 1024
	EB       = PB * 1024

	K int64 = 1000
	M       = K * 1000
	B       = M * 1000
	T       = B * 1000

	// IntType will hold a value regarding current system is 32 bit or 64 bit.
	IntType = 32 << (^uint(0) >> 63) // 64 or 32
)

var (
	GlobalBufferSize = 1024

	EOF = NewError("EOF") // EOF can be updated by io.EOF or any other eg. `gosl.EOF = io.EOF`
)

