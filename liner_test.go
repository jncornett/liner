package liner

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildIndex(t *testing.T) {
	tests := map[string][]int{
		"":        nil,
		"aaa":     nil,
		"aaa\n":   {3},
		"\n":      {0},
		"\naaa":   {0},
		"\n\n":    {0, 1},
		"aaa\n\n": {3, 4},
		"\n\naaa": {0, 1},
		"\naaa\n": {0, 4},
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			got, err := BuildIndex(strings.NewReader(input))
			require.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}

func TestBuildIndex_readError(t *testing.T) {
	_, err := BuildIndex(readerFunc(func([]byte) (int, error) {
		return 0, errors.New("oops")
	}))
	require.Error(t, err)
}

func TestRowCol(t *testing.T) {
	tests := []struct {
		source           string
		offset           int
		wantRow, wantCol int
	}{
		{"", 0, 0, 0},
		{"", 1, 0, 1},
		{"", 2, 0, 2},
		{"abc", 0, 0, 0},
		{"abc", 1, 0, 1},
		{"abc", 2, 0, 2},
		{"abc", 3, 0, 3},
		{"abc\ndef", 0, 0, 0},
		{"abc\ndef", 1, 0, 1},
		{"abc\ndef", 2, 0, 2},
		{"abc\ndef", 3, 0, 3},
		{"abc\ndef", 4, 1, 0},
		{"abc\ndef", 5, 1, 1},
		{"abc\ndef", 6, 1, 2},
		{"abc\ndef", 7, 1, 3},
		{"abc\ndef\n", 7, 1, 3},
		{"abc\ndef\n", 8, 2, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("[%q, %d]", string(tt.source), tt.offset), func(t *testing.T) {
			row, col := RowCol([]byte(tt.source), tt.offset)
			assert.Equal(t, tt.wantRow, row)
			assert.Equal(t, tt.wantCol, col)
		})
	}
}
