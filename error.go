// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/14/2022

package gosl

// Err is string type for error to save errors on constant.
// However, when this is being used, need to return a pointer of Err
// to save allocations
type Err string

// Error is to meet error interface
func (e Err) Error() string {
	return string(e)
}

// NewError takes string and creates an error
// If the string received is empty, it will return error with nil value.
func NewError(s string) error {
	if s == "" {
		return nil
	}
	e := Err(s)
	return &e
}

// errWrap will be used to wrap an error
type errWrap struct {
	err  string
	prev error
}

// Unwrap for the error interface with wrap-able error
func (e *errWrap) Unwrap() error {
	if e.prev != nil {
		return e.prev
	}
	return nil
}

// Error to meet the error interface
func (e *errWrap) Error() string {
	if e.prev != nil {
		buf := make(Buf, 0, 1024)
		buf = buf.WriteString(e.err)
		// buf = buf.WriteBytes(':')
		// buf = buf.WriteString(e.prev.Error())
		// println("cur:", e.Err, "prv:", e.prev.Error())
		return buf.String()
	}
	return e.err
}

// IsError will check if error is same, or contains the given error
func IsError(err, lookup error) bool {
	if err == lookup {
		return true
	}
	if ew, ok := err.(interface {
		Error() string
		Unwrap() error
	}); ok {
		if euw := ew.Unwrap(); euw != nil {
			if euw == lookup {
				return true
			}
			return IsError(ew.Unwrap(), lookup)
		}
	}
	return err == lookup
}

// UnwrapError will unwrap error if available
func UnwrapError(e error) error {
	if ew, ok := e.(interface {
		Error() string
		Unwrap() error
	}); ok {
		return ew.Unwrap()
	}
	return nil
}

// WrapError will create a new error that has info
func WrapError(info string, e error) error {
	if e != nil {
		return &errWrap{
			err:  info + ": " + e.Error(),
			prev: e,
		}
	}
	return &errWrap{
		err: info,
	}
}

// IfPanic will take name and a function func(error),
// if function f is given, it will use that function,
// if not given, print the message using println (stdout)
// Usage:
//     func hello() (out string) {
//         defer IfPanic(func(m interface{}) { out = m.(string) })
//         panic("whatever")
//     }
func IfPanic(f func(a interface{})) {
	// This will only execute when recover() has something
	if r := recover(); r != nil && f != nil {
		f(r)
	}
}
