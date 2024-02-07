package main

import (
	"testing"
)

func TestCopy(t *testing.T) {
	// testCases := []struct {
	// 	from     string
	// 	to       string
	// 	limit    int64
	// 	offset   int64
	// 	expected string
	// }{
	// 	{
	// 		from:     "testdata/input.txt",
	// 		to:       "out_offset0_limit0.txt",
	// 		offset:   0,
	// 		limit:    0,
	// 		expected: "testdata/out_offset0_limit0.txt",
	// 	},
	// 	{
	// 		from:     "testdata/input.txt",
	// 		to:       "out_offset0_limit10.txt",
	// 		offset:   0,
	// 		limit:    10,
	// 		expected: "testdata/out_offset0_limit10.txt",
	// 	},
	// 	{
	// 		from:     "testdata/input.txt",
	// 		to:       "out_offset0_limit1000.txt",
	// 		offset:   0,
	// 		limit:    1000,
	// 		expected: "testdata/out_offset0_limit1000.txt",
	// 	},
	// }

	// for _, testCase := range testCases {
	// 	testCase := testCase
	// 	t.Run(testCase.to, func(t *testing.T) {
	// 		err := Copy(testCase.from, testCase.to, testCase.offset, testCase.limit)
	// 		require.NoError(t, err)
	// 		cmp := equalfile.New(nil, equalfile.Options{})
	// 		isFileEqual, err := cmp.CompareFile(testCase.to, testCase.expected)
	// 		require.NoError(t, err)
	// 		require.Equal(t, true, isFileEqual)
	// 	})
	// }
}
