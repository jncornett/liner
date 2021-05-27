// Package liner provides functions for indexing and converting between absolute and row/column offsets.
package liner

import (
	"bytes"
	"io"
	"sort"
)

// BuildIndex builds an index of row offsets from r.
func BuildIndex(r io.Reader) ([]int, error) {
	var out []int
	scan := newIndexScanner(r, '\n')
	for scan.Scan() {
		out = append(out, scan.Pos())
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func RowColIndex(index []int, offset int) (row, col int) {
	if offset < 0 {
		panic("negative offset")
	}
	row = sort.SearchInts(index, offset)
	if row == 0 {
		return 0, offset
	}
	return row, offset - index[row-1] - 1
}

func RowCol(source []byte, offset int) (row, col int) {
	index, _ := BuildIndex(bytes.NewReader(source))
	return RowColIndex(index, offset)
}
