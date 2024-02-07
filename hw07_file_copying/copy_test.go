package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type errorTestCase struct {
	from     string
	to       string
	limit    int64
	offset   int64
	expected error
}

func TestCopy(t *testing.T) {
	testCaseOffsetExceeds := errorTestCase{
		from:     "testdata/input.txt",
		to:       "out_offset0_limit0.txt",
		offset:   7000,
		limit:    0,
		expected: ErrOffsetExceedsFileSize,
	}

	t.Run("Error offset exceeds file size", func(t *testing.T) {
		err := Copy(testCaseOffsetExceeds.from, testCaseOffsetExceeds.to, testCaseOffsetExceeds.offset, testCaseOffsetExceeds.limit)
		require.ErrorIs(t, err, testCaseOffsetExceeds.expected)
	})
}
