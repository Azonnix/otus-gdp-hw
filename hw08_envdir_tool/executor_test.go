package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("failed return code", func(t *testing.T) {
		code := RunCmd([]string{}, nil)
		require.Equal(t, 1, code)
	})
}
