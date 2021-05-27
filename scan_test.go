package liner

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexScanner(t *testing.T) {
	scan := newIndexScanner(
		bufio.NewReader(
			strings.NewReader(
				"abc\ndef\n",
			),
		),
		'\n',
	)
	require.True(t, scan.Scan())
	assert.Equal(t, 3, scan.Pos())
	require.True(t, scan.Scan())
	assert.Equal(t, 7, scan.Pos())
	assert.False(t, scan.Scan())
	assert.NoError(t, scan.Err())
}

func TestIndexScanner_zeroLength(t *testing.T) {
	scan := newIndexScanner(
		bufio.NewReader(
			bytes.NewReader(nil),
		),
		'\n',
	)
	require.False(t, scan.Scan())
	assert.NoError(t, scan.Err())
}

func TestIndexScanner_consecutiveDelimiters(t *testing.T) {
	scan := newIndexScanner(
		bufio.NewReader(
			strings.NewReader("\n\n"),
		),
		'\n',
	)
	require.True(t, scan.Scan())
	assert.Equal(t, 0, scan.Pos())
	require.True(t, scan.Scan())
	assert.Equal(t, 1, scan.Pos())
	assert.False(t, scan.Scan())
	assert.NoError(t, scan.Err())
}

func TestIndexScanner_propagateReaderError(t *testing.T) {
	scan := newIndexScanner(
		readerFunc(func([]byte) (int, error) {
			return 0, errors.New("oops")
		}),
		'\n',
	)
	require.False(t, scan.Scan())
	assert.Error(t, scan.Err())
}

type readerFunc func([]byte) (int, error)

func (fn readerFunc) Read(p []byte) (int, error) { return fn(p) }
