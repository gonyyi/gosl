// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/9/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

type tester bool

func (t *tester) Name() string {
	return "FakeTester"
}

func (t tester) Failed() bool {
	return bool(t)
}

func (t *tester) Reset() {
	*t = false
}

func (t *tester) Fail() {
	*t = true
}

func TestTesting(t *testing.T) {
	var x tester

	t.Run("bool", func(t *testing.T) {
		x.Reset()
		gosl.Test(&x, false, false)
		if x == true {
			t.Fail()
		}
	})
	t.Run("byte", func(t *testing.T) {
		x.Reset()
		gosl.Test(&x, 'a', 'a')
		if x == true {
			t.Fail()
		}
	})

	t.Run("int", func(t *testing.T) {
		x.Reset()
		gosl.Test(&x, 123, 123)
		if x == true {
			t.Fail()
		}
	})
	t.Run("int64", func(t *testing.T) {
		x.Reset()
		gosl.Test(&x, int64(123), int64(123))
		if x == true {
			t.Fail()
		}
	})
	t.Run("string", func(t *testing.T) {
		x.Reset()
		gosl.Test(&x, "abc", "abc")
		if x == true {
			t.Fail()
		}
	})

	t.Run("float", func(t *testing.T) {
		x.Reset()
		gosl.Test(&x, float64(123.123), float64(123.123))
		if x == true {
			t.Fail()
		}
	})

	t.Run("[]string - ok", func(t *testing.T) {
		x.Reset()
		a := []string{"a", "b", "c"}
		b := []string{"a", "b", "c"}
		gosl.Test(&x, a, b)
		if x == true {
			t.Fail()
		}
	})

	t.Run("[]int - ok", func(t *testing.T) {
		x.Reset()
		a := []int{1, 3, 5}
		b := []int{1, 3, 5}
		gosl.Test(&x, a, b)
		if x == true {
			t.Fail()
		}
	})
}

func TestTesting_Fails(t *testing.T) {
	var x tester
	var skipAll = true

	t.Run("bool", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		gosl.Test(&x, false, true)
		if x == false {
			t.Fail()
		}
	})
	t.Run("byte", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		gosl.Test(&x, 'a', 'b')
		if x == false {
			t.Fail()
		}
	})

	t.Run("int", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		gosl.Test(&x, 123, 1234)
		if x == false {
			t.Fail()
		}
	})
	t.Run("int64", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		gosl.Test(&x, int64(123), int64(1234))
		if x == false {
			t.Fail()
		}
	})
	t.Run("string", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		gosl.Test(&x, "abc", "abcd")
		if x == false {
			t.Fail()
		}
	})

	t.Run("float", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		gosl.Test(&x, float64(123.123), float64(123.1234))
		if x == false {
			t.Fail()
		}
	})

	t.Run("[]string - diff size", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}

		x.Reset()
		a := []string{"a", "b", "c"}
		b := []string{"a", "b", "c", "d"}
		gosl.Test(&x, a, b)
		if x == false {
			t.Fail()
		}
	})

	t.Run("[]string - diff content", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}
		x.Reset()
		a := []string{"a", "b", "c"}
		b := []string{"a", "d", "c"}
		gosl.Test(&x, a, b)
		if x == false {
			t.Fail()
		}
	})

	t.Run("[]int - diff size", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}
		x.Reset()
		a := []int{1, 3, 5}
		b := []int{1, 3, 5, 6}
		gosl.Test(&x, a, b)
		if x == false {
			t.Fail()
		}
	})

	t.Run("[]int - diff content", func(t *testing.T) {
		if skipAll {
			t.SkipNow()
		}
		x.Reset()
		a := []int{1, 3, 5}
		b := []int{1, 3, 4}
		gosl.Test(&x, a, b)
		if x == false {
			t.Fail()
		}
	})
}

