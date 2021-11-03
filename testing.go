// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

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

func TestByte(t Tester, expected, actual byte) {
	if expected != actual {
		buf := getBufpBuffer()
		buf.WriteString(t.Name())
		buf.WriteString("() -> EXP=byte(")
		buf.Buf = AppendInt(buf.Buf, int(expected), false)
		buf.WriteString("), ACT=byte(")
		buf.Buf = AppendInt(buf.Buf, int(actual), false)
		buf.WriteString(")")
		println(buf.String())
		buf.ReturnBuffer()
		t.Fail()
	}
}

func TestInt(t Tester, expected, actual int) {
	if expected != actual {
		var buf []byte
		for _, v := range t.Name() + "() -> EXP=" + Itoaf(expected, false) + ", ACT=" + Itoaf(actual, false) + "" {
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


