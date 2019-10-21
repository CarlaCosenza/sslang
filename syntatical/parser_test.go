package syntatical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	tt := map[string]struct {
		s string

		n  int
		ok bool
	}{
		"test invalid rule": {
			s:  "1324",
			n:  -1,
			ok: false,
		},
		"test valid rule": {
			s:  "s123",
			n:  123,
			ok: true,
		},
		"test empty rule": {
			s:  "",
			n:  -1,
			ok: false,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			n, ok := shift(table.s)

			assert.Equal(t, table.n, n)
			assert.Equal(t, table.ok, ok)
		})
	}
}

func TestReduce(t *testing.T) {
	tt := map[string]struct {
		s string

		n  Rule
		ok bool
	}{
		"test invalid rule": {
			s:  "1324",
			n:  -1,
			ok: false,
		},
		"test valid rule": {
			s:  "r123",
			n:  123,
			ok: true,
		},
		"test empty rule": {
			s:  "",
			n:  -1,
			ok: false,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			n, ok := reduce(table.s)

			assert.Equal(t, table.n, n)
			assert.Equal(t, table.ok, ok)
		})
	}
}
