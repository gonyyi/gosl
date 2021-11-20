# Go Small Library (gosl)

Copyright Gon Y. Yi 2021 <https://gonyyi.com/copyright>


## Goal

__General__

- No import of any library whatsoever including standard library.
- Most of the code should have zero memory allocation.
- Only frequently used functions.
- Only very minimum functions.
- Safe code
    - 99%+ code coverage
    - All the code should have tests, benchmarks and examples
- Minimize default allocation caused by importing the library
    - Currently `bufp` is allocated at global level when importing the library
      as this is required for the logger. 


## append.go

- `AppendBool(dst []byte, b bool) []byte`
- `AppendInt(dst []byte, i int, comma bool) (b []byte)`
- `AppendFloat64(dst []byte, value float64, decimal uint8, comma bool) []byte`
- `AppendString(dst []byte, s string, trim bool) []byte`
- `AppendPath(dst []byte, path ...string) []byte`
- `AppendFills(dst []byte, src []byte, n int) []byte`
- `AppendRepeats(dst []byte, rep []byte, n int) []byte`
- `AppendRepeat(dst []byte, rep byte, n int) []byte`
- `AppendStringFit(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte`
- `AppendStringFitCenter(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte`
- `AppendStringFitRight(dst []byte, s string, length int, filler byte, overflowMarker bool) []byte`


## bitflag.go

- `BitsAdd(b1, b2 uint32) uint32`
- `BitsSub(bFrom, bTo uint32) uint32`
- `BitsToggle(bFrom, bToggle uint32) uint32`
- `BitsAnd(b1, b2 uint32) uint32`
- `BitsAny(bFrom, bTo uint32) bool`
- `BitsHas(b1, b2 uint32) bool`
- `NewBitflag() Bitflag`
- `type Bitflag uint32`
    - `(Bitflag) All() Bitflag`
    - `(Bitflag) None() Bitflag`
    - `(Bitflag) Nth(Nth ...uint8) Bitflag`
    - `(Bitflag) Reverse() Bitflag`
    - `(Bitflag) Add(b Bitflag) Bitflag`
    - `(Bitflag) Sub(b Bitflag) Bitflag`
    - `(Bitflag) Toggle(b Bitflag) Bitflag`
    - `(Bitflag) And(b Bitflag) Bitflag`
    - `(Bitflag) Any(b Bitflag) bool`
    - `(Bitflag) Has(b Bitflag) bool`
    - `(Bitflag) Output(dst []byte) []byte`


## buf.go

- `type Buf []byte`
    - `(Buf) WriteBytes(bytes ...byte) Buf`
    - `(Buf) WriteBool(t bool) Buf`
    - `(Buf) WriteInt(i int) Buf`
    - `(Buf) WriteFloat64(f64 float64) Buf`
    - `(Buf) WriteString(s string) Buf`
    - `(Buf) Last() byte`
    - `(Buf) Trim(n uint) Buf`
    - `(Buf) Cap() int`
    - `(Buf) Len() int`
    - `(Buf) Reset() Buf`
    - `(Buf) String() string`
    - `(Buf) WriteTo(w Writer) (n int, err error)`
    - `(Buf) Println()`
    - `(*Buf) Write(p []byte) (n int, err error)`


## bufPool.go

- `GetBuffer() *poolBuf`
    - `(*poolBuf) Free()`
    - `(*poolBuf) Init(size int)`
    - `(*poolBuf) Write(p []byte) (n int, err error)`
    - `(*poolBuf) WriteBytes(a ...byte) *poolBuf`
    - `(*poolBuf) WriteBool(t bool) *poolBuf`
    - `(*poolBuf) WriteInt(i int) *poolBuf`
    - `(*poolBuf) WriteFloat64(f64 float64) *poolBuf`
    - `(*poolBuf) WriteString(s string) *poolBuf`
    - `(*poolBuf) Last() byte`
    - `(*poolBuf) Trim(n uint)`
    - `(*poolBuf) Cap() int`
    - `(*poolBuf) Len() int`
    - `(*poolBuf) Reset()`
    - `(*poolBuf) String() string`
    - `(*poolBuf) Bytes() []byte`
    - `(*poolBuf) WriteTo(w Writer) (n int, err error)`


## bytes.go

- `NewBytesFilter(allow bool, list []byte) func([]byte) []byte`
- `BytesInsert(dst []byte, index int, p []byte) []byte`
- `BytesReverse(dst []byte) []byte`
- `BytesToUpper(dst []byte) []byte`
- `BytesToLower(dst []byte) []byte`


## conv.go 

- `Itoa(i int) (s string)`
- `Itoaf(i int, comma bool) (s string)`
- `Ftoa(f64 float64) (s string)`
- `Ftoaf(f64 float64, decimal uint8, comma bool) (s string)`
- `Atoi(s string) (num int, ok bool)`
- `MustAtoi(s string, fallback int) int`
- `ToUpper(s string) string`
- `ToLower(s string) string`


## err.go

- `NewError(s string) error`
- `IfErr(key string, e error)`
- `IfPanic(name string, f func(error))`


## gosl.go

- `DoNothing()`


## int.go

- `IntsJoin(dst []byte, p []int, delim ...byte) []byte`


## logger.go

- `NewLogger(w Writer) Logger`
- `type Logger`
    - `(l Logger) SetPrefix(prefix6 string) Logger`
    - `(l Logger) SetOutput(w Writer) Logger`
    - `(l Logger) Enable(enable bool) Logger`
    - `(l Logger) IfErr(key string, err error) (isError bool)`
    - `(l Logger) ifErr(key string, err error)`
    - `(l Logger) write(p []byte)`
    - `(l Logger) String(s string)`
    - `(l Logger) string(s string)`
    - `(l Logger) KeyBool(key string, val bool)`
    - `(l Logger) keyBool(key string, val bool)`
    - `(l Logger) KeyInt(key string, val int)`
    - `(l Logger) keyInt(key string, val int)`
    - `(l Logger) KeyFloat64(key string, val float64)`
    - `(l Logger) keyFloat64(key string, val float64)`
    - `(l Logger) KeyString(key string, val string)`
    - `(l Logger) keyString(key string, val string)`
    - `(l Logger) KeyError(key string, val error)`
    - `(l Logger) keyError(key string, err error)`
    - `(l Logger) Write(p []byte) (n int, err error)`
    - `(l Logger) Enabled() bool`
    - `(l Logger) Close() error`


## mutex.go

- `type Mutex chan struct{}`
    - `(Mutex) Init() Mutex`
    - `(Mutex) Unlock()`
    - `(Mutex) Lock()`
    - `(Mutex) LockFor(f func())`
    - `(Mutex) LockIfNot() (ok bool)`
    - `(Mutex) Locked() bool`

- `type Once chan struct{}`
    - `(Once) Init(n int) Once`
    - `(Once) Reset()`
    - `(Once) Do(f func()) bool`
    - `(Once) Available() int`
    - `(Once) Close() error`

- `type MuBool struct`
    - `(MuBool) Init() MuBool`
    - `(*MuBool) Get() (t bool)`
    - `(*MuBool) Set(t bool)`

- `type MuInt struct`
    - `(MuInt) Init() MuInt`
    - `(*MuInt) Get() (n int)`
    - `(*MuInt) Add(n int)`


## pool.go

- `type Pool`
    - `(Pool) Init(size int) Pool`
    - `(*Pool) Get() interface{}`
    - `(*Pool) Put(b interface{})`
    

## sort.go

- `SortSlice(dst []interface{}, compare func(idx1, idx2 int) bool) (ok bool)`
- `SortStrings(dst []string, compare func(idx1, idx2 int) bool) (ok bool)`
- `SortInts(dst []int)`


## string.go

- `IsNumber(s string) bool`
- `Split(dst []string, s string, delim rune) []string`
- `Join(dst []byte, p []string, delim ...byte) []byte`
- `Trim(s string) string`
- `TrimLeft(s string) string`
- `TrimRight(s string) string`
- `trim(s string, trimLeft, trimRight bool) string`
- `HasPrefix(s string, prefix string) bool`
- `HasSuffix(s string, suffix string) bool`
- `TrimPrefix(s string, prefix string) string`
- `TrimSuffix(s string, suffix string) string`


## tester.go

- `Test(t interface{}, expected, actual interface{})`


## ver.go

- `NewVer(name string, major, minor, patch, build int) Ver`
- `type Ver string`
    - `(v Ver) String() string`
    - `(v Ver) Name() string`
    - `(v Ver) Version() (major, minor, patch, build int)`
    - `(v Ver) IsNewer(old Ver) bool`
    - `(v Ver) Set(name string, major, minor, patch, build int) Ver`
    - `(v Ver) Clean() Ver`
    - `(v Ver) Parse() (name string, major, minor, patch, build int)`


## writer.go

- `type Writer interface`
- `type Closer interface`
- `var Discard writer`
    - `(discard) Write(p []byte) (n int, err error)`
- `Close(w interface{}) error`
- `NewAltWriter(dst Writer, f func([]byte) []byte) Writer`
    - `(altWriter) Write(b []byte) (n int, err error)`
    - `(altWriter) Close() error`
- `NewPrefixWriter(prefix string, w Writer) Writer`
- `NewMultiWriter(w ...Writer) Writer`


## runner/

### job.go

- `type Job interface`
    - `ID()string`
    - `Accept()`
    - `Reject()`
    - `Cancel()`
    - `Run()`
- `NewJob(ID string, fnRun func()) *simpleJob`
    - `(*simpleJob) SetID(id string) *simpleJob`
    - `(*simpleJob) SetReject(f func()) *simpleJob`
    - `(*simpleJob) SetAccept(f func()) *simpleJob`
    - `(*simpleJob) SetCancel(f func()) *simpleJob`
    - `(*simpleJob) SetRun(f func()) *simpleJob`
    - `(*simpleJob) ID() string`
    - `(*simpleJob) Accept()`
    - `(*simpleJob) Reject()`
    - `(*simpleJob) Cancel()`
    - `(*simpleJob) Run()`


### runner.go

- `func NewRunner(queue uint, workers uint, async bool, finish bool) *Runner`
    - `func (b *Runner) Stats() (rejected, accepted, cancelled, completed int)`
    - `func (b *Runner) String() string`
    - `func (b *Runner) SetLoggerOutput(debug gosl.Writer)`
    - `func (b *Runner) Closed() bool`
    - `func (b *Runner) Stop()`
    - `func (b *Runner) WaitClose()`
    - `func (b *Runner) Add(f Job) (ok bool)`
    - `func (b *Runner) Run() *Runner`
    - `func (b *Runner) Queue() int`
    - `func (b *Runner) Running() int`
    - `func (b *Runner) add(f Job) (ok bool)`
    - `func (b *Runner) run()`
    - `func (b *Runner) waitToFinish()`
    - `func (b *Runner) close()`



