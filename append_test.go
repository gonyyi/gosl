// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Append_Path(t *testing.T) {
	var tmp []byte
	tmp = gosl.AppendPath(tmp, "/aaa", "bbb")
	gosl.TestString(t, "/aaa/bbb", string(tmp))
	tmp = gosl.AppendPath(tmp, "", "d", "e")
	gosl.TestString(t, "/aaa/bbb/d/e", string(tmp))
	tmp = gosl.AppendPath(tmp[:0], "/aaa/", "/bbb/")
	gosl.TestString(t, "/aaa/bbb/", string(tmp))
}

func Benchmark_Append_Path(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		var tmp []byte
		for i := 0; i < b.N; i++ {
			// tmp = tmp[:0]
			tmp = gosl.AppendPath(tmp[:0], "/aaa", "bbb", "ccc", "/ddd")
		}
		// println(string(tmp))
	})
}


