// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl

// Buf is a byte buffer similar to bytes.Buffer.
type Buf []byte

// Bytes will return Buf as a byte slice,
// however, this is a much trimmed down version.
func (b Buf) Bytes() []byte {
	return b
}

// Cap will return capacity of Buf
func (b Buf) Cap() int {
	return cap(b)
}

// Len will return 0 if Buf is nil,
// otherwise, it will return the size of bytes
func (b Buf) Len() int {
	return len(b)
}

// Println will print current buffer into stdout
func (b Buf) Println() Buf {
	println(b.String())
	return b
}

// Reset will reset the Buf
func (b Buf) Reset() Buf {
	return b[:0]
}

// Set will reset Buf with string s without allocation new
func (b Buf) Set(s string) Buf {
	return b.Reset().WriteString(s)
}

// String will convert Buf into string.
// String meets fmt.Stringer interface.
func (b Buf) String() string {
	return string(b)
}

// Write for io.Writer interface
func (b *Buf) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

// WriteBytes will take byte(s) and append to Buf
func (b Buf) WriteBytes(bytes ...byte) Buf {
	return append(b, bytes...)
}

func (b Buf) WriteBool(t bool) Buf {
	return BytesAppendBool(b, t)
}

func (b Buf) WriteFloat(f float64, dec uint8) Buf {
	return BytesAppendFloat(b, f, dec)
}

func (b Buf) WriteInt(i int) Buf {
	return BytesAppendInt(b, i)
}

// WriteString will take a string `s` and append to Buf
func (b Buf) WriteString(s string) Buf {
	return append(b, s...)
}

// WriteStrings will take a string slice `s` and append to Buf
// If delim is not 0, it will append delim.
func (b Buf) WriteStrings(s []string, delim ...byte) Buf {
	return BytesAppendStrings(b, s, delim...)
}

// WriteTo will take a writer and dump Buf contents into it.
// After that Buf will be reset..
// This meets io.WriterTo interface.
func (b Buf) WriteTo(w Writer) (n int, err error) {
	n, err = w.Write(b)
	b.Reset()
	return
}

// ************************************************************************************************************
// IMPLEMENTATION FROM BYTES* FUNCTIONS
// ************************************************************************************************************

// AppendPrefix will take prefix and if Buf does not have that suffix,
// it will append, otherwise return Buf.
// eg. `/abc/def` --> AppendSuffix('/') --> `/abc/def`
//     `abc/def`  --> AppendSuffix('/') --> `/abc/def`
func (b Buf) AppendPrefix(prefix ...byte) Buf {
	return BytesAppendPrefix(b, prefix...)
}

// AppendPrefixString will append string prefix
func (b Buf) AppendPrefixString(prefix string) Buf {
	return BytesAppendPrefixString(b, prefix)
}

// AppendSuffix will take suffix and if Buf does not have that suffix,
// it will append, otherwise return Buf.
// eg. `abc/def`  --> AppendSuffix('/') --> `abc/def/`
//     `abc/def/` --> AppendSuffix('/') --> `abc/def/`
func (b Buf) AppendSuffix(suffix ...byte) Buf {
	return BytesAppendSuffix(b, suffix...)
}

// AppendSuffixString will append string suffix
func (b Buf) AppendSuffixString(suffix string) Buf {
	return BytesAppendSuffixString(b, suffix)
}

// Copy will create new Buf
func (b Buf) Copy() Buf {
	return BytesCopy(b)
}

// Elem gets Nth item of Buf split with delim.
// Eg. `/abc/def/ghi/` => Elem('/', 0) => Buf("")
// Eg. `/abc/def/ghi/` => Elem('/', 1) => Buf("abc")
// Eg. `/abc/def/ghi/` => Elem('/', 3) => Buf("ghi")
func (b Buf) Elem(delim byte, index int) Buf {
	return BytesElem(b, delim, index)
}

// Equal will compare Buf with another Buf
// This will return true if both items are equal
// Otherwise, it will return false.
func (b Buf) Equal(p Buf) bool {
	return BytesEqual(b, p)
}

// HasPrefix will check if Buf has a prefix
func (b Buf) HasPrefix(prefix ...byte) bool {
	return BytesHasPrefix(b, prefix...)
}

// HasPrefixString will check if Buf has a prefix
func (b Buf) HasPrefixString(prefix string) bool {
	return BytesHasPrefixString(b, prefix)
}

// HasSuffix will check if Buf has the suffix
func (b Buf) HasSuffix(suffix ...byte) bool {
	return BytesHasSuffix(b, suffix...)
}

// HasSuffixString will check if Buf has the suffix
func (b Buf) HasSuffixString(suffix string) bool {
	return BytesHasSuffixString(b, suffix)
}

// Index will search first byte c and return its index
// If not found, it will return -1.
func (b Buf) Index(c ...byte) int {
	return BytesIndex(b, c...)
}

// IndexString will search first byte c and return its index
// If not found, it will return -1.
func (b Buf) IndexString(s string) int {
	return BytesIndexString(b, s)
}

// Insert will insert p into the given index of Buf
func (b Buf) Insert(index int, p ...byte) Buf {
	return BytesInsert(b, index, p...)
}

// InsertString will insert string s into the given index of Buf
func (b Buf) InsertString(index int, s string) Buf {
	return BytesInsertString(b, index, s)
}

// LastByte will return a last byte. If Buf were empty, it will return 0.
func (b Buf) LastByte() byte {
	return BytesLastByte(b)
}

// Replace will replace a byte `old` to `new`
func (b Buf) Replace(old, new byte) Buf {
	BytesReplace(b, old, new)
	return b
}

// Reverse will reverse the Buf
func (b Buf) Reverse() Buf {
	BytesReverse(b)
	return b
}

// Shift does not return bool ok as all the moves are within
func (b Buf) Shift(index, length, shift int) Buf {
	BytesShift(b, index, length, shift)
	return b
}

// ToLower will convert Buf to lowercase
func (b Buf) ToLower() Buf {
	BytesToLower(b)
	return b
}

// ToUpper will convert Buf to uppercase
func (b Buf) ToUpper() Buf {
	BytesToUpper(b)
	return b
}

// TrimPrefix will trim a byte prefix if exists
func (b Buf) TrimPrefix(prefix ...byte) Buf {
	return BytesTrimPrefix(b, prefix...)
}

// TrimPrefixString will trim a string prefix if exists
func (b Buf) TrimPrefixString(prefix string) Buf {
	return BytesTrimPrefixString(b, prefix)
}

// TrimSuffix will trim a byte suffix if exists
func (b Buf) TrimSuffix(suffix ...byte) Buf {
	return BytesTrimSuffix(b, suffix...)
}

// TrimSuffixString will trim a string suffix if exists
func (b Buf) TrimSuffixString(suffix string) Buf {
	return BytesTrimSuffixString(b, suffix)
}

// ************************************************************************************************************
// BufItem
// ************************************************************************************************************

// NewBuffer will return a bufItem.
// bufItem uses a pointer and update without need for returning value unlike Buf.
// eg. (Buf)     out = out.WriteString("abc")
//     (bufItem) out.WriteString("abc")
func NewBuffer(size int) bufItem {
	return bufItem{
		Buf: make(Buf, 0, size),
	}
}

// bufItem holds actual buffer Buf in it with a max buffer size maxBufSize
type bufItem struct {
	Buf Buf
}

// Bytes returns []byte of current buffer
func (b *bufItem) Bytes() []byte {
	return b.Buf.Bytes()
}

// Cap returns current capacity
func (b *bufItem) Cap() int {
	return b.Buf.Cap()
}

// Len returns current length
func (b *bufItem) Len() int {
	return b.Buf.Len()
}

// Println prints current buffer
func (b *bufItem) Println() {
	b.Buf.Println()
}

// Reset resets current buffer
func (b *bufItem) Reset() *bufItem {
	b.Buf = b.Buf.Reset()
	return b
}

// Set resets and set current buffer with a given string
func (b *bufItem) Set(s string) *bufItem {
	b.Buf = b.Buf.Set(s)
	return b
}

// String returns current buffer in string format
func (b *bufItem) String() string {
	return b.Buf.String()
}

// Write writes bytes into current buffer
// This meets Writer interface.
func (b *bufItem) Write(p []byte) (n int, err error) {
	return b.Buf.Write(p)
}

// WriteBytes will write byte or bytes to current buffer
func (b *bufItem) WriteBytes(bytes ...byte) *bufItem {
	b.Buf = b.Buf.WriteBytes(bytes...)
	return b
}

// WriteBool will write boolean t to current buffer in string
// true => "true", false => "false"
func (b *bufItem) WriteBool(t bool) *bufItem {
	b.Buf = b.Buf.WriteBool(t)
	return b
}

// WriteFloat will write float f with decimal point dec to current buffer
func (b *bufItem) WriteFloat(f float64, dec uint8) *bufItem {
	b.Buf = b.Buf.WriteFloat(f, dec)
	return b
}

// WriteInt will write integer i to buffer as string format
func (b *bufItem) WriteInt(i int) *bufItem {
	b.Buf = b.Buf.WriteInt(i)
	return b
}

// WriteString will write a string s to current buffer
func (b *bufItem) WriteString(s string) *bufItem {
	b.Buf = b.Buf.WriteString(s)
	return b
}

// WriteStrings will write string slice to current buffer with a delimiter
func (b *bufItem) WriteStrings(s []string, delim ...byte) *bufItem {
	b.Buf = b.Buf.WriteStrings(s, delim...)
	return b
}

// WriteTo will write current buffer to writer w.
func (b *bufItem) WriteTo(w Writer) (n int, err error) {
	return b.Buf.WriteTo(w)
}

// ************************************************************************************************************
// Buffer Pool - Global
// ************************************************************************************************************

// bufPool is a byte Buf pool to be use internally for the logger, etc.
var bufPool = NewBufferPool(256, 2048)

// GetBuffer returns global buffer from the pool
func GetBuffer() *bufItem {
	buf := bufPool.Get()
	return buf
}

// PutBuffer returns buf back to pool
func PutBuffer(buf *bufItem) {
	bufPool.Put(buf)
}

// ************************************************************************************************************
// Buffer Pool
// ************************************************************************************************************

// NewBufferPool creates a buffer pool BufPool
func NewBufferPool(poolSize, bufSize int) BufPool {
	return BufPool{
		pool: Pool{
			New: func() interface{} {
				// bufPoolCreated.Add(1)
				return &bufItem{
					// id: bufPoolCreated.Get(),
					Buf: make([]byte, 0, bufSize),
				}
			},
		}.Init(poolSize),
		maxBufSize: bufSize,
	}
}

// BufPool holds a pool inside that has buffer
type BufPool struct {
	pool       Pool
	maxBufSize int
}

// Get will get buf
func (bp *BufPool) Get() *bufItem {
	buf := bp.pool.Get().(*bufItem)
	buf.Buf = buf.Buf[:0]
	return buf
}

// Put will put buf back to the pool
func (bp *BufPool) Put(buf *bufItem) {
	if cap(buf.Buf) > bp.maxBufSize {
		return
	}
	bp.pool.Put(buf)
}
