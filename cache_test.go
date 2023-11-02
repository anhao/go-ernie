package go_ernie

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()

	cache.Set("test", "test", 3*time.Second)

	value, ok := cache.Get("test")
	if ok {
		t.Log(value == "test")
	}
	time.Sleep(3 * time.Second)
	_, exists := cache.Get("test")
	t.Log(exists)
}
