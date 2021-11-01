// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/18/2021

package gosl

// TODO: check if this can utilize internal bufp, if not drop the support.

// Itoa converts int to string
func Itoa(i int, comma bool) (s string) {
	return string(AppendInt(make([]byte, 0, 128), i, comma))
}

func Ftoa(f64 float64, decimal uint8, comma bool) (s string) {
	return string(AppendFloat64(make([]byte, 0, 128), f64, decimal, comma))
}
