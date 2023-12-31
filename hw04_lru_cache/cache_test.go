package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("logic of pushing out items due to queue size", func(t *testing.T) {
		c := NewCache(4)

		ok := c.Set("a", 1)
		require.False(t, ok)

		ok = c.Set("b", 2)
		require.False(t, ok)

		ok = c.Set("c", 3)
		require.False(t, ok)

		ok = c.Set("d", 4)
		require.False(t, ok)

		ok = c.Set("e", 5)
		require.False(t, ok)

		val, ok := c.Get("a")
		require.Nil(t, val)
		require.False(t, ok)

		val, ok = c.Get("b")
		require.Equal(t, 2, val)
		require.True(t, ok)

		val, ok = c.Get("c")
		require.Equal(t, 3, val)
		require.True(t, ok)

		val, ok = c.Get("d")
		require.Equal(t, 4, val)
		require.True(t, ok)

		val, ok = c.Get("e")
		require.Equal(t, 5, val)
		require.True(t, ok)
	})

	t.Run("logic of pushing out long-used items", func(t *testing.T) {
		c := NewCache(4)

		ok := c.Set("a", 1)
		require.False(t, ok)

		ok = c.Set("b", 2)
		require.False(t, ok)

		ok = c.Set("c", 3)
		require.False(t, ok)

		ok = c.Set("d", 4)
		require.False(t, ok)

		val, ok := c.Get("a")
		require.Equal(t, 1, val)
		require.True(t, ok)

		ok = c.Set("e", 5)
		require.False(t, ok)

		val, ok = c.Get("b")
		require.Nil(t, val)
		require.False(t, ok)

		val, ok = c.Get("a")
		require.Equal(t, 1, val)
		require.True(t, ok)

		val, ok = c.Get("c")
		require.Equal(t, 3, val)
		require.True(t, ok)

		val, ok = c.Get("d")
		require.Equal(t, 4, val)
		require.True(t, ok)

		val, ok = c.Get("e")
		require.Equal(t, 5, val)
		require.True(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
