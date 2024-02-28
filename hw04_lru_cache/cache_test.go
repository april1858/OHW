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

	t.Run("purge logic", func(t *testing.T) {
		// Write me
		c := NewCache(4)

		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)
		c.Set("d", 4)

		c.Clear()

		for _, val := range []Key{"a", "b", "c", "d"} {
			_, ok := c.Get(val)
			require.False(t, ok)
		}
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
	t.Name()
}

func TestPushing(t *testing.T) {
	t.Run("pushing", func(t *testing.T) {
		cache := []struct {
			key Key
			val interface{}
		}{
			{key: "a", val: 1},
			{key: "b", val: 2},
			{key: "c", val: 3},
			{key: "d", val: 4},
			{key: "e", val: 5},
		}
		c := NewCache(3)
		for _, item := range cache {
			ok := c.Set(item.key, item.val)
			require.False(t, ok)
		}

		val, _ := c.Get("a")
		require.Equal(t, nil, val)

		val, _ = c.Get("b")
		require.Equal(t, nil, val)
	})

	t.Run("pushingOld", func(t *testing.T) {
		cache := []struct {
			key Key
			val interface{}
		}{
			{key: "a", val: 1},
			{key: "b", val: 2},
			{key: "c", val: 3},
		}
		c := NewCache(3)
		for _, item := range cache {
			ok := c.Set(item.key, item.val)
			require.False(t, ok)
		}
		// "a" - is the last element (back) and "c" -  is the first element (front)

		val, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 1, val)
		// "b" - now is the last element (back)

		ok = c.Set("b", 22)
		require.True(t, ok)

		// "c" - now is the last element (back)

		ok = c.Set("d", 4)
		require.False(t, ok)

		val, _ = c.Get("c")
		require.Equal(t, nil, val)
	})
}
