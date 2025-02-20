package pokecache

import (
	"testing"
	"time"
)

func TestAddGetCacheEntry(t *testing.T) {
	cases := map[string]struct {
		url  string
		data []byte
	}{
		"test 1": {
			url:  "https://api.example.com/test1/",
			data: []byte("test 1 data"),
		},
		"test 2": {
			url:  "https://api.example.com/test2/",
			data: []byte("test 2 data"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(time.Second)
			err := cache.Add(c.url, c.data)
			if err != nil {
				t.Errorf("error adding entry to cache: %v", err)
				return
			}

			entry, ok := cache.Get(c.url)
			if !ok {
				t.Errorf("error getting entry from cache")
				return
			}

			if string(entry) != string(c.data) {
				t.Errorf("unexpected results: entry: %v != c.data: %v", entry, c.data)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cases := map[string]struct {
		url  string
		data []byte
	}{
		"test 1": {
			url:  "https://api.example.com/test1/",
			data: []byte("test 1 data"),
		},
		"test 2": {
			url:  "https://api.example.com/test2/",
			data: []byte("test 2 data"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(time.Second)
			err := cache.Add(c.url, c.data)
			if err != nil {
				t.Errorf("error adding entry to cache: %v", err)
				return
			}

			_, ok := cache.Get(c.url)
			if !ok {
				t.Errorf("unexpected results: entry not added")
				return
			}

			time.Sleep(time.Second * 2)

			_, ok = cache.Get(c.url)
			if ok {
				t.Errorf("unexpected results: entry should be reaped")
				return
			}
		})
	}
}
