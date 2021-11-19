// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/19/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_String_IntsJoin(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var a = []int{1, 3, 5, 8, 100, 0, -200, 23}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteBytes('[')
		buf = gosl.IntsJoin(buf, a, ',')
		buf = buf.WriteBytes(']')
		gosl.Test(t, "[1,3,5,8,100,0,-200,23]", buf.String())
	})

	t.Run("null", func(t *testing.T) {
		var a = []int{}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteBytes('[')
		buf = gosl.IntsJoin(buf, a, ',')
		buf = buf.WriteBytes(']')
		gosl.Test(t, "[]", buf.String())
	})
}
