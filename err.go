// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// err is string type for error to save errors on constant.
// However, when this is being used, need to return a pointer of err
// to save allocations
type err string

// Error is to meet error interface
func (e err) Error() string {
	return string(e)
}

// NewError takes string and creates an error
func NewError(s string) error {
	return err(s)
}

// IfErr is a simple function that takes an error ID and error..
// If error is not nil, then it will print error message.
// This has zero allocation.
func IfErr(key string, e error) {
	// If given error is not nil, then get a buffer from the internal buffer pool,
	// then write an error as "key = value" format and then write it to
	// os.Stdout using println().
	if e != nil {
		buf := getBufpBuffer()
		buf.WriteString(key).WriteByte('=').WriteString(e.Error())
		println(buf.String())
		buf.ReturnBuffer()
	}
}


