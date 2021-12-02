// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/01/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestKeyVal_ValBool(t *testing.T) {
	kv := gosl.KeyVal{Key: "1", Val: true}
	gosl.Test(t, "1", kv.Key)
	gosl.Test(t, true, kv.ValBool())
}

func TestKeyVal_ValInt(t *testing.T) {
	kv := gosl.KeyVal{Key: "1", Val: 1}
	gosl.Test(t, "1", kv.Key)
	gosl.Test(t, 1, kv.ValInt())
}

func TestKeyVal_ValFloat64(t *testing.T) {
	kv := gosl.KeyVal{Key: "1", Val: 1.23}
	gosl.Test(t, "1", kv.Key)
	gosl.Test(t, 1.23, kv.ValFloat64())
}

func TestKeyVal_ValString(t *testing.T) {
	kv := gosl.KeyVal{Key: "1", Val: "test"}
	gosl.Test(t, "1", kv.Key)
	gosl.Test(t, "test", kv.ValString())
}

func TestKeyVal_ValByte(t *testing.T) {
	kv := gosl.KeyVal{Key: "1", Val: []byte("abc")}
	gosl.Test(t, "1", kv.Key)
	gosl.Test(t, "abc", string(kv.ValBytes()))
}
