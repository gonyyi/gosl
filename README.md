# Go Small Library (GoSL)

Copyright Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>

Go Small Library is a collection of frequently used functions. There are two goals for this library. First, minimal
memory allocation. Although target is to have zero memory allocation, dealing with such as a string conversion requires
unavoidable memory allocation. To minimize it, GoSL have a global level byte buffer pool which can be used by the end
user as well as library itself.

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

## GetBuffer()

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
	buf.Free()
}
```

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

## Logger

This is a very minimal logger -- it has very few methods. For logging methods, there are only two:

- `Write(p []byte) (n int, err error)`
- `WriteString(s string) (n int, err error)`

However, GoSL's logger has one very distinct feature -- initialization is not required, and will not cause a panic.

```go
package main

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"os"
)

type Something struct {
	Some  string
	Thing int
	Log   gosl.Logger
}

func main() {
	// Create an object `Something`, but did not initialized the logger.
	// However, when the logger is being used, it will not cause any panic,
	// and it will run very fast. This logger will be useful as a diagnosis
	// for small libraries where you don't want to use a heavy logger.
	s := Something{}
	s.Log.WriteString("created something") // this will not cause any panic

	s.Log = s.Log.SetOutput(os.Stdout, false)
	// Below will print:
	// > set logger's output without newline
	// > test A-1test A-2
	// >
	s.Log.WriteString("set logger's output")
	s.Log.WriteString(" without newline\n")
	s.Log.Write([]byte("test A-1"))
	s.Log.Write([]byte("test A-2"))
	s.Log.WriteString("\n")

	s.Log = s.Log.SetOutput(os.Stdout, true)
	// Below will print:
	// > set logger's output
	// >  without newline
	// > test B-1
	// > test B-2
	// >
	s.Log.WriteString("set logger's output")
	s.Log.WriteString(" without newline")
	s.Log.Write([]byte("test B-1"))
	s.Log.Write([]byte("test B-2"))

	// Logger also is compatible as a io.Writer.
	// NOTE: since `newline` was set to true, `fmt.Fprintf()` below without
	//   a newline will still append a newline. If there was a newline at the
	//   end, it will not append another newline.
	// Output:
	// > [Info] hello test 1
	// > [Info] hello test 2
	// > [Info] hello test 3
	fmt.Fprintf(s.Log, "[%s] %s", "Info", "hello test 1")
	fmt.Fprintf(s.Log, "[%s] %s", "Info", "hello test 2\n")
	fmt.Fprintf(s.Log, "[%s] %s", "Info", "hello test 3")
}
```

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

