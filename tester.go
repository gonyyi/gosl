// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/29/2021

package gosl

func Test(t interface{}, expected, actual interface{}, whenFail ...func()) {
	tx, ok := t.(interface {
		Name() string
		Fail()
		Failed() bool
	})
	if !ok {
		println("gosl.Test(): unexpected t given")
		return
	}

	buf := make(Buf, 0, 2048)
	buf = buf.WriteString(tx.Name()).
		WriteString("()\n\tEXP => ")
	switch exp := expected.(type) {
	case bool:
		act, ok := actual.(bool)
		if !ok {
			buf = buf.WriteBool(exp).WriteString(" (bool)").
				WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}

		if exp != act {
			buf = buf.WriteBool(exp).WriteString(" (bool)").
				WriteString("\n\tACT => ").
				WriteBool(act).
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case rune:
		exps := Itoa(int(exp))
		act, ok := actual.(rune)
		if !ok {
			buf = buf.WriteString(exps).WriteString(" (rune)").
				WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		acts := Itoa(int(act))
		if exp != act {
			buf = buf.WriteString(exps).
				WriteBytes('(').WriteBytes(byte(exp)).WriteBytes(')').
				WriteString(" (rune)").
				WriteString("\n\tACT => ").
				WriteString(acts).
				WriteBytes('(').WriteBytes(byte(act)).WriteBytes(')').
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case byte:
		exps := Itoa(int(exp))
		act, ok := actual.(byte)
		if !ok {
			buf = buf.WriteString(exps).WriteString(" (byte)").
				WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		acts := Itoa(int(act))
		if exp != act {
			buf = buf.WriteString(exps).
				WriteBytes('(').WriteBytes(exp).WriteBytes(')').
				WriteString(" (byte)").
				WriteString("\n\tACT => ").
				WriteString(acts).
				WriteBytes('(').WriteBytes(act).WriteBytes(')').
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case int64:
		buf = buf.WriteInt(int(exp)).WriteString(" (int64)")
		act, ok := actual.(int64)
		if !ok {
			buf = buf.WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		if int(exp) != int(act) {
			buf = buf.WriteString("\n\tACT => ").
				WriteInt(int(act)).
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case int:
		buf = buf.WriteInt(exp).WriteString(" (int)")
		act, ok := actual.(int)
		if !ok {
			buf = buf.WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		if exp != act {
			buf = buf.WriteString("\n\tACT => ").
				WriteInt(act).
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case string:
		act, ok := actual.(string)
		if !ok {
			buf = buf.WriteString(exp).WriteString(" (string)").
				WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		if exp != act {
			buf = buf.WriteString(exp).WriteString(" (string)").
				WriteString("\n\tACT => ").
				WriteString(act).
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case float64:
		buf = buf.WriteFloat64(exp).WriteString(" (float64)")
		act, ok := actual.(float64)
		if !ok {
			buf = buf.WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		if exp != act {
			buf = buf.WriteString("\n\tACT => ").
				WriteFloat64(act).
				WriteBytes('\n')
			print(buf.String())
			tx.Fail()
		}
	case []int:
		buf = buf.WriteBytes('[')
		buf = IntsJoin(buf, exp, ',')
		buf = buf.WriteBytes(']')
		act, ok := actual.([]int)
		if !ok {
			buf = buf.WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		isMatch := true
		if len(exp) != len(act) {
			isMatch = false
		}
		if isMatch {
			for i := 0; i < len(exp); i++ {
				if exp[i] != act[i] {
					isMatch = false
					break
				}
			}
		}
		if isMatch == false {
			buf = buf.WriteString("\n\tACT => [")
			buf = IntsJoin(buf, act, ',')
			buf = buf.WriteBytes(']', '\n')
			print(buf.String())
			tx.Fail()
		}
	case []string:
		buf = buf.WriteBytes('[')
		buf = Join(buf, exp, ',')
		buf = buf.WriteBytes(']')
		act, ok := actual.([]string)
		if !ok {
			buf = buf.WriteString("\n\tACT => (err) Unexpected-Type\n")
			print(buf.String())
			tx.Fail()
		}
		isMatch := true
		if len(exp) != len(act) {
			isMatch = false
		}
		if isMatch {
			for i := 0; i < len(exp); i++ {
				if exp[i] != act[i] {
					isMatch = false
					break
				}
			}
		}
		if isMatch == false {
			buf = buf.WriteString("\n\tACT => [")
			buf = Join(buf, act, ',')
			buf = buf.WriteBytes(']', '\n')
			print(buf.String())
			tx.Fail()
		}

	default:
		print("(err) Unsupported-Type")
		tx.Fail()
	}

	// If failed, run all optional whenFail functions
	if tx.Failed() {
		for _, f := range whenFail {
			f()
		}
	}
}

