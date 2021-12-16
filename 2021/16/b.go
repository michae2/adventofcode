package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type nybble byte

type reader struct {
	br  *bufio.Reader
	buf nybble
	pos int
}

func (r *reader) next() {
	hex, err := r.br.ReadByte()
	if err != nil || hex < '0' || (hex > '9' && hex < 'A') || hex > 'F' {
		panic("parser error")
	} else if r.pos%4 != 0 {
		panic("reader error")
	}
	if hex >= 'A' {
		r.buf = nybble(hex - 'A' + 10)
	} else {
		r.buf = nybble(hex - '0')
	}
}

func (r *reader) read(n int) (val int) {
	for n > 0 {
		avail := 4 - r.pos%4
		mask := 1<<avail - 1
		temp := int(r.buf) & mask
		if n >= avail {
			temp <<= n - avail
			r.pos += avail
			n -= avail
			r.next()
		} else {
			temp >>= avail - n
			r.pos += n
			n -= n
		}
		val += temp
	}
	return
}

func literal(r *reader) (val int64) {
	for group := 0x10; group&0x10 != 0; {
		group = r.read(5)
		val <<= 4
		val += int64(group & 0x0F)
	}
	return
}

func sum(r *reader) (val int64) {
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		lenPackets := r.read(15)
		for start := r.pos; r.pos < start+lenPackets; {
			val += packet(r)
		}
	case 1:
		numPackets := r.read(11)
		for ; numPackets > 0; numPackets-- {
			val += packet(r)
		}
	}
	return
}

func product(r *reader) (val int64) {
	val = 1
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		lenPackets := r.read(15)
		for start := r.pos; r.pos < start+lenPackets; {
			val *= packet(r)
		}
	case 1:
		numPackets := r.read(11)
		for ; numPackets > 0; numPackets-- {
			val *= packet(r)
		}
	}
	return
}

func minimum(r *reader) (val int64) {
	val = math.MaxInt64
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		lenPackets := r.read(15)
		for start := r.pos; r.pos < start+lenPackets; {
			if v := packet(r); v < val {
				val = v
			}
		}
	case 1:
		numPackets := r.read(11)
		for ; numPackets > 0; numPackets-- {
			if v := packet(r); v < val {
				val = v
			}
		}
	}
	return
}

func maximum(r *reader) (val int64) {
	val = math.MinInt64
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		lenPackets := r.read(15)
		for start := r.pos; r.pos < start+lenPackets; {
			if v := packet(r); v > val {
				val = v
			}
		}
	case 1:
		numPackets := r.read(11)
		for ; numPackets > 0; numPackets-- {
			if v := packet(r); v > val {
				val = v
			}
		}
	}
	return
}

func gt(r *reader) int64 {
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		r.read(15)
	case 1:
		r.read(11)
	}
	if packet(r) > packet(r) {
		return 1
	}
	return 0
}

func lt(r *reader) int64 {
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		r.read(15)
	case 1:
		r.read(11)
	}
	if packet(r) < packet(r) {
		return 1
	}
	return 0
}

func eq(r *reader) int64 {
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		r.read(15)
	case 1:
		r.read(11)
	}
	if packet(r) == packet(r) {
		return 1
	}
	return 0
}

func packet(r *reader) int64 {
	r.read(3)
	typeID := r.read(3)
	switch typeID {
	case 0:
		return sum(r)
	case 1:
		return product(r)
	case 2:
		return minimum(r)
	case 3:
		return maximum(r)
	case 4:
		return literal(r)
	case 5:
		return gt(r)
	case 6:
		return lt(r)
	case 7:
		return eq(r)
	}
	panic("parser error")
}

func main() {
	r := &reader{br: bufio.NewReader(os.Stdin)}
	r.next()
	fmt.Println(packet(r))
}
