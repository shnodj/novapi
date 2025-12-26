package backend

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

// InitMemoryCache initializes the memory cache
func InitMemoryCache() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	Cache = cache.New(5*time.Minute, 10*time.Minute)
	log.Println("Memory cache initialized")
}
