package main

import (
	"bufio"
	"fmt"
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

func literal(r *reader) {
	for r.read(5)&0x10 != 0 {
	}
}

func operator(r *reader) (versionSum int) {
	lengthTypeID := r.read(1)
	switch lengthTypeID {
	case 0:
		lenPackets := r.read(15)
		for start := r.pos; r.pos < start+lenPackets; {
			versionSum += packet(r)
		}
	case 1:
		numPackets := r.read(11)
		for ; numPackets > 0; numPackets-- {
			versionSum += packet(r)
		}
	}
	return
}

func packet(r *reader) int {
	version := r.read(3)
	typeID := r.read(3)
	switch typeID {
	case 4:
		literal(r)
		return version
	default:
		return version + operator(r)
	}
}

func main() {
	r := &reader{br: bufio.NewReader(os.Stdin)}
	r.next()
	fmt.Println(packet(r))
}
