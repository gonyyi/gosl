package gosl_test

import (
	"github.com/gonyyi/gosl"
	"os"
	"time"
)

func ExampleLvWriter_WriteAny() {
	// WriteAny is a special method - unlike other methods used,
	// WriteAny takes variadic params of interface{}. Currently, for
	// the simplicity, only few data types are supposed:
	//     `string`, `int`, `bool`, `[]byte`, `func([]byte)[]byte`
	// `func([]byte)[]byte` is a magic key for LvWriter as it can be
	// used to add current time OR header of LvWriter - making it
	// a full function logger yet zero or minimum memory allocation.

	// LvWriter can be created multiple way, but `NewLvWriter()` is
	// simply shortcut for `LvWriter{}.SetOutput().SetLevel()` with
	// one command.
	lw := gosl.NewLvWriter(os.Stdout, 0) // 0 for lvl is lowest level

	// Unlike builtin println, WriteAny() will not add a space between params.
	// Also note that WriteAny() and WriteString() will check if there's
	// newline at the end of input, if missing a newline, it will append a
	// newline.
	lw.WriteAny("my", "name", "is", "gon")  // using string
	lw.WriteAny("i am ", 100, " years old") // using number

	head := func(dst []byte) []byte {
		return append(dst, "[MyLog] "...)
	}
	lw.WriteAny(head, "my", "name", "is", "gon") // anonymous function head is the first argument,
	lw.WriteAny(head, "i am ", 150, " years old")

	// Any method in LvWriter can have level. Either by `LwWriter{}.Lv(LvLevel)...`
	// or `LwWriter{}.Info()...`, `LwWriter{}.Error()...`.
	// Let's reset LvWriter with Info level.
	lw = lw.SetLevel(gosl.LvInfo)                        // Note that `lw = lw.SetLevel()...` instead of just `lw.SetLevel()...`
	lw.Debug().WriteAny(head, "my", "name", "is", "gon") // will not be printed as current minimum level is Info.
	lw.Error().WriteAny(head, "i am ", 200, " years old")

	// Now, let's create a function that prints current time to the log.
	// for the test, instead of getting current time, use static time as below
	var now time.Time
	showTime := func(dst []byte) []byte {
		// now = time.Now()
		now, _ = time.Parse("20060102150405", "20220202010200")
		return now.AppendFormat(dst, "2006/01/02 Mon 15:04:05 ")
	}

	// print current time; then, head record.
	lw.Debug().WriteAny(showTime, head, "my", "name", "is", "gon")
	lw.Fatal().WriteAny(showTime, head, "i am ", 300, " years old")

	// User can further create a closure as below:
	MyDebug := func(name string, age int) {
		lw.Debug().WriteAny(showTime, "[DBG] ", name, " is ", age, " years old")
	}
	MyError := func(name string, age int) {
		lw.Error().WriteAny(showTime, "[ERR] ", name, " is ", age, " years old")
	}

	MyDebug("Gon", 40) // Will not print, because level (debug) is below minimum (info)
	MyError("Don", 90)

	// Output:
	// mynameisgon
	// i am 100 years old
	// [MyLog] mynameisgon
	// [MyLog] i am 150 years old
	// [MyLog] i am 200 years old
	// 2022/02/02 Wed 01:02:00 [MyLog] i am 300 years old
	// 2022/02/02 Wed 01:02:00 [ERR] Don is 90 years old
}

func ExampleNewLvWriter() {
	// LvWriter is not initiated yet, but using it won't cause any panic.
	var lw gosl.LvWriter
	lw.WriteString("hello 1") // This won't cause panic.

	lw = lw.SetOutput(os.Stdout) // Now output will be to stdout
	lw.WriteString("hello 2")    // This will be printed

	// By setting writer with LvWarn level,
	// now only record with level LvWarn or above will be printed.
	lw = lw.SetLevel(gosl.LvWarn)
	lw.Trace().WriteString("this is trace")
	lw.Debug().WriteString("this is debug")
	lw.Info().WriteString("this is info")
	lw.Warn().WriteString("this is warn")   // Will be printed
	lw.Error().WriteString("this is error") // Will be printed
	lw.Fatal().WriteString("this is fatal") // Will be printed

	// Instead of Trace(), Debug(), etc., a user can use Lv(LvLevel) method as well.
	// This Lv() method can be used when custom log level (LvLevel) is being used.
	// Since LvLevel is just an alias for uint8, any uint8 or alias of uint8 can be
	// used.
	lw.Lv(gosl.LvInfo).WriteString("another way of info")   // WILL NOT BE PRINTED
	lw.Lv(gosl.LvWarn).WriteString("another way of warn")   // will be printed
	lw.Lv(gosl.LvError).WriteString("another way of error") // will be printed

	// Output:
	// hello 2
	// this is warn
	// this is error
	// this is fatal
	// another way of warn
	// another way of error
}
