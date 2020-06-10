package cache

import (
	"fmt"
	"github.com/allegro/bigcache"
	"time"
)

type BigCache struct {
}

func (bc *BigCache) Test() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))
}
