package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond
	c := NewCache(interval)
	if c.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond
	c := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
		{
			inputKey: "",
			inputVal: []byte("value3"),
		},
	}

	for _, cas := range cases {
		c.Add(cas.inputKey, cas.inputVal)

		actual, ok := c.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesn't match %s", string(actual), string(cas.inputVal))
			continue
		}
	}
}

func TestCacheReapLoop(t *testing.T) {
	interval := time.Millisecond * 10
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("value1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := c.cache[keyOne]
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestCacheReapLoopFail(t *testing.T) {
	interval := time.Millisecond * 10
	c := NewCache(interval)

	keyOne := "key1"
	c.Add(keyOne, []byte("value1"))

	time.Sleep(interval / 2)

	_, ok := c.cache[keyOne]
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}
}
