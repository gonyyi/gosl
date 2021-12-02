// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/01/2021

package gosl

// KeyVal is a key value set.
// This will be useful for variadic function (...)
// Example would be:
//   func AddButton(kv ...KeyVal) // where Key is name of the button, Val is the response ID.
// Val is an interface{} type and anything can be in it.
type KeyVal struct {
	Key string
	Val interface{}
}

// ValBool will type-cast Val into bool
func (kv *KeyVal) ValBool() bool {
	if v, ok := kv.Val.(bool); ok {
		return v
	}
	return false
}

// ValInt will type-cast Val into int
func (kv *KeyVal) ValInt() int {
	if v, ok := kv.Val.(int); ok {
		return v
	}
	return -1
}

// ValFloat64 will type-cast Val into []byte
func (kv *KeyVal) ValFloat64() float64 {
	if v, ok := kv.Val.(float64); ok {
		return v
	}
	return -1
}

// ValString will type-cast Val into string
func (kv *KeyVal) ValString() string {
	if v, ok := kv.Val.(string); ok {
		return v
	}
	return ""
}

// ValBytes will type-cast Val into []byte
func (kv *KeyVal) ValBytes() []byte {
	if v, ok := kv.Val.([]byte); ok {
		return v
	}
	return nil
}
