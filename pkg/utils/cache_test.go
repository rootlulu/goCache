package utils

import (
	"fmt"
	"testing"
)

func TestCacheNew(t *testing.T) {
	t.Run(
		"new cache", func(t *testing.T) {
			cache1 := NewCache[string, string]()
			if cache1 == nil {
				t.Errorf("new cache failed")
			}
			cache1.Set("key", "value")
			if cache1.Get("key") != "value" {
				t.Errorf("set the cache failed, get the value: %s, wanted: %s", cache1.Get("key"), "value")
			}

			cache2 := NewCache[int, string]()
			if cache2 == nil {
				t.Errorf("new cache failed")
			}
			cache2.Set(1, "value")
			if cache2.Get(1) != "value" {
				t.Errorf("set the cache failed, get the value: %s, wanted: %s", cache2.Get(1), "value")
			}
			v, ok := cache2.GetBool(1)
			if !ok {
				t.Errorf("get the cache failed")
			}
			if v != "value" {
				t.Errorf("get the cache failed, get the value: %s, wanted: %s", v, "value")
			}
		},
	)
}

func TestCacheRange(t *testing.T) {
	t.Run(
		"range the cache", func(t *testing.T) {
			cache := NewCache[string, string]()
			cache.Set("key1", "value1")
			cache.Set("key2", "value2")
			cache.Set("key3", "value3")
			cache.Set("key4", "value4")
			cache.Set("key5", "value5")
			cache.Set("key6", "value6")
			cache.Set("key7", "value7")
			cache.Set("key8", "value8")
			cache.Set("key9", "value9")
			cache.Set("key10", "value10")

			for item := range cache.Range() {
				fmt.Printf("the key: %s, the value: %s", item.Key, item.Val)
			}
		},
	)
}
