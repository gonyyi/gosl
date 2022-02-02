# Go Small Library (Gosl)

(c) Gon Y. Yi 2021-2022

<https://gonn.org> / <https://gonyyi.com>

![Gosl Mascot](gosl.png "Gosl")

Go Small Library is a collection of frequently used functions. There are two goals for this library. First, minimal
memory allocation. Although target is to have zero memory allocation, sometimes, dealing with such as a string
conversion requires unavoidable memory allocation. To minimize it, GoSL have a global level byte buffer pool which can
be used by the end user as well as library itself.

1. Gosl is entirely built only with Golang's built-in keywords (Except for the testing code)
2. Gosl does not allocate memory
3. Gosl should be fully compatible with TinyGo for microprocessors.

Table of Contents

- [Benchmark](#benchmark)
- [LvWriter](#lvwriter)
- [Buf](#buf)
	- [GetBuffer()](#getbuffer)
- [Bytes](#bytes)
- [Mutex](#mutex)


## Benchmark

- As of v0.7.7
- GOOS: darwin
- GOARCH: arm64
- PKG: github.com/gonyyi/gosl/tests

| Benchmark                                        | # of run   | time/op   | b/op | alloc/op |             
| :----------------------------------------------- | ---------: | --------: | ---: | -------: | 
| BenchmarkBuf/Write()-8                           |   64397343 |  18.65 ns |  0 B |  0 alloc |
| BenchmarkNewBuffer-8                             |   91040132 |  16.50 ns |  0 B |  0 alloc |
| BenchmarkBuffer/T1-8                             |   25762336 |  51.18 ns |  0 B |  0 alloc |
| BenchmarkBytesAppends/BytesAppendBool-8          | 1000000000 | 0.8354 ns |  0 B |  0 alloc |
| BenchmarkBytesAppends/BytesAppendInt-8           |   99758914 |  15.56 ns |  0 B |  0 alloc |
| BenchmarkBytesAppends/BytesAppendPrefix-8        |  100000000 |  12.29 ns |  0 B |  0 alloc |
| BenchmarkBytesAppends/BytesAppendPrefixString-8  |  100000000 |  11.33 ns |  0 B |  0 alloc |
| BenchmarkBytesAppends/BytesAppendSuffix-8        |  624065390 |  1.916 ns |  0 B |  0 alloc |
| BenchmarkBytesAppends/BytesAppendSuffixString-8  |  638502922 |  1.875 ns |  0 B |  0 alloc |
| BenchmarkBytesEquals/BytesEqual-8                |  313417296 |  3.765 ns |  0 B |  0 alloc |
| BenchmarkBytesEquals/BytesEqual:false-8          | 1000000000 | 0.3579 ns |  0 B |  0 alloc |
| BenchmarkBytesFilterAny/keep=true-8              |   68192146 |  17.12 ns |  0 B |  0 alloc |
| BenchmarkBytesFilterAny/keep=false-8             |   70628917 |  16.91 ns |  0 B |  0 alloc |
| BenchmarkBytesHases/BytesHasPrefix-8             |  625212474 |  1.905 ns |  0 B |  0 alloc |
| BenchmarkBytesHases/BytesHasPrefixString-8       |  762768020 |  1.569 ns |  0 B |  0 alloc |
| BenchmarkBytesHases/BytesHasSuffix-8             |  955929650 |  1.253 ns |  0 B |  0 alloc |
| BenchmarkBytesHases/BytesHasSuffixString-8       |  957725401 |  1.251 ns |  0 B |  0 alloc |
| BenchmarkBytesIndexes/BytesIndex-8               |  209662804 |  5.690 ns |  0 B |  0 alloc |
| BenchmarkBytesIndexes/BytesIndexString-8         |  153219141 |  7.569 ns |  0 B |  0 alloc |
| BenchmarkBytesInserts/BytesInsert-8              |  140897906 |  8.494 ns |  0 B |  0 alloc |
| BenchmarkBytesInserts/BytesInsertString-8        |  147377509 |  8.286 ns |  0 B |  0 alloc |
| BenchmarkBytesTos/BytesToLower-8                 |  235865358 |  5.086 ns |  0 B |  0 alloc |
| BenchmarkBytesTos/BytesToUpper-8                 |  239079325 |  5.014 ns |  0 B |  0 alloc |
| BenchmarkBytesTrims/BytesTrimPrefix-8            |  271787853 |  4.446 ns |  0 B |  0 alloc |
| BenchmarkBytesTrims/BytesTrimPrefixString-8      |  288879405 |  4.118 ns |  0 B |  0 alloc |
| BenchmarkBytesTrims/BytesTrimSuffix-8            |  338991747 |  3.568 ns |  0 B |  0 alloc |
| BenchmarkBytesTrims/BytesTrimSuffixString-8      |  303717277 |  3.945 ns |  0 B |  0 alloc |
| BenchmarkBytesElem-8                             |  152857048 |  7.945 ns |  0 B |  0 alloc |
| BenchmarkBytesEtc/BytesLastByte-8                |  452576361 |  2.659 ns |  0 B |  0 alloc |
| BenchmarkBytesEtc/BytesReplace-8                 |   60801052 |  19.64 ns |  0 B |  0 alloc |
| BenchmarkBytesEtc/BytesReverse-8                 |  252620676 |  4.639 ns |  0 B |  0 alloc |
| BenchmarkBytesEtc/BytesShift-8                   |  149032028 |  8.036 ns |  0 B |  0 alloc |
| BenchmarkIfPanic/basic-8                         |  544778833 |  2.204 ns |  0 B |  0 alloc |
| Benchmark_Sort_SortInts-8                        |   64256167 |  18.10 ns |  0 B |  0 alloc |
| Benchmark_Sort_SortStrings-8                     |   17430354 |  67.58 ns |  0 B |  0 alloc |
| Benchmark_String_Atoi/basic-8                    |  126459487 |  9.423 ns |  0 B |  0 alloc |
| Benchmark_String_Itoa/Plain-8                    |   70785496 |  20.59 ns |  0 B |  0 alloc |
| Benchmark_Mutex/Mutex/LockUnlock-8               |   44735431 |  26.44 ns |  0 B |  0 alloc |
| Benchmark_Mutex/Mutex/LockFor-8                  |   45084532 |  26.48 ns |  0 B |  0 alloc |
| Benchmark_Mutex/Mutex/LockIfNot-8                |   42992517 |  27.33 ns |  0 B |  0 alloc |
| Benchmark_Mutex/MuInt-8                          |   44020138 |  27.23 ns |  0 B |  0 alloc |
| Benchmark_Pool/x1-8                              |   41051253 |  29.33 ns |  0 B |  0 alloc |
| Benchmark_Pool/x256-8                            |   40796031 |  28.98 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/WriteString()-8                |   28001535 |  46.75 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/WriteAny()/enabled-8           |   19065726 |  67.64 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/WriteAny()/disabled-8          |  254944800 |  4.687 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/WriteAny()+timeFunc/enabled-8  |    4774202 |  253.4 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/WriteAny()+timeFunc/disabled-8 |  306364662 |  3.915 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/Write():_enabled-8             |  167933073 |  7.440 ns |  0 B |  0 alloc |
| BenchmarkLvWriter/Write():_disabled-8            | 1000000000 | 0.8071 ns |  0 B |  0 alloc |

^[Top](#go-small-library-gosl)


## LvWriter

`LvWriter` is a multifunctional light writer wrapper which can be used as a levelled logger as well.

Basic usage: (see `example_writer_test.go`)

```go
package main

import (
	"github.com/gonyyi/gosl"
	"os"
)

func main() {
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
```

`WriteAny()` is a new method added to `LvWriter{}`. Usage is as below:

```go
package main

import (
	"github.com/gonyyi/gosl"
	"os"
	"time"
)

func main() {
	// Newly added WriteAny() is a special method - unlike other methods 
	// used, WriteAny() takes variadic params of interface{}. Currently, 
	// for the simplicity, only few data types are supposed:
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
	lw.Debug().WriteAny(showTime, head, "my", "name", "is", "gon") // Will not print because of level
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
```

^[Top](#go-small-library-gosl)


## Buf

`Buf` is a struct for byte slice with many useful methods. To avoid potential leak, it is recommended to create one with
a pre-defined size.

NOTE: `Buf` is fully compatible with `[]byte`

```go
package main

import "github.com/gonyyi/gosl"

func main() {
	// create Buf with 1k
	buf := make(gosl.Buf, 0, 1024)

	// write string `test` to buffer.
	buf = buf.WriteString("test")
	buf.Println() // prints "test"

	// Buf is compatible with []byte
	buf = append(buf, "-1"...)
	buf.Println() // prints "test-1"
}
```

^[Top](#go-small-library-gosl)


### GetBuffer()

Also a buffer pool is created upon importing the GoSL library. This can be called by `gosl.GetBuff()`.

```go
package main

import "github.com/gonyyi/gosl"

func main() {
	// Get a Buf from the pool
	buf := gosl.GetBuffer()

	// Write a string `test` and integer `123` to the buffer
	buf.Buf = buf.Buf.WriteString("test").WriteInt(123)
	buf.Buf.Println() // prints "test123"

	// Now set the buf with "abc"
	buf.Buf = buf.Buf.Set("abc")
	println(buf.Buf.String()) // prints "abc"

	// Return the buffer back to pool. 
	// NOTE: if this buffer wasn't freed, there will be a memory allocation.
	gosl.PutBuffer(buf)
}
```

^[Top](#go-small-library-gosl)


## Bytes

There are many standalone functions for byte slices. They all have prefix
`Bytes` in its name such as `BytesAppendInt()`

```go
package main

import "github.com/gonyyi/gosl"

func main() {
	// Create a byte slice for a string "Hello"
	// And convert it to uppercase
	tmp := []byte("Hello")
	gosl.BytesToUpper(tmp)
	println(string(tmp)) // prints "HELLO"

	// Create a 32 byte Buf with "Hello" in it. 
	// NOTE: gosl.Buf is fully compatible with `[]byte`.
	buf := make(gosl.Buf, 0, 32).Set("Hello")
	gosl.BytesToUpper(buf)
	buf.Println() // prints "HELLO"

	// Reset the Buf as "Pi is ", and then append float 3.1415 with
	// two decimal places. 
	buf = buf.Set("Pi is ")
	buf = gosl.BytesAppendFloat(buf, 3.1415, 2)
	buf.Println() // prints "Pi is 3.14"
}
```

^[Top](#go-small-library-gosl)


## Mutex

This Mutex is created using channel. It's very simple straight forward, however is about 4 times slower than
using `sync.Mutex`.

There are two Mutex structs -- `Mutex` and `MuInt` where `Mutex` does very same thing as `sync.Mutex`, but `MuInt` is
similar to `sync.WaitGroup`.

```go
package main

import "github.com/gonyyi/gosl"

func main() {
	// There are 3 ways to create a Mutex.
	// 1. By using `NewMutex()`: mu := gosl.NewMutex()
	// 2. By using `Init()`: var mu gosl.Mutex; mu = mu.Init();
	// 3. By using `make`: mu := make(gosl.Mutex, 1)
	mu := gosl.NewMutex()
	total := 0 // total will be used to have sum from 1 to 100.

	// Since this code will run
	mi := gosl.NewMuInt()
	mi.Set(0) // unlike Mutex, MuInt uses pointer receiver.

	for i := 0; i < 100; i++ {
		// `mi` will be used as a WaitGroup
		// `*MuInt.Add()` take a change `i`. It can be any number,
		// also negative or positive integer as well.
		mi.Add(1)

		go func(i int) {
			// // as i starts from 0 and to 99, add +1 to i.
			// mu.Lock()
			// total += i + 1 
			// mi.Add(-1)
			// mu.Unlock()

			// `Mutex.LockFor()` takes a function, and wrap it 
			// with Mutex.Lock() and Mutex.Unlock(). Below code 
			// will do same as above code.
			mu.LockFor(func() {
				// as i starts from 0 and 
				// to 99, add +1 to i.
				total += i + 1
				mi.Add(-1)
			})
		}(i)
	}

	// Wait until all goroutines are finished. 
	// (until the value of `mi` become 0)
	mi.Wait(0)
	println(total) // prints 5050
}
```

^[Top](#go-small-library-gosl)

