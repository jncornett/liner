package liner

import (
	"bufio"
	"errors"
	"io"
)

type indexScanner struct {
	buf   *bufio.Reader
	pos   int
	delim byte
	err   error
}

func newIndexScanner(r io.Reader, delim byte) *indexScanner {
	return &indexScanner{
		buf:   bufio.NewReader(r),
		delim: delim,
	}
}

func (s *indexScanner) Scan() bool {
	for {
		b, err := s.buf.ReadByte()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				s.err = err
			}
			return false
		}
		s.pos++
		if b == s.delim {
			return true
		}
	}
}

func (s *indexScanner) Err() error {
	return s.err
}

func (s *indexScanner) Pos() int {
	return s.pos - 1
}
