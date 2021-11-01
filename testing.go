// Tester is an interface for testing.T

package gosl

type Tester interface {
	Name() string
	Fail()
}

func TestString(t Tester, expected, actual string) {
	if expected != actual {
		var buf []byte
		for _, v := range t.Name() + "() -> EXP=\"" + expected + "\", ACT=\"" + actual + "\"" {
			if v == '\n' {
				buf = append(buf, "\\n"...)
			} else {
				buf = append(buf, byte(v))
			}
		}
		println(string(buf))
		t.Fail()
	}
}

func TestInt(t Tester, expected, actual int) {
	if expected != actual {
		var buf []byte
		for _, v := range t.Name() + "() -> EXP=" + Itoa(expected,false) + ", ACT=" + Itoa(actual, false) + "" {
			if v == '\n' {
				buf = append(buf, "\\n"...)
			} else {
				buf = append(buf, byte(v))
			}
		}
		println(string(buf))
		t.Fail()
	}
}

func TestBool(t Tester, expected, actual bool) {
	if expected != actual {
		var buf []byte
		exp, act := "false", "false"
		if expected == true {
			exp = "true"
		}
		if actual == true {
			act = "true"
		}
		for _, v := range t.Name() + "() -> EXP=" + exp + ", ACT=" + act + "" {
			if v == '\n' {
				buf = append(buf, "\\n"...)
			} else {
				buf = append(buf, byte(v))
			}
		}
		println(string(buf))
		t.Fail()
	}
}
