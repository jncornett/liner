package main

import (
	"bufio"
	"bytes"
	"log"
)

func main() {
	data := bytes.NewReader([]byte("abbbbbbbbbbbbbbbbba\ncbbbbbc\ndcccd"))
	buf := bufio.NewReaderSize(data, 2)
	log.Println(buf.Size())
	p, err := buf.ReadSlice('\n')
	log.Printf("%q, %v", string(p), err)
	p, err = buf.ReadSlice('\n')
	log.Printf("%q, %v", string(p), err)
}
