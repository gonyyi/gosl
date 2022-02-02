# GOSLJ

(c) Gon Y. Yi 2021-2022

<https://gonn.org> / <https://gonyyi.com>

GOSLJ is part of extended GOSL package. GOSLJ is also a zero allocation JSON builder.

__Background__ 
: I was building a microservice that constantly returns JSON formatted data using Golang's built-in 
JSON library. However I noticed a memory allocations and tried to eliminate allocations by creating
a custom JSON builder that won't require struct to be created, and also won't allocate to memory.


## Usage

Simplest

```go
package main

import (
	goslj "github.com/gonyyi/gosl/json"
	"os"
)

func main() {
	goslj.NewJSON(1024).Start().
		String("name", "Gon").
		End().Write(os.Stdout)
	// Output:
	// {"name":"Gon"}
}
```

Complex - using `Sub()` method

```go
package main

import (
	goslj "github.com/gonyyi/gosl/json"
	"os"
)

func main() {
	jp := goslj.NewPool(20) // create a JSON pool with 20 objects

	j1 := jp.Get() // get JSON from the pool 
	j2 := jp.Get()
	j3 := jp.Get()

	j1.Start(). // JSON always starts with `Start()`, and ends with `End()`
		String("city", "conway").
		String("state", "arkansas").
		Int("zip", 72034).
		End() // end
	j2.Start().
		String("name", "gonn corp").
		Int("tin", 123456789).
		Int("income", 123456).
		End()
	j3.Start().
		String("name", "Gon Yi").
		Int("age", 100).
		Sub("address", j1). // add j1 into j
		Sub("employer", j2). // add j2 into j
		End().
		Write(os.Stdout) // print to screen
		// Output:
	    // {"name":"Gon Yi","age":100,"address":{"city":"conway","state":"arkansas","zip":72034},"employer":{"name":"gonn corp","tin":123456789,"income":123456}}
	
	j1.Putback() // return to pool 
	j2.Putback()
	j3.Putback()
}
```
