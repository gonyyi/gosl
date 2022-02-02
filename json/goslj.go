// (c) Gon Y. Yi 2022 <https://gonyyi.com/copyright>
//
// JSON is a part of Gosl package's extended package. JSON is a zero allocation JSON builder.
// Creating a JSON allocates memory each time and for microservices, this gives stress to the garbage collection.
// Especially when it comes to the cloud services, speed, memory usages, and frequency of garbage collection will
// directly correlate to the cost.

package goslj

import "github.com/gonyyi/gosl"

var BufferSize = 1024 // Sets JSON's buffer size when created

// NewPool will create a pool of JSON
func NewPool(PoolSize int) *Pool {
	p := &Pool{}
	p.pool = p.pool.Init(PoolSize)
	p.pool.New = func() interface{} {
		kvj := &JSON{
			buf: make(gosl.Buf, 0, BufferSize),
		}
		p.created += 1 // todo: need mutex
		kvj.pool = p
		return kvj
	}
	return p
}

// Pool is JSON pool
type Pool struct {
	pool    gosl.Pool
	created int
	inUse   int
}

// Stats will return how many objects were created and how many are in use.
func (p *Pool) Stats() (created, inUse int) {
	return p.created, p.inUse
}

// Get will obtain *JSON from the pool
func (p *Pool) Get() *JSON {
	p.inUse += 1
	return p.pool.Get().(*JSON).Reset()
}

// Put will put *JSON to the pool
// This can be done by `*JSON.Putback()` as well
func (p *Pool) Put(kvj *JSON) {
	p.inUse -= 1
	p.pool.Put(kvj)
}

// NewJSON takes a buffer size and creates a JSON
func NewJSON(bufSize int) *JSON {
	if bufSize < 0 {
		bufSize = BufferSize
	}
	return &JSON{
		buf: make(gosl.Buf, 0, bufSize),
	}
}

// JSON is a very simple writer for JSON without memory allocation
type JSON struct {
	pool *Pool // to able to self-return
	buf  gosl.Buf
}

// Reset will clear current JSON
func (j *JSON) Reset() *JSON {
	j.buf = j.buf.Reset()
	return j
}

// Start will begin JSON
func (j *JSON) Start() *JSON {
	return j.b('{')
}

// End will remove extra comma if exists, and also add '}'
func (j *JSON) End() *JSON {
	return j.rmLast(',').b('}')
}

// String will add key-value pair of string
func (j *JSON) String(name, s string) *JSON {
	return j.string(name).b(':').string(s).b(',')
}

// Int will add key-value pair of integer
func (j *JSON) Int(name string, i int) *JSON {
	return j.string(name).b(':').int(i).b(',')
}

// IntArray will add integers
func (j *JSON) IntArray(name string, nums ...int) *JSON {
	j.string(name).b(':')
	j.b('[')
	for _, num := range nums {
		j.int(num)
		j.b(',')
	}
	j.rmLast(',')
	j.b(']')
	j.b(',')
	return j
}

// StringArray will add strings
func (j *JSON) StringArray(name string, s ...string) *JSON {
	j.string(name).b(':')
	j.b('[')
	for _, v := range s {
		j.string(v)
		j.b(',')
	}
	j.rmLast(',')
	j.b(']')
	j.b(',')
	return j
}

// Write writes JSON to the Writer
func (j *JSON) Write(w gosl.Writer) *JSON {
	j.rmLast(',')
	j.buf.WriteTo(w)
	return j
}

// Sub will take other JSON and add
func (j *JSON) Sub(name string, src *JSON) *JSON {
	if j == src { // do not allow self being included
		return j
	}
	j.string(name).b(':')
	j.buf = j.buf.WriteBytes(src.buf...)
	return j.b(',')
}

// Putback will return JSON to the pool if it was from the pool
func (j *JSON) Putback() bool {
	if j.pool != nil {
		j.pool.Put(j)
		return true
	}
	return false
}

// b will add a byte
func (j *JSON) b(b byte) *JSON {
	j.buf = append(j.buf, b)
	return j
}

// rmLast will remove byte b suffix if exists
func (j *JSON) rmLast(b byte) *JSON {
	j.buf = j.buf.TrimSuffix(b)
	return j
}

// int will convert an integer and add
func (j *JSON) int(i int) *JSON {
	j.buf = gosl.BytesAppendInt(j.buf, i)
	return j
}

// string will add string s
func (j *JSON) string(s string) *JSON {
	j.buf = j.buf.WriteBytes('"')

	for i := 0; i < len(s); i++ {
		if app := stringEscapes[s[i]]; app != 0 {
			j.buf = append(j.buf, '\\', app)
		} else {
			j.buf = append(j.buf, s[i])
		}
	}

	j.buf = j.buf.WriteBytes('"')
	return j
}

// stringEscapes will hold what strings need to be escaped
// At this point, &, <, > will not be converted to \u0026, \u003c, \u003e. Not sure if I need to...
var stringEscapes = [256]byte{'"': '"', '\\': '\\', '\r': 'r', '\n': 'n', '\b': 'b', '\f': 'f', '\t': 't'}
