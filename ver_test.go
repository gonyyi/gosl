// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/18/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func TestVer_name(t *testing.T) {
	var v1 gosl.Ver

	v1 = "Gon 1.2.3-123"
	gosl.Test(t, "Gon", v1.Name())

	v1 = "1.2.3-123"
	gosl.Test(t, "", v1.Name())

	v1 = ""
	gosl.Test(t, "", v1.Name())
}

func TestVer_IsNewer(t *testing.T) {
	test := func(v1, v2 gosl.Ver, exp string) {
		var out = v1.String()
		if v1.IsNewer(v2) {
			out += " > "
		} else if v2.IsNewer(v1) {
			out += " < "
		} else {
			out += " = "
		}
		out += v2.String()
		gosl.Test(t, v1.String()+" "+exp+" "+v2.String(), out)
	}

	test("Gon 1.2.3-123", "Gon 1.2.3-123", "=")
	test("Gon 1.2.3-123", "Gon 1.2.3-124", "<")
	test("Gon 1.2.4-123", "Gon 1.2.3-124", ">")
	test("Gon 1.2.4-123", "Gon 1.3.3-124", "<")
	test("Gon 2.2.4-123", "Gon 1.3.3-124", ">")
	test("Gon 21.2.4-123", "Gon 132.3.3-124", "<")

	test("Gon v1.2.3-123", "Gon 1.2.3-123", "=")
	test("Gon v1.2.3-123", "Gon 1.2.3-124", "<")
	test("Gon v1.2.4-123", "Gon 1.2.3-124", ">")
	test("Gon v1.2.4-123", "Gon 1.3.3-124", "<")
	test("Gon v2.2.4-123", "Gon 1.3.3-124", ">")
	test("Gon v21.2.4-123", "Gon 132.3.3-124", "<")

	test("Gon 1.2.3-123", "Gon v1.2.3-123", "=")
	test("Gon 1.2.3-123", "Gon v1.2.3-124", "<")
	test("Gon 1.2.4-123", "Gon v1.2.3-124", ">")
	test("Gon 1.2.4-123", "Gon v1.3.3-124", "<")
	test("Gon 2.2.4-123", "Gon v1.3.3-124", ">")
	test("Gon 21.2.4-123", "Gon v132.3.3-124", "<")

	test("Gon v1.2.3-123", "Gon v1.2.3-123", "=")
	test("Gon v1.2.3-123", "Gon v1.2.3-124", "<")
	test("Gon v1.2.4-123", "Gon v1.2.3-124", ">")
	test("Gon v1.2.4-123", "Gon v1.3.3-124", "<")
	test("Gon v2.2.4-123", "Gon v1.3.3-124", ">")
	test("Gon v21.2.4-123", "Gon v132.3.3-124", "<")
}

func TestVer_Parse(t *testing.T) {
	test := func(v gosl.Ver, exp string) {
		name, maj, min, pat, bld := v.Parse()
		gosl.Test(t, exp, fmt.Sprintf("<%s> %d.%d.%d-%d", name, maj, min, pat, bld))
	}

	test("", "<> 0.0.0-0")
	test("123", "<> 123.0.0-0")
	test("gon 1.2.3-456", "<gon> 1.2.3-456")
	test("gon 1.2.3-", "<gon> 1.2.3-0")
	test("1.2.3-456", "<> 1.2.3-456")
	test("1.2.3", "<> 1.2.3-0")
	test("1.2", "<> 1.2.0-0")
	test("1-456", "<> 1.0.0-456")
	test("gon super 1 1-456", "<gon super 1> 1.0.0-456")

	test("Gosl 1 2.1.1-123", "<Gosl 1> 2.1.1-123")
	test("Gosl 1 0.1.1-123", "<Gosl 1> 0.1.1-123")
	test("Gosl 1 0.1.1", "<Gosl 1> 0.1.1-0")
	test("Gosl 1 0.1-123", "<Gosl 1> 0.1.0-123")
	test("Gosl 1 0.1", "<Gosl 1> 0.1.0-0")
	test("Gosl 1 1-123", "<Gosl 1> 1.0.0-123")
	test("Gosl 1", "<Gosl> 1.0.0-0")

	test("2.1.1-123", "<> 2.1.1-123")
	test("0.1.1-123", "<> 0.1.1-123")
	test("0.1.1", "<> 0.1.1-0")
	test("0.1-123", "<> 0.1.0-123")
	test("0.1", "<> 0.1.0-0")
	test("1-123", "<> 1.0.0-123")
	test("1", "<> 1.0.0-0")

	test("Gosl 1 v2.1.1-123", "<Gosl 1> 2.1.1-123")
	test("Gosl 1 v0.1.1-123", "<Gosl 1> 0.1.1-123")
	test("Gosl 1 v0.1.1", "<Gosl 1> 0.1.1-0")
	test("Gosl 1 v0.1-123", "<Gosl 1> 0.1.0-123")
	test("Gosl 1 v0.1", "<Gosl 1> 0.1.0-0")
	test("Gosl 1 v1-123", "<Gosl 1> 1.0.0-123")
	test("Gosl v1", "<Gosl> 1.0.0-0")

	test("v2.1.1-123", "<> 2.1.1-123")
	test("v0.1.1-123", "<> 0.1.1-123")
	test("v0.1.1", "<> 0.1.1-0")
	test("v0.1-123", "<> 0.1.0-123")
	test("v0.1", "<> 0.1.0-0")
	test("v1-123", "<> 1.0.0-123")
	test("v1", "<> 1.0.0-0")
}

func BenchmarkVer(b *testing.B) {
	b.ReportAllocs()
	var a gosl.Ver = "Slack Bot Interface 1.2.3-45"

	// var major, minor, patch, build int
	for i := 0; i < b.N; i++ {
		a.Name()
		a.Parse()
	}
}

func TestVer_Clean(t *testing.T) {
	test := func(a gosl.Ver, exp string) {
		a = a.Clean()
		gosl.Test(t, exp, a.String())
	}

	test("Gon-World", "Gon v0.0.0-0")                  // Note that trailing space is required
	test("Gon-World ", "Gon-World v0.0.0-0")           // Note that trailing space is required
	test("Gon", "Gon v0.0.0-0")                        // Note that trailing space is not required
	test("Gon1", "Gon1 v0.0.0-0")                      // Note that trailing space is not required
	test("Gon World Here ", "Gon World Here v0.0.0-0") // Note that trailing space is required
	test("Gon 1", "Gon v1.0.0-0")
	test("Gon 1.2.3", "Gon v1.2.3-0")
	test("Gon v1", "Gon v1.0.0-0")
	test("Gon v1.2.3", "Gon v1.2.3-0")
	test("Gon World 1", "Gon World v1.0.0-0")
	test("Gon World v1", "Gon World v1.0.0-0")
	test("Gon World v1.2.3", "Gon World v1.2.3-0")
}

func BenchmarkVer_Clean(b *testing.B) {
	b.ReportAllocs()
	var a1 gosl.Ver = "Gon 1-123"

	// var major, minor, patch, build int
	for i := 0; i < b.N; i++ {
		a1 = a1.Clean()
	}
	// println(a1)
}
